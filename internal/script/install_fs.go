package script

import (
	"marcuson/scriptman/internal/utils/fsext"
	"os"
	"path"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

type fsScriptInstaller struct{}

func (obj *fsScriptInstaller) prepare(ctx *scriptInstallCtx) error {
	ctx.InstallFromLocalFile = ctx.RawUri
	return nil
}

func (obj *fsScriptInstaller) installMainFile(ctx *scriptInstallCtx) error {
	_, err := fsext.CopyFile(ctx.InstallFromLocalFile, ctx.InstallTargetMainFile)
	return err
}

type fsAssetInstaller struct{}

func (obj *fsAssetInstaller) installAsset(ctx *assetInstallCtx) error {
	var err error
	if ctx.IsRelative() {
		installFromLocalDir := path.Dir(ctx.ScriptInstallCtx.InstallFromLocalFile)
		err = obj.installFromGlob(installFromLocalDir, ctx.RawUri, ctx.ScriptInstallCtx.InstallTargetDir)
	} else {
		rawUriSplit := strings.Split(ctx.AssetUri.Path, "|>")
		var basePath string
		var glob string
		if len(rawUriSplit) > 1 {
			basePath = rawUriSplit[0]
			glob = rawUriSplit[1]
		} else {
			basePath = path.Dir(ctx.AssetUri.Path)
			glob = path.Base(ctx.AssetUri.Path)
		}
		err = obj.installFromGlob(basePath, glob, ctx.ScriptInstallCtx.InstallTargetDir)
	}

	return err
}

func (obj *fsAssetInstaller) installFromGlob(
	basePath string, glob string, scriptTargetDir string) error {

	installFromLocalDir := basePath
	installFromDirFSys := os.DirFS(installFromLocalDir)

	files, err := doublestar.Glob(installFromDirFSys, glob, doublestar.WithFilesOnly())
	if err != nil {
		return err
	}

	for _, f := range files {
		fInstallFromPath := installFromLocalDir + "/" + f
		fInstallPath := scriptTargetDir + "/" + f
		_, err = fsext.CopyFile(fInstallFromPath, fInstallPath)
		if err != nil {
			return err
		}
	}

	return nil
}
