package script

import (
	"fmt"
	"os"
	"path"

	"marcuson/scriptman/internal/script/internal/scriptmeta"
)

func Uninstall(id string) error {
	found, scriptPath := scriptmeta.GetScriptPathFromId(id)
	if !found {
		return fmt.Errorf("unable to find script with id '%s' for uninstall", id)
	}

	installDir := path.Dir(scriptPath)

	err := os.RemoveAll(installDir)
	if err != nil {
		return err
	}

	return nil
}
