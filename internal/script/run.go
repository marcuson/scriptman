package script

import (
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/run"
)

func Run(idOrPath string, sections ...string) error {
	secRewriter := rewriter.NewSecRewriter(sections...)
	_, err := run.RunWithHooks(idOrPath, run.RunHooks{}, secRewriter)
	return err
}
