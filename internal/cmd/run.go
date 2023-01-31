package cmd

import (
	"fmt"

	"github.com/roger-russel/go-k8s-cli/internal/k8s"
	"github.com/roger-russel/go-k8s-cli/pkg/core"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	kcli := setupClientset(cmd)
	action := parseActionFlag(cmd)
	command(action, kcli)
}

func setupClientset(cmd *cobra.Command) k8s.Client {
	kconfig, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		core.Out("failed to get kubeconfig flag error:", err)
		core.Exit(1)
	}

	cli, err := k8s.NewClient(
		k8s.Config{
			AuthType:   k8s.BuildConfigFromFlags,
			Kubeconfig: kconfig,
		})

	if err != nil {
		core.Out(
			fmt.Errorf("fail to create a new clientset with the kubeconfig: %w", err),
		)
		core.Exit(1)
	}

	return cli
}

func parseActionFlag(cmd *cobra.Command) string {
	readCommand, err := cmd.Flags().GetString("read")
	if err != nil {
		core.Out("failed to get read flag error:", err)
		core.Exit(1)
	}

	var readEnum ReadEnum
	if err := readEnum.Set(readCommand); err != nil {
		core.Out("invalid -r flag given:", err)
		core.Exit(1)
	}

	return readEnum.String()
}

func command(action string, kcli k8s.Client) {
	var (
		total    int
		err      error
		readEnum ReadEnum
	)

	switch action {
	case "pods":
		total, err = kcli.CountPodsNumber()
	case "nodes":
		total, err = kcli.CountNodesNumber()
	default:
		core.Out("unknown command")
		core.Exit(1)
	}

	if err != nil {
		core.Out(
			fmt.Errorf("fail to get list of %v error: %w", readEnum.String(), err),
		)
		core.Exit(1)
	}

	core.Out(total)
}
