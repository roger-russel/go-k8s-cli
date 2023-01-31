package cli

import (
	"fmt"

	"github.com/roger-russel/go-k8s-cli/internal/cmd"
)

func Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("some thing went wrong:", r)
		}
	}()

	cmd.Root()
}
