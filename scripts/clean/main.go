package main

import (
	"os"
)

var (
	dirs = []string{"_build"}
)

func main() {
	var err error

	for _, dir := range dirs {
		err = os.RemoveAll(dir)

		if err != nil {
			panic(err)
		}
	}

	os.Exit(0)
}
