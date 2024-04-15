package script

import (
	"fmt"
	"os"
	"path/filepath"

	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/fsext"
	"marcuson/scriptman/internal/utils/pathext"

	"github.com/adrg/xdg"
)

func Install(uri string) (*scriptmeta.ScriptMetadata, error) {
	if !pathext.Exists(uri) {
		return nil, fmt.Errorf("script not found at '%s'", uri)
	}

	meta, err := ParseMetadata(uri)
	if err != nil {
		return nil, err
	}

	installPath, _ := xdg.DataFile(config.SCRIPT_HOME_DEFAULT + "/" +
		meta.InstallScriptIdDir() + "/" + meta.Name + meta.Ext)

	_, err = fsext.CopyFile(uri, installPath)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func Uninstall(id string) (*scriptmeta.ScriptMetadata, error) {
	found, scriptPath := scriptmeta.GetScriptPathFromId(id)
	if !found {
		return nil, fmt.Errorf("unable to find script with id '%s' for uninstall", id)
	}

	meta, err := ParseMetadata(scriptPath)
	if err != nil {
		return nil, err
	}

	iDir, _ := xdg.DataFile(config.SCRIPT_HOME_DEFAULT + "/" + meta.InstallScriptIdDir())
	installDir := filepath.ToSlash(iDir)

	err = os.RemoveAll(installDir)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
