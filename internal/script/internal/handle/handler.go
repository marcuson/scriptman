package handle

import (
	"io"
	"marcuson/scriptman/internal/script/internal/processor"
	"marcuson/scriptman/internal/script/internal/scan"
)

type Handler struct {
	scanner    *scan.Scanner
	processors []processor.Processor
}

func NewHandler(r io.Reader, processors ...processor.Processor) *Handler {
	return &Handler{
		scanner:    scan.NewScanner(r),
		processors: processors,
	}
}

func (obj *Handler) Handle() error {
	for _, p := range obj.processors {
		if err := p.ProcessStart(); err != nil {
			return err
		}
	}

	for obj.scanner.Scan() {
		if err := obj.scanner.Err(); err != nil {
			return err
		}

		for _, p := range obj.processors {
			if err := p.ProcessLine(obj.scanner.Line()); err != nil {
				return err
			}
		}
	}

	for _, p := range obj.processors {
		if err := p.ProcessEnd(); err != nil {
			return err
		}
	}

	return nil
}
