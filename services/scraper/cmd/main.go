package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"net/http"
	"scraper/api"
	"scraper/configs"
)

func main() {
	log.Println("Run start")
	config := configs.NewEnv()
	port := config.App.Port
	if port != "" {
		port = "8088"
	}
	db, err := sql.Open(config.DB.Dialect, config.DB.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)
	router := mux.NewRouter()
	a := &api.Api{}
	a.Context = context.Background()
	a.Env = config
	a.Db = db
	router.HandleFunc("/scraping", a.Scraper).Methods("POST")
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}
	log.Printf("server listen %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
