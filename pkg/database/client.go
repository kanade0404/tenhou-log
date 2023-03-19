package database

import (
	"context"
	"fmt"
	"github.com/kanade0404/tenhou-log/services/ent"
)

func CreateClient(ctx context.Context, dialect, connectionString string) (*ent.Client, error) {
	client, err := ent.Open(dialect, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed database connection: %v", err)
	}
	return client, nil
}
