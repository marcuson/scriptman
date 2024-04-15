package script

import (
	"marcuson/scriptman/internal/script/internal/handle"
	"marcuson/scriptman/internal/script/internal/processor"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"os"
)

func ParseMetadata(path string) (*scriptmeta.ScriptMetadata, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	metaProcessor := processor.NewMetadataProcessor(path)
	handler := handle.NewHandler(file, metaProcessor)
	if err = handler.Handle(); err != nil {
		return nil, err
	}

	return metaProcessor.Metadata(), nil
}
