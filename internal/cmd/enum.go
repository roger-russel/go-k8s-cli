package cmd

import (
	"fmt"
)

type ReadEnum string

const (
	Pods  ReadEnum = "pods"
	Nodes ReadEnum = "nodes"
)

func (r *ReadEnum) String() string {
	return string(*r)
}

func (r *ReadEnum) Type() string {
	return "enum"
}

func (ReadEnum) Values() []string {
	return []string{string(Pods), string(Nodes)}
}

func (r *ReadEnum) Set(value string) error {
	if value != string(Pods) && value != string(Nodes) {
		return fmt.Errorf("invalid resource type %q, must be one of %v", value, r.Values())
	}
	*r = ReadEnum(value)
	return nil
}
