package script

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/fsext"
	"marcuson/scriptman/internal/utils/httpext"
	"marcuson/scriptman/internal/utils/pathext"

	"github.com/adrg/xdg"
	"github.com/bmatcuk/doublestar/v4"
)

type installCtx struct {
	RawUri    string
	ScriptUri *url.URL

	InstallFromLocalPath string
	Meta                 *scriptmeta.ScriptMetadata
	InstallScriptPath    string
	InstallScriptDir     string
}

func newInstallCtx(uri string) *installCtx {
	ctx := &installCtx{
		RawUri: uri,
	}
	return ctx
}

func (obj *installCtx) prepareInstallation() error {
	uriParsed, err := url.Parse(obj.RawUri)
	if err != nil {
		return err
	}

	if uriParsed.Scheme == "" {
		uriParsed.Scheme = "file"
	}

	obj.ScriptUri = uriParsed

	switch obj.ScriptUri.Scheme {
	case "file":
		obj.InstallFromLocalPath = obj.RawUri
	case "http":
	case "https":
		tmpInstallPath, err := xdg.DataFile(config.TMP_ROOT_DIR + "/" +
			path.Base(obj.ScriptUri.Path))
		if err != nil {
			return err
		}

		err = httpext.DownloadFile(obj.RawUri, tmpInstallPath)
		if err != nil {
			return err
		}

		obj.InstallFromLocalPath = tmpInstallPath
	}

	return nil
}

func installMainFile(ctx *installCtx) error {
	installDir, err := xdg.DataFile(config.SCRIPT_HOME + "/" +
		ctx.Meta.InstallScriptIdDir())
	if err != nil {
		return err
	}
	ctx.InstallScriptDir = installDir
	ctx.InstallScriptPath = installDir + "/" + ctx.Meta.Name + ctx.Meta.Ext

	switch ctx.ScriptUri.Scheme {
	case "file":
		_, err = fsext.CopyFile(ctx.InstallFromLocalPath, ctx.InstallScriptPath)
		if err != nil {
			return err
		}
	case "http":
	case "https":
		err = os.MkdirAll(ctx.InstallScriptDir, 0777) // FIXME: perm
		if err != nil {
			return err
		}

		err = os.Rename(ctx.InstallFromLocalPath, ctx.InstallScriptPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func installAsset(assetUri string, ctx *installCtx) error {
	absUri := assetUri
	uriParsed, err := url.Parse(assetUri)
	if err != nil {
		return err
	}

	if uriParsed.Scheme == "" {
		absUri = path.Dir(ctx.RawUri) + "/" + assetUri // FIXME wrong concat
		uriParsed, err = url.Parse(absUri)
		if err != nil {
			return err
		}
	}

	if uriParsed.Scheme == "" {
		uriParsed.Scheme = ctx.ScriptUri.Scheme
	}

	switch uriParsed.Scheme {
	case "file":
		assetGlob, err := filepath.Rel(path.Dir(ctx.InstallFromLocalPath), uriParsed.Path)
		if err != nil {
			return err
		}
		err = installAssetFromLocal(assetGlob, ctx)
		if err != nil {
			return err
		}
	case "http":
	case "https":
		assetRelPath, err := filepath.Rel(path.Dir(ctx.RawUri), absUri)
		if err != nil {
			return err
		}

		err = os.MkdirAll(path.Dir(ctx.InstallScriptDir+"/"+assetRelPath), 0777) // FIXME: perm
		if err != nil {
			return err
		}
		err = httpext.DownloadFile(absUri, ctx.InstallScriptDir+"/"+assetRelPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func installAssetFromLocal(assetGlob string, ctx *installCtx) error {
	installFromLocalDir := path.Dir(ctx.InstallFromLocalPath)
	installFromDirFSys := os.DirFS(installFromLocalDir)

	files, err := doublestar.Glob(installFromDirFSys, assetGlob, doublestar.WithFilesOnly())
	if err != nil {
		return err
	}

	for _, f := range files {
		fInstallFromPath := installFromLocalDir + "/" + f
		fInstallPath := ctx.InstallScriptDir + "/" + f
		_, err = fsext.CopyFile(fInstallFromPath, fInstallPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func Install(uri string) (*scriptmeta.ScriptMetadata, error) {
	ctx := newInstallCtx(uri)
	err := ctx.prepareInstallation()
	if err != nil {
		return nil, err
	}

	if !pathext.Exists(ctx.InstallFromLocalPath) {
		return nil, fmt.Errorf("script not found at '%s'", ctx.InstallFromLocalPath)
	}

	meta, err := ParseMetadata(ctx.InstallFromLocalPath)
	if err != nil {
		return nil, err
	}
	ctx.Meta = meta

	err = installMainFile(ctx)
	if err != nil {
		return nil, err
	}

	for _, assetUri := range ctx.Meta.Assets {
		err = installAsset(assetUri, ctx)
		if err != nil {
			return nil, err
		}
	}

	return ctx.Meta, nil
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
