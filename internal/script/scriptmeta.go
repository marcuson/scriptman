package script

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/codeext"
	"marcuson/scriptman/internal/utils/pathext"
	"path"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

type ScriptMetadata struct {
	Namespace   string
	Name        string
	Interpreter string
	Ext         string
}

func (obj *ScriptMetadata) FillMissingMetadata(scriptPath string) {
	obj.Ext = path.Ext(scriptPath)

	codeext.SetIf(&obj.Namespace, obj.Namespace == "", "_nons_")
	codeext.SetIf(&obj.Name, obj.Name == "", pathext.Name(scriptPath))
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
	iDir, _ := xdg.DataFile(config.SCRIPT_HOME_DEFAULT + "/" + ns + "/" + name)
	installDir := filepath.ToSlash(iDir)

	files, _ := filepath.Glob(installDir + "/" + name + ".*")

	if len(files) <= 0 {
		return false, ""
	}

	return true, filepath.ToSlash(files[0])
}
