package script

import (
	"cmp"
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/pathext"
	"marcuson/scriptman/internal/utils/slicesext"
	"path/filepath"
	"slices"
	"strings"

	"github.com/adrg/xdg"
)

func GetInstalledList() ([]*ScriptMetadata, error) {
	scriptsHome, _ := xdg.DataFile(config.SCRIPT_HOME_DEFAULT)
	scriptFilepaths, err := filepath.Glob(scriptsHome + "/*/*/*.*")
	if err != nil {
		return nil, err
	}

	scriptFilepaths = slicesext.Map(
		scriptFilepaths,
		func(a string) string { return filepath.ToSlash(a) },
	)
	scriptFilepaths = slices.DeleteFunc(scriptFilepaths, func(a string) bool {
		pathSplit := strings.Split(a, "/")
		dirName := pathSplit[len(pathSplit)-2]
		fileName := pathSplit[len(pathSplit)-1]
		fileNameNoExt := pathext.Name(fileName)
		return dirName != fileNameNoExt
	})
	scriptsMeta := []*ScriptMetadata{}

	for _, fPath := range scriptFilepaths {
		meta, err := ParseMetadata(fPath)
		if err != nil {
			return nil, err
		}
		scriptsMeta = append(scriptsMeta, meta)
	}

	scriptMetaComp := func(a, b *ScriptMetadata) int {
		return cmp.Compare(a.ScriptId(), b.ScriptId())
	}

	slices.SortFunc(scriptsMeta, scriptMetaComp)

	return scriptsMeta, nil
}