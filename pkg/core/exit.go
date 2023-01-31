package core

import "os"

var fakeExit bool = false

func Exit(i int) {
	if !fakeExit {
		os.Exit(i)
	}

	panic("fake exit called")
}
