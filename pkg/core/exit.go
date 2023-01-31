package core

import "os"

var fakeExit *bool

func Exit(i int) {
	if fakeExit == nil {
		fe := os.Getenv("FAKE_EXIT") != ""
		fakeExit = &fe
	}

	if !*fakeExit {
		os.Exit(i)
	}

	panic("fake exit called")
}
