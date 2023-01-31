package core

import (
	"os"
	"testing"
)

func TestExit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != "fake exit called" {
				t.Errorf("unexpected panic: %v", r)
			}
		} else {
			t.Error("Exit did not panic")
		}
	}()

	os.Setenv("FAKE_EXIT", "true")
	Exit(1)
}
