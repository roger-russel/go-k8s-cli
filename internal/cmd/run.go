package cmd

import (
	"context"
	"fmt"

	"github.com/roger-russel/go-k8s-cli/pkg/core"
	"github.com/spf13/cobra"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func run(cmd *cobra.Command, args []string) {
	kcore := setupClientset(cmd).CoreV1()
	action := parseActionFlag(cmd)
	command(action, kcore)
}

func setupClientset(cmd *cobra.Command) *kubernetes.Clientset {
	kconfig, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		core.Out("failed to get kubeconfig flag error:", err)
		core.Exit(1)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kconfig)

	if err != nil {
		core.Out(
			fmt.Errorf("fail to build kubeconfig: %w", err),
		)
		core.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		core.Out(
			fmt.Errorf("fail to create a new clientset with the kubeconfig: %w", err),
		)
		core.Exit(1)
	}

	return clientset
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

func command(action string, kcore v1.CoreV1Interface) {
	var (
		total    int
		err      error
		pods     *coreV1.PodList
		nodes    *coreV1.NodeList
		readEnum ReadEnum
	)

	switch action {
	case "pods":
		pods, err = kcore.Pods("").List(context.TODO(), metav1.ListOptions{})
		total = len(pods.Items)
	case "nodes":
		nodes, err = kcore.Nodes().List(context.TODO(), metav1.ListOptions{})
		total = len(nodes.Items)
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
