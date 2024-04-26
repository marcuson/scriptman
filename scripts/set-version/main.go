package main

import (
	"log"
	"os"
)

func main() {
	ver := os.Args[1]
	envData := "SCRIPTMAN_VERSION=" + ver

	err := os.MkdirAll("_build", 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("_build/build.env", []byte(envData), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
