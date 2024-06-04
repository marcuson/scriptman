package cliapp

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	app *cli.App
)

func InitCli() *cli.App {
	asciiArt := `
	.---------------------------------------------------------------.
	|   ____               _         _                              |
	|  / ___|   ___  _ __ (_) _ __  | |_  _ __ ___    __ _  _ __    |
	|  \___ \  / __|| '__|| || '_ \ | __|| '_ ' _ \  / _' || '_ \   |
	|   ___) || (__ | |   | || |_) || |_ | | | | | || (_| || | | |  |
	|  |____/  \___||_|   |_|| .__/  \__||_| |_| |_| \__,_||_| |_|  |
	|                        |_|                                    |
	'---------------------------------------------------------------'`

	app = &cli.App{
		Name:    "sman",
		Usage:   "Scriptman: a cross-platform script manager written in Go.",
		Version: Version,
	}

	app.CustomAppHelpTemplate = fmt.Sprintf(
		"%s\n\n%s",
		asciiArt,
		cli.AppHelpTemplate,
	)

	app.Commands = getCmds()

	return app
}
