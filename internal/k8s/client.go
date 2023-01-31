package k8s

import (
	"context"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generage mockgen -source=$GOFILE -package=$GOPACKAGE -destination=mock_$GOFILE
type Client interface {
	CountPodsNumber() (int, error)
	CountNodesNumber() (int, error)
}

type ClientImpl struct {
	clientSet *kubernetes.Clientset
}

type Config struct {
	AuthType   AuthType
	Kubeconfig string
}

func NewClient(conf Config) (Client, error) {
	var (
		kconf *rest.Config
		err   error
	)

	switch conf.AuthType {
	case InClusterConfig:
		kconf, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get in cluster config: %v", err)
		}
	case BuildConfigFromFlags:
		kconf, err = buildConfigFromFlags(conf.Kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("failed to get kubernetes config: %v", err)
		}
	}

	clientSet, err := kubernetes.NewForConfig(kconf)

	if err != nil {
		return nil, fmt.Errorf("fail to create a new clientset with the kubeconfig: %w", err)
	}

	return &ClientImpl{
		clientSet,
	}, nil
}

func (c *ClientImpl) CountPodsNumber() (int, error) {
	pods, err := c.clientSet.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{})
	return len(pods.Items), err
}

func (c *ClientImpl) CountNodesNumber() (int, error) {
	nodes, err := c.clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	return len(nodes.Items), err
}

func buildConfigFromFlags(kconfig string) (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kconfig)

	if err != nil {
		return nil, fmt.Errorf("fail to build kubeconfig: %w", err)
	}

	return config, nil
}
