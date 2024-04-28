package script

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/pathext"
	"os"

	"github.com/adrg/xdg"
)

func Unlink(scriptId string) error {
	linkPath, _ := xdg.DataFile(config.BIN_HOME + "/" + scriptId)

	if pathext.Exists(linkPath) {
		err := os.Remove(linkPath)
		if err != nil {
			return err
		}
	}

	return nil
}
