package script

import (
	"fmt"
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/run"
	"marcuson/scriptman/internal/script/internal/scriptutils"
	"os"
)

type RunOptions struct {
	ScriptIdOrPath string
	EnvFilePath    string
}

func NewRunOpts(idOrPath string) *RunOptions {
	opts := &RunOptions{
		ScriptIdOrPath: idOrPath,
	}
	return opts
}

func Run(opts *RunOptions, sections ...string) error {
	found, scriptPath := scriptutils.FindScriptPath(opts.ScriptIdOrPath)
	if !found {
		return fmt.Errorf("cannot find script '%s' by id or path", opts.ScriptIdOrPath)
	}

	inter, err := ParseInterpreter(scriptPath)
	if err != nil {
		return err
	}

	rewriters := []rewriter.Rewriter{rewriter.NewSecRewriter(sections...)}
	if opts.EnvFilePath != "" {
		file, err := os.Open(opts.EnvFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		rewriters = append(rewriters, rewriter.NewDotenvInjectorRewriter(file, inter))
	}
	_, err = run.RunWithHooks(scriptPath, run.RunHooks{}, rewriters...)
	return err
}
