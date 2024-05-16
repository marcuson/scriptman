package script

import (
	"marcuson/scriptman/internal/script/internal/handle"
	"marcuson/scriptman/internal/script/internal/processor"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"os"
)

func parseMetadata(path string, headOnly bool) (*scriptmeta.ScriptMetadata, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	metaProcessor := processor.NewMetadataProcessor(path)
	metaProcessor.SetHeadOnly(headOnly)
	handler := handle.NewHandler(file, metaProcessor)
	if err = handler.Handle(); err != nil {
		return nil, err
	}

	return metaProcessor.Metadata(), nil
}

func ParseMetadataHeaderOnly(path string) (*scriptmeta.ScriptMetadata, error) {
	return parseMetadata(path, true)
}

func ParseMetadata(path string) (*scriptmeta.ScriptMetadata, error) {
	return parseMetadata(path, false)
}

func ParseInterpreter(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	handler := handle.NewHandler(file)
	if err = handler.Handle(); err != nil {
		return "", err
	}

	return handler.Interpreter(), nil
}
