package interpreter

import (
	"path/filepath"
	"strings"
)

func (obj *bashInterpreter) GetNormalizedScriptPath(scriptPath string) string {
	return strings.Replace(filepath.ToSlash(scriptPath), "C:", "/mnt/c", 1)
}
