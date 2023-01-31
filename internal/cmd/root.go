package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/roger-russel/go-k8s-cli/pkg/core"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

// Root Command
func Root() {
	var read string
	var kconfig string
	var readEnum ReadEnum

	rootCmd := &cobra.Command{
		Use: "go-k8s-cli",
		Run: run,
	}

	rootCmd.PersistentFlags().StringVarP(
		&read, "read", "r", "",
		fmt.Sprintf("read the number of %v: -r pods", readEnum.Values()),
	)

	if err := rootCmd.MarkPersistentFlagRequired("read"); err != nil {
		core.Out("fail to mark read as required flag error:", err)
		core.Exit(1)
	}

	kubeconfigFlag(rootCmd, &kconfig)

	rootCmd.Execute()
}

func kubeconfigFlag(rootCmd *cobra.Command, kconfig *string) {
	var message, defaultValue string

	if home := homedir.HomeDir(); home != "" {
		defaultValue = filepath.Join(home, ".kube", "config")
		message = "kubeconfig " + defaultValue + " (optional) absolute path to the kubeconfig file"
	} else {
		message = "kubeconfig absolute path to the kubeconfig file"
	}

	rootCmd.PersistentFlags().StringVarP(
		kconfig, "kubeconfig", "c", defaultValue,
		message,
	)
}
