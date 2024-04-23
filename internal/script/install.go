package script

import (
	"fmt"
	"os"
	"path"

	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/fsext"
	"marcuson/scriptman/internal/utils/pathext"

	"github.com/adrg/xdg"
	"github.com/bmatcuk/doublestar/v4"
)

func Install(uri string) (*scriptmeta.ScriptMetadata, error) {
	installFromLocalPath := uri
	installFromLocalDir := path.Dir(installFromLocalPath)
	if !pathext.Exists(installFromLocalPath) {
		return nil, fmt.Errorf("script not found at '%s'", installFromLocalPath)
	}

	meta, err := ParseMetadata(installFromLocalPath)
	if err != nil {
		return nil, err
	}

	installDir, err := xdg.DataFile(config.SCRIPT_HOME_DEFAULT + "/" +
		meta.InstallScriptIdDir())
	if err != nil {
		return nil, err
	}
	installPath := installDir + "/" + meta.Name + meta.Ext

	_, err = fsext.CopyFile(installFromLocalPath, installPath)
	if err != nil {
		return nil, err
	}

	installFromDirFSys := os.DirFS(installFromLocalDir)
	for _, assetGlob := range meta.Assets {
		files, err := doublestar.Glob(installFromDirFSys, assetGlob, doublestar.WithFilesOnly())
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			fInstallFromPath := installFromLocalDir + "/" + f
			fInstallPath := installDir + "/" + f
			_, err = fsext.CopyFile(fInstallFromPath, fInstallPath)
			if err != nil {
				return nil, err
			}
		}
	}

	return meta, nil
}

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
