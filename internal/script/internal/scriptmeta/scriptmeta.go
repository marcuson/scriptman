package scriptmeta

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/interpreter"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

type ScriptSection struct {
	LineStart int32
	LineEnd   int32
}

func NewScriptSection() *ScriptSection {
	sec := ScriptSection{
		LineStart: -1,
		LineEnd:   -1,
	}

	return &sec
}

type ScriptMetadata struct {
	Namespace   string
	Name        string
	Interpreter string
	Ext         string

	GetargsTpl string

	Sections map[string]*ScriptSection
	Assets   []string
}

func NewScriptMetadata() *ScriptMetadata {
	var meta ScriptMetadata
	meta.Sections = make(map[string]*ScriptSection)
	meta.Assets = []string{}
	return &meta
}

func (obj *ScriptMetadata) InterpreterInfo() interpreter.InterpreterInfo {
	return interpreter.Interpreters[obj.Interpreter]
}

func (obj *ScriptMetadata) GetOrAddSection(section string) *ScriptSection {
	sec, found := obj.Sections[section]
	if found {
		return sec
	}

	newSec := NewScriptSection()
	obj.Sections[section] = newSec
	return newSec
}

func (obj *ScriptMetadata) ScriptId() string {
	return obj.Namespace + "-" + obj.Name
}

func (obj *ScriptMetadata) InstallScriptIdDir() string {
	return obj.Namespace + "/" + obj.Name
}

func GetScriptPathFromId(id string) (bool, string) {
	idSplit := strings.Split(id, "-")
	ns := idSplit[0]
	name := strings.Join(idSplit[1:], "-")
	iDir, _ := xdg.SearchDataFile(config.SCRIPT_HOME + "/" + ns + "/" + name)
	installDir := filepath.ToSlash(iDir)

	files, _ := filepath.Glob(installDir + "/" + name + ".*")

	if len(files) <= 0 {
		return false, ""
	}

	return true, filepath.ToSlash(files[0])
}
