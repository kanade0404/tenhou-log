package main

import (
	"context"
	"database/sql"
	"external/driver/secret_manager"
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
	ctx := context.Background()
	isLocal := configs.IsLocal()
	var env *configs.Env
	if isLocal {
		env = configs.NewLocalEnv()
	} else {
		sm, err := secret_manager.NewSecretManager(ctx)
		if err != nil {
			log.Fatal(err)
		}
		env, err = configs.NewDevEnv(sm)
		if err != nil {
			log.Fatal(err)
		}
	}
	port := env.AppPort
	if port != "" {
		port = "8088"
	}
	db, err := sql.Open(env.Dialect, env.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)
	router := mux.NewRouter()
	a := &api.Api{}
	a.Context = ctx
	a.Env = env
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
