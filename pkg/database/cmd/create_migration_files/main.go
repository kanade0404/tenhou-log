package main

import (
	atlas "ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/services/ent/migrate"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
)

/*
atlasのversioned migration fileを作成する。
*/
func main() {
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required.")
	}
	ctx := context.Background()
	f, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	dir, err := atlas.NewLocalDir(fmt.Sprintf("%s/%s/%s", filepath.Dir(f), filepath.Base(f), "services/database/migrations"))
	if err != nil {
		log.Fatalln(err)
	}
	c, err := config.NewEnv(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	if err = migrate.NamedDiff(ctx, c.ConnectionString(), os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
