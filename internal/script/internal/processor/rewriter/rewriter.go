package rewriter

import "marcuson/scriptman/internal/script/internal/scan"

type Rewriter interface {
	RewriteBeforeAll() (string, error)
	RewriteLine(line *scan.LineScript) (string, error)
	RewriteAfterAll() (string, error)
}
