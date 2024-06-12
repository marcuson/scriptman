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

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "preserve-user",
			Aliases: []string{"p"},
			Usage: "If Scriptman is run with sudo and this flag is set to true, all operations " +
				"will be performed on the data dirs of the user which invoked the sudo command. " +
				"If set to false, operations will be performed on root user's data dirs.",
		},
	}
	app.Commands = getCmds()

	return app
}
