package main

import (
	atlasmigrate "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/pkg/driver/secret_manager"
	"github.com/kanade0404/tenhou-log/services/ent"
	"github.com/kanade0404/tenhou-log/services/ent/migrate"
	"github.com/kanade0404/tenhou-log/services/scraper/api"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	log.Println("Run start")
	ctx := context.Background()
	isLocal := config.IsLocal()
	var (
		env *config.Config
		err error
	)
	if isLocal {
		env, err = config.NewLocalEnv()
		if err != nil {
			log.Fatalf("failed load dev env: %v", err)
		}
	} else {
		sm, err := secret_manager.NewSecretManager(ctx)
		if err != nil {
			log.Fatalf("failed initialize SecretManager: %v", err)
		}
		env, err = config.NewRemoteEnv(sm)
		if err != nil {
			log.Fatalf("failed load remote env: %v", err)
		}
	}
	port := env.Port()
	if port != "" {
		port = "8088"
	}
	client, err := ent.Open(env.Dialect(), env.ConnectionString())
	if err != nil {
		log.Fatalf("failed database connection: %v", err)
	}
	if config.IsLocal() {
		// ローカル環境
		if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
			log.Fatalf("failed creation schema resouce: %v", err)
		}
	} else {
		// リモート環境
		dir, err := sqltool.NewDBMateDir("./migrations")
		if err != nil {
			log.Fatalf("failed creating migration directory: %v", err)
		}
		sum, err := dir.Checksum()
		if err != nil {
			log.Fatalf("failed check atlas.sum: %v", err)
		}
		if err := atlasmigrate.WriteSumFile(dir, sum); err != nil {
			log.Fatalf("failed writing atlas.sum: %v", err)
		}
		opts := []schema.MigrateOption{
			schema.WithDir(dir),                         // provide migration directory
			schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
			schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		}
		if err := migrate.Diff(ctx, env.ConnectionString(), opts...); err != nil {
			log.Fatalf("failed generationg migration file: %v", err)
		}
	}
	router := mux.NewRouter()
	a := api.New(ctx, client, env)
	router.HandleFunc("/scraping", a.Scraper).Methods("POST")
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}
	log.Printf("server listen %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed listening server: %v", err)
	}
}
