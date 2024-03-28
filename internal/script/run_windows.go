//go:build windows

package script

import (
	"path/filepath"
	"strings"
)

func getNormalizedScriptPath(scriptPath string, interpreter string) string {
	var scriptPathNormalized string
	switch interpreter {
	case "bash":
		scriptPathNormalized = strings.Replace(filepath.ToSlash(scriptPath), "C:", "/mnt/c", 1)
	default:
		scriptPathNormalized = scriptPath
	}

	return scriptPathNormalized
}
