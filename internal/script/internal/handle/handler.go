package handle

import (
	"io"
	"marcuson/scriptman/internal/script/internal/processor"
	"marcuson/scriptman/internal/script/internal/scan"
	"marcuson/scriptman/internal/utils/slicesext"
)

type Handler struct {
	scanner              *scan.Scanner
	interpreterProcessor *processor.InterpreterProcessor
	processors           []processor.Processor
}

func NewHandler(r io.Reader, processors ...processor.Processor) *Handler {
	obj := &Handler{
		scanner:              scan.NewScanner(r),
		interpreterProcessor: &processor.InterpreterProcessor{},
	}
	obj.processors = []processor.Processor{obj.interpreterProcessor}
	obj.processors = append(obj.processors, processors...)
	return obj
}

func (obj *Handler) Interpreter() string {
	return obj.interpreterProcessor.Interpreter()
}

func (obj *Handler) handleLine() (error, bool) {
	if err := obj.scanner.Err(); err != nil {
		return err, false
	}

	for _, p := range obj.processors {
		if err := p.ProcessLine(obj.scanner.Line()); err != nil {
			return err, false
		}
	}

	shouldExitEarly := slicesext.Reduce(
		obj.processors,
		func(acc bool, p processor.Processor) bool {
			return acc && p.IsProcessCompletedEarly()
		},
		true)
	return nil, shouldExitEarly
}

func (obj *Handler) Handle() error {
	for _, p := range obj.processors {
		if err := p.ProcessStart(); err != nil {
			return err
		}
	}

	// First line (for interpreter hint)
	obj.scanner.Scan()
	err, shouldExitEarly := obj.handleLine()
	if err != nil {
		return err
	}

	obj.scanner.SetInterpreter(obj.interpreterProcessor.Interpreter())
	if shouldExitEarly {
		return nil
	}

	// Second line and on
	for obj.scanner.Scan() {
		err, shouldExitEarly := obj.handleLine()
		if err != nil {
			return err
		}
		if shouldExitEarly {
			break
		}
	}
	if obj.scanner.Err() != nil {
		return err
	}

	for _, p := range obj.processors {
		if err := p.ProcessEnd(); err != nil {
			return err
		}
	}

	return nil
}
