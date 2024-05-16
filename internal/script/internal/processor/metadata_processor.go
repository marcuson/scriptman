package processor

import (
	"fmt"
	"marcuson/scriptman/internal/script/internal/scan"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/codeext"
	"marcuson/scriptman/internal/utils/pathext"
	"path"
)

type MetadataProcessor struct {
	scriptPath      string
	interpreterProc *InterpreterProcessor
	headOnly        bool
	isHeadFinished  bool
	_meta           scriptmeta.ScriptMetadata
}

func NewMetadataProcessor(scriptPath string) *MetadataProcessor {
	return &MetadataProcessor{
		scriptPath:      scriptPath,
		interpreterProc: NewInterpreterProcessor(),
		isHeadFinished:  false,
		_meta:           *scriptmeta.NewScriptMetadata(),
	}
}

func (obj *MetadataProcessor) SetHeadOnly(headOnly bool) *MetadataProcessor {
	obj.headOnly = headOnly
	return obj
}

func (obj *MetadataProcessor) Metadata() *scriptmeta.ScriptMetadata {
	return &obj._meta
}

func (obj *MetadataProcessor) ProcessStart() error {
	return obj.interpreterProc.ProcessStart()
}

func (obj *MetadataProcessor) ProcessLine(line *scan.LineScript) error {
	if obj.headOnly && obj.isHeadFinished {
		return nil
	}

	if line.LineIndex == 0 {
		err := obj.interpreterProc.ProcessLine(line)
		obj._meta.Interpreter = obj.interpreterProc.interpreter
		return err
	}

	if !line.IsEmpty && !line.IsComment {
		obj.isHeadFinished = true
	}

	if !line.IsMetadata {
		return nil
	}

	return obj.parseMetadata(line)
}

func (obj *MetadataProcessor) ProcessEnd() error {
	err := obj.interpreterProc.ProcessEnd()
	if err != nil {
		return err
	}

	if obj.scriptPath == "" {
		return nil
	}

	obj._meta.Ext = path.Ext(obj.scriptPath)

	codeext.SetIf(&obj._meta.Namespace, obj._meta.Namespace == "", "_nons_")
	codeext.SetIf(&obj._meta.Name, obj._meta.Name == "", pathext.Name(obj.scriptPath))
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
	case "asset":
		obj._meta.Assets = append(obj._meta.Assets, metaValue)
	case "getargs-tpl":
		obj._meta.GetargsTpl = metaValue
	case "sec:start":
		if obj.headOnly {
			obj.isHeadFinished = true
			return nil
		}
		obj._meta.GetOrAddSection(metaValue).LineStart = line.LineIndex
	case "sec:end":
		if obj.headOnly {
			obj.isHeadFinished = true
			return nil
		}
		obj._meta.GetOrAddSection(metaValue).LineEnd = line.LineIndex
	default:
		return fmt.Errorf("unknown meta key: %s", metaKey)
	}

	return nil
}

func (obj *MetadataProcessor) IsProcessCompletedEarly() bool {
	return obj.headOnly && obj.isHeadFinished
}
