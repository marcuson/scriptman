package processor

import (
	"marcuson/scriptman/internal/script/internal/scan"
)

type Processor interface {
	ProcessStart() error
	ProcessLine(line *scan.LineScript) error
	ProcessEnd() error
}
