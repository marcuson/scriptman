package main

import (
	_ "embed"
	"marcuson/scriptman/internal/cliapp"
	"marcuson/scriptman/internal/logger"

	"os"
)

func main() {
	app := cliapp.InitCli()
	err := app.Run(os.Args)

	if err != nil {
		logger.Get().Fatal(err)
	}
}
