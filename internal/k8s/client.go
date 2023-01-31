package k8s

import (
	"context"
	"fmt"

	"github.com/roger-russel/go-k8s-cli/pkg/core"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client interface {
	CountPodsNumber() (int, error)
	CountNodesNumber() (int, error)
}

type ClientImpl struct {
	clientSet *kubernetes.Clientset
}

type Config struct {
	Type string
}

func NewClient(conf *rest.Config) Client {
	clientSet, err := kubernetes.NewForConfig(conf)

	if err != nil {
		core.Out(
			fmt.Errorf("fail to create a new clientset with the kubeconfig: %w", err),
		)
		core.Exit(1)
	}

	return &ClientImpl{
		clientSet,
	}
}

func (c *ClientImpl) CountPodsNumber() (int, error) {
	pods, err := c.clientSet.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{})
	return len(pods.Items), err
}

func (c *ClientImpl) CountNodesNumber() (int, error) {
	nodes, err := c.clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	return len(nodes.Items), err
}
