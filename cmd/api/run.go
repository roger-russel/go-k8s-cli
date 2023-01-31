package api

import (
	"os"

	"github.com/roger-russel/go-k8s-cli/internal/api"
)

func Run() {
	addr := os.Getenv("API_ADDRESS")
	if addr == "" {
		addr = ":8080"
	}

	api.Server(api.Config{
		Addr: addr,
	})
}
