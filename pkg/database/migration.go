package database

import (
	"context"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/services/ent"
	"github.com/kanade0404/tenhou-log/services/ent/migrate"
	"log"
)

func Migration(ctx context.Context, client *ent.Client) error {
	if config.IsLocal() {
		log.Println("Exec migration.")
		if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
			return fmt.Errorf("failed creation schema resource: %v", err)
		}
	}
	return nil
}
