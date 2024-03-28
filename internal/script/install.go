package script

import (
	"fmt"
	"os"
	"path/filepath"

	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/fsext"
	"marcuson/scriptman/internal/utils/pathext"

	"github.com/adrg/xdg"
)

func Install(uri string) (*ScriptMetadata, error) {
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

