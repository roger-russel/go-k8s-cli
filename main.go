package main

import (
	"os"

	"github.com/roger-russel/go-k8s-cli/cmd/api"
	"github.com/roger-russel/go-k8s-cli/cmd/cli"
)

func main() {
	if os.Getenv("MODE") == "API" {
		api.Run()
	} else {
		cli.Run()
	}
}
