package database

import (
	"context"
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
	}
	return nil
}
