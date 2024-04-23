package processor

import (
	"fmt"
	"marcuson/scriptman/internal/script/internal/scan"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/codeext"
	"marcuson/scriptman/internal/utils/pathext"
	"path"
	"strings"
)

type MetadataProcessor struct {
	scriptPath string
	_meta      scriptmeta.ScriptMetadata
}

func NewMetadataProcessor(scriptPath string) *MetadataProcessor {
	return &MetadataProcessor{
		scriptPath: scriptPath,
		_meta:      *scriptmeta.NewScriptMetadata(),
	}
}

func (obj *MetadataProcessor) Metadata() *scriptmeta.ScriptMetadata {
	return &obj._meta
}

func (obj *MetadataProcessor) ProcessStart() error {
	return nil
}

func (obj *MetadataProcessor) ProcessLine(line *scan.LineScript) error {
	switch {
	case line.IsShebang:
		return obj.parseShebang(line)
	case line.IsMetadata:
		return obj.parseMetadata(line)
	default:
		return nil
	}
}

func (obj *MetadataProcessor) ProcessEnd() error {
	if obj.scriptPath == "" {
		return nil
	}

	obj._meta.Ext = path.Ext(obj.scriptPath)

	codeext.SetIf(&obj._meta.Namespace, obj._meta.Namespace == "", "_nons_")
	codeext.SetIf(&obj._meta.Name, obj._meta.Name == "", pathext.Name(obj.scriptPath))
	return nil
}

func (obj *MetadataProcessor) parseShebang(line *scan.LineScript) error {
	lineSplit := line.LineSplit()
	interpreter := lineSplit[len(lineSplit)-1]
	interpreter = strings.Replace(interpreter, "#!", "", 1)
	obj._meta.Interpreter = interpreter
	return nil
}

func (obj *MetadataProcessor) parseMetadata(line *scan.LineScript) error {
	lineSplit := line.LineSplit()

	metaKey := lineSplit[2]
	metaValue := lineSplit[3]

	switch metaKey {
	case "namespace":
		obj._meta.Namespace = metaValue
	case "name":
		obj._meta.Name = metaValue
	case "interpreter":
		if obj._meta.Interpreter == "" {
			obj._meta.Interpreter = metaValue
		}
	case "sec:start":
		obj._meta.GetOrAddSection(metaValue).LineStart = line.LineIndex
	case "sec:end":
		obj._meta.GetOrAddSection(metaValue).LineEnd = line.LineIndex
	case "asset":
		obj._meta.Assets = append(obj._meta.Assets, metaValue)
	default:
		return fmt.Errorf("unknown meta key: %s", metaKey)
	}

	return nil
}

func (obj *MetadataProcessor) IsProcessCompletedEarly() bool {
	return false
}
