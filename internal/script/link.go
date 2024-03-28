package script

import (
	"fmt"
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/pathext"
	"os"

	"github.com/adrg/xdg"
)

func Link(scriptId string) error {
	found, scriptPath := GetScriptPathFromId(scriptId)
	if !found {
		return fmt.Errorf("unable to find script with id '%s' for link", scriptId)
	}

	linkPath, _ := xdg.DataFile(config.BIN_HOME_DEFAULT + "/" + scriptId)
	if pathext.Exists(linkPath) {
		err := os.Remove(linkPath)
		if err != nil {
			return err
		}
	}

	err := os.Symlink(scriptPath, linkPath)
	if err != nil {
		return err
	}

	return nil
}
