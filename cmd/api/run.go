package api

import (
	"context"

	log "github.com/google/logger"
	"github.com/roger-russel/go-k8s-cli/internal/api"
	"github.com/roger-russel/go-k8s-cli/internal/config"
	"github.com/roger-russel/go-k8s-cli/pkg/core"
)

func Run() {
	ctx := context.Background()
	err := config.LoadEnvs(ctx)

	if err != nil {
		log.Errorf("Failed to load environment variables %v", err)
		core.Exit(1)
	}

	api.Server(ctx, api.Config{
		Addr: config.Envs.APIAddress,
	})
}
