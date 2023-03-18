package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/pkg/database"
	"github.com/kanade0404/tenhou-log/services/scraper/api"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	log.Println("Run start")
	ctx := context.Background()
	env, err := config.NewEnv(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := database.CreateClient(ctx, env.Dialect(), env.ConnectionString())
	if err != nil {
		log.Fatalf("failed database connection: %v", err)
	}
	if err := database.Migration(ctx, client, env); err != nil {
		log.Fatalln(err)
	}
	router := mux.NewRouter()
	a := api.New(ctx, client, env)
	router.HandleFunc("/scraping", a.Scraper).Methods("POST")
	port := env.Port()
	if port != "" {
		port = "8088"
	}
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}
	log.Printf("server listen %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed listening server: %v", err)
	}
}
