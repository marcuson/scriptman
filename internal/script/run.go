package script

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/run"
	"marcuson/scriptman/internal/utils/codeext"
)

func Run(idOrPath string, sections ...string) error {
	sec := codeext.Tern(len(sections) > 0, sections, []string{config.RUN_SECTION})
	secRewriter := rewriter.NewSecRewriter(sec...)
	_, err := run.RunWithHooks(idOrPath, run.RunHooks{}, secRewriter)
	return err
}
