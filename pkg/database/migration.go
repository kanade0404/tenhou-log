package database

import (
	atlasmigrate "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/services/ent"
	"github.com/kanade0404/tenhou-log/services/ent/migrate"
)

func Migration(ctx context.Context, client *ent.Client, env *config.Config) error {
	if config.IsLocal() {
		if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
			return fmt.Errorf("failed creation schema resource: %v", err)
		}
	} else {
		// Create a local migration directory able to understand dbmate migration file format for replay.
		dir, err := sqltool.NewDBMateDir("./migrations")
		if err != nil {
			return fmt.Errorf("failed creating migration directory: %v", err)
		}
		// Write atlas.sum file to the local migration directory.
		sum, err := dir.Checksum()
		if err != nil {
			return fmt.Errorf("failed check atlas.sum: %v", err)
		}
		if err := atlasmigrate.WriteSumFile(dir, sum); err != nil {
			return fmt.Errorf("failed writing atlas.sum: %v", err)
		}
		// Migrate diff options.
		opts := []schema.MigrateOption{
			schema.WithDir(dir),                         // provide migration directory
			schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
			schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		}
		// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
		if err := migrate.Diff(ctx, env.ConnectionString(), opts...); err != nil {
			return fmt.Errorf("failed generating migration file: %v", err)
		}
	}
	return nil
}
