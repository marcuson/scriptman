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

func uninstallCmd(cCtx *cli.Context) error {
	id := cCtx.Args().Get(0)
	err := script.Unlink(id)
	if err != nil {
		return err
	}

	_, err = script.Uninstall(id)
	return err
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
		{
			Name:      "uninstall",
			Aliases:   []string{"u"},
			Usage:     "Uninstall a previously installed script (given its id).",
			Args:      true,
			ArgsUsage: "<id>",
			Action:    uninstallCmd,
		},
	}

	return cmds
}
