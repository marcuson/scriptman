package cliapp

import (
	"fmt"
	"marcuson/scriptman/internal/script"

	"github.com/urfave/cli/v2"
)

func installCmd(cCtx *cli.Context) error {
	scriptPath := cCtx.Args().Get(0)
	meta, err := script.Install(scriptPath)
	if err != nil {
		return err
	}

	return script.Link(meta.ScriptId())
}

func getCmds() []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:      "install",
			Aliases:   []string{"i"},
			Usage:     "Install a script from filesystem (given its path).",
			Args:      true,
			ArgsUsage: "<path>",
			Action:    installCmd,
		},
	}

	return cmds
}
