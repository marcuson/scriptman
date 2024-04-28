package script

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/httpext"
	"os"
	"path"
	"strings"

	"github.com/adrg/xdg"
)

type httpScriptInstaller struct{}

func (obj *httpScriptInstaller) prepare(ctx *scriptInstallCtx) error {
	tmpInstallPath, err := xdg.DataFile(config.TMP_HTTP_INSTALL_DIR + "/" +
		path.Base(ctx.ScriptUri.Path))
	if err != nil {
		return err
	}

	err = httpext.DownloadFile(ctx.RawUri, tmpInstallPath)
	if err != nil {
		return err
	}

	ctx.InstallFromLocalFile = tmpInstallPath

	return nil
}

func (obj *httpScriptInstaller) installMainFile(ctx *scriptInstallCtx) error {
	err := os.MkdirAll(ctx.InstallTargetDir, 0777) // FIXME: perm
	if err != nil {
		return err
	}

	err = os.Rename(ctx.InstallFromLocalFile, ctx.InstallTargetMainFile)
	return err
}

type httpAssetInstaller struct{}

func (obj *httpAssetInstaller) installAsset(ctx *assetInstallCtx) error {
	downloadUri := ctx.RawUri
	if ctx.IsRelative() {
		rawSplit := strings.Split(ctx.ScriptInstallCtx.RawUri, "/")
		rawSplit = rawSplit[:len(rawSplit)-1]
		downloadUri = strings.Join(rawSplit, "/") + "/" + ctx.RawUri
	}

	err := os.MkdirAll(path.Dir(ctx.InstallTargetFile), 0777) // FIXME: perm
	if err != nil {
		return err
	}

	err = httpext.DownloadFile(downloadUri, ctx.InstallTargetFile)
	return err
}
