package script

import (
	"marcuson/scriptman/internal/utils/fsext"
	"os"
	"path"

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
		err = obj.installFromRelGlob(ctx)
	} else {
		err = obj.installFromAbsPath(ctx)
	}

	if err != nil {
		return err
	}

	return nil
}

func (obj *fsAssetInstaller) installFromRelGlob(ctx *assetInstallCtx) error {
	installFromLocalDir := path.Dir(ctx.ScriptInstallCtx.InstallFromLocalFile)
	installFromDirFSys := os.DirFS(installFromLocalDir)

	files, err := doublestar.Glob(installFromDirFSys, ctx.RawUri, doublestar.WithFilesOnly())
	if err != nil {
		return err
	}

	for _, f := range files {
		fInstallFromPath := installFromLocalDir + "/" + f
		fInstallPath := ctx.ScriptInstallCtx.InstallTargetDir + "/" + f
		_, err = fsext.CopyFile(fInstallFromPath, fInstallPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (obj *fsAssetInstaller) installFromAbsPath(ctx *assetInstallCtx) error {
	_, err := fsext.CopyFile(ctx.RawUri, ctx.InstallTargetFile)
	return err
}
