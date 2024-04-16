package processor

import (
	"marcuson/scriptman/internal/script/internal/scan"
)

type InterpreterProcessor struct {
	metaProcessor *MetadataProcessor
}

func NewInterpreterProcessor() *InterpreterProcessor {
	return &InterpreterProcessor{
		metaProcessor: NewMetadataProcessor(""),
	}
}

func (obj *InterpreterProcessor) Interpreter() string {
	return obj.metaProcessor.Metadata().Interpreter
}

func (obj *InterpreterProcessor) ProcessStart() error {
	return nil
}

func (obj *InterpreterProcessor) ProcessLine(line *scan.LineScript) error {
	return obj.metaProcessor.ProcessLine(line)
}

func (obj *InterpreterProcessor) ProcessEnd() error {
	return obj.metaProcessor.ProcessEnd()
}

func (obj *InterpreterProcessor) IsProcessCompletedEarly() bool {
	return obj.Interpreter() != ""
}
