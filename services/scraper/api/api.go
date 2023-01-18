package api

import (
	"context"
	"github.com/kanade0404/tenhou-log/pkg/config"
	"github.com/kanade0404/tenhou-log/services/ent"
)

type Api struct {
	context.Context
	client *ent.Client
	config *config.Config
}

func New(ctx context.Context, client *ent.Client, cnf *config.Config) *Api {
	return &Api{
		ctx,
		client,
		cnf,
	}
}
