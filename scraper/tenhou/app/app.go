package app

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"tenhou/configs"
	"tenhou/entities"
	"tenhou/internal"
	"time"
)

type App struct {
	*mux.Router
	DB *sql.DB
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	Response(w, code, map[string]string{"error": message})
}
func Response(w http.ResponseWriter, code int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func (a *App) Scraper(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://tenhou.net/sc/raw/list.cgi")
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
	}(res.Body)
	if res.StatusCode != 200 {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("text: ", doc.Text())
	logFiles := internal.UnmarshalFileListText(doc.Text())
	log.Printf("logs: \n%v", logFiles)
	var logs []*entities.Log
	for _, str := range logFiles {
		log.Println("str: ", str)
		l, err := entities.Unmarshal(str)
		if err != nil {
			log.Println("error unmarshal entity: ", err)
			ErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		logs = append(logs, l)
	}
	//rgx := regexp.MustCompile(`[a-z]`)
	for _, l := range logs {
		client := new(http.Client)
		compressLogURL := fmt.Sprintf("http://tenhou.net/sc/raw/dat/%s", l.File)
		log.Println(compressLogURL)
		request, err := http.NewRequest("GET", compressLogURL, nil)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to initialize request. err: %s", err.Error()))
			return
		}
		request.Header.Add("Accept-Encoding", "gzip")

		res, err := client.Do(request)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to requests log file. err: %s", err.Error()))
			return
		}
		defer res.Body.Close()
		var f []byte
		if _, err := res.Body.Read(f); err != nil {
			ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to read request body. err: %s", err.Error()))
			return
		}
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to initialize gzip reader. err: %s", err.Error()))
			return
		}
		defer reader.Close()
		g, err := ioutil.ReadAll(reader)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to read gzip. err: %s", err.Error()))
			return
		}
		err = ioutil.WriteFile(fmt.Sprintf("/tmp/%s", strings.ReplaceAll(l.File, "log", "html")), g, os.ModePerm)
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to write gzip. err: %s", err.Error()))
			return
		}
		//// L0001 | 21:06 | 四－東喰赤－ | <a href="http://tenhou.net/0/?log=2022091221gm-0101-0001-4738510e">牌譜</a> | こり蔵(+23.5) ゲルルン(+8.9) 千葉ぱぐぱぐ(-9.0) ひみつのあつこ(-23.4)<br>
		//logText := string(g)
		//log.Println(logText)
		//logURLRegexp := regexp.MustCompile("http://tenhou.net/0/\\?log=[a-z\\d\\-]*")
		//logURL := logURLRegexp.FindString(logText)
		//log.Println(logURL)
		//u, err := url.Parse(logURL)
		//if err != nil {
		//	ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to parse url. err: %s", err.Error()))
		//	return
		//}
		//var ID string
		//for key, values := range u.Query() {
		//	if key == "log" {
		//		for _, value := range values {
		//			ID = value
		//		}
		//	}
		//}
		//rawLogBaseURL := "https://tenhou.net/0/log/?"
		//rawLogURL := strings.ReplaceAll(ID, "log", "html")
		//log.Println(rawLogURL)
		//res, err = http.Get(fmt.Sprintf("%s%s", rawLogBaseURL, rawLogURL))
		//if err != nil {
		//	ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to parse url. err: %s", err.Error()))
		//	return
		//}
		//var data []byte
		//defer res.Body.Close()
		//_, err = res.Body.Read(data)
		//if err != nil {
		//	ErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to parse url. err: %s", err.Error()))
		//	return
		//}
		//log.Println(string(data))
	}
	Response(w, http.StatusOK, logs)
}

func (a *App) Run() {
	log.Println("Run start")
	a.Router = mux.NewRouter()
	env := configs.NewEnv()
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DB.UserName, env.DB.Password, env.DB.Host, env.DB.Port, env.DB.Name)
	schemaURL := "file://database/schema"
	m, err := migrate.New(schemaURL, connectionString)
	if err != nil {
		log.Fatal("database connection error: ", err)
	}
	err = m.Up()
	if err != migrate.ErrNoChange {
		log.Fatal("database migration error: ", err)
	}
	port := env.App.Port
	if port == "" {
		port = "8000"
	}
	db, err := sql.Open(env.DB.Dialect, env.DB.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)
	a.DB = db
	a.Router.HandleFunc("/", a.Scraper)
	server := &http.Server{
		Handler:      a.Router,
		Addr:         fmt.Sprintf("0.0.0.0:%s", env.App.Port),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	log.Println(fmt.Sprintf("server listen :%s", port))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
