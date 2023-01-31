package k8s

import (
	"fmt"
)

type AuthType string

const (
	InClusterConfig      AuthType = "inClusterConfig"
	BuildConfigFromFlags AuthType = "buildConfigFromFlags"
)

func (a *AuthType) String() string {
	return string(*a)
}

func (a *AuthType) Type() string {
	return "enum"
}

func (AuthType) Values() []string {
	return []string{string(InClusterConfig), string(BuildConfigFromFlags)}
}

func (a *AuthType) Set(value string) error {
	if value != string(InClusterConfig) && value != string(BuildConfigFromFlags) {
		return fmt.Errorf("invalid resource type %q, must be one of %v", value, a.Values())
	}
	*a = AuthType(value)
	return nil
}
