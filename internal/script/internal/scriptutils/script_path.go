package scriptutils

import (
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/pathext"
)

func FindScriptPath(idOrPath string) (bool, string) {
	instFound, scriptPath := scriptmeta.GetScriptPathFromId(idOrPath)
	if !instFound {
		scriptPath = idOrPath
	}
	if !pathext.Exists(scriptPath) {
		return false, ""
	}

	return true, scriptPath
}
