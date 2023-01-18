package database

import (
	atlasmigrate "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/kanade0404/tenhou-log/services/ent/migrate"
	"log"
)

func Migration() {
	ctx := context.Background()
	// Create a local migration directory able to understand dbmate migration file format for replay.
	dir, err := sqltool.NewDBMateDir("./migrations")
	if err != nil {
		log.Fatalf("failed creating migration directory: %v", err)
	}
	// Write atlas.sum file to the local migration directory.
	sum, err := dir.Checksum()
	if err != nil {
		log.Fatalf("failed check atlas.sum: %v", err)
	}
	if err := atlasmigrate.WriteSumFile(dir, sum); err != nil {
		log.Fatalf("failed writing atlas.sum: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.Diff(ctx, "mysql://root:password@localhost:3306/atlas_dev_database", opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
