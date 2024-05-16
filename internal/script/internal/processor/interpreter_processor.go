package processor

import (
	"marcuson/scriptman/internal/script/internal/scan"
	"strings"
)

type InterpreterProcessor struct {
	isOverLine0 bool
	interpreter string
}

func NewInterpreterProcessor() *InterpreterProcessor {
	return &InterpreterProcessor{
		isOverLine0: false,
	}
}

func (obj *InterpreterProcessor) Interpreter() string {
	return obj.interpreter
}

func (obj *InterpreterProcessor) ProcessStart() error {
	return nil
}

func (obj *InterpreterProcessor) ProcessLine(line *scan.LineScript) error {
	if line.LineIndex > 0 {
		obj.isOverLine0 = true
		return nil
	}

	switch {
	case line.IsShebang:
		lineSplit := line.LineSplit()
		inter := lineSplit[len(lineSplit)-1]
		inter = strings.Replace(inter, "#!", "", 1)
		obj.interpreter = inter
	case line.IsMetadata:
		lineSplit := line.LineSplit()
		metaKey := lineSplit[2]
		metaValue := lineSplit[3]
		if metaKey == "interpreter" {
			obj.interpreter = metaValue
		}
	}

	return nil
}

func (obj *InterpreterProcessor) ProcessEnd() error {
	return nil
}

func (obj *InterpreterProcessor) IsProcessCompletedEarly() bool {
	return obj.isOverLine0 || obj.Interpreter() != ""
}
