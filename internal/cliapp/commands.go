package cliapp

import (
	"fmt"
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script"
	"marcuson/scriptman/internal/utils/codeext"

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

func listCmd(cCtx *cli.Context) error {
	scriptsMeta, err := script.GetInstalledList()
	if err != nil {
		return err
	}

	for _, m := range scriptsMeta {
		fmt.Println(m.ScriptId())
	}

	return nil
}

func runCmd(cCtx *cli.Context) error {
	idOrPath := cCtx.Args().Get(0)
	sections := cCtx.StringSlice("section")
	sections = codeext.Tern(len(sections) > 0, sections, []string{config.RUN_SECTION})
	return script.Run(idOrPath, sections...)
}

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
		{
			Name:    "list",
			Usage:   "List installed scripts.",
			Aliases: []string{"ls"},
			Action:  listCmd,
		},
		{
			Name: "run",
			Usage: "Run a script previsouly installed (given its unique <scriptId>) or " +
				"directly from filesystem (given its path).",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:    "section",
					Aliases: []string{"s"},
					Usage:   `Sections to run, default to "run"`,
				},
			},
			Args:      true,
			ArgsUsage: "<script id or path>",
			Action:    runCmd,
		},
	}

	return cmds
}
