package pathext

import (
	"errors"
	"os"
	"path"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}

func Name(fsPath string) string {
	basename := path.Base(fsPath)
	return basename[:len(basename)-len(path.Ext(basename))]
}
