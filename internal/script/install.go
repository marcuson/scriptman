package script

import (
	"fmt"
	"net/url"
	"strings"

	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/codeext"
	"marcuson/scriptman/internal/utils/pathext"

	"github.com/adrg/xdg"
)

const (
	FS_INSTALL_PROTOCOL   = "fs"
	HTTP_INSTALL_PROTOCOL = "http"
	GIT_INSTALL_PROTOCOL  = "git"
)

type scriptInstallCtx struct {
	RawUri    string
	ScriptUri *url.URL

	InstallFromLocalFile  string
	Meta                  *scriptmeta.ScriptMetadata
	InstallTargetMainFile string
	InstallTargetDir      string
}

func newScriptInstallCtx(uri string) (*scriptInstallCtx, error) {
	ctx := &scriptInstallCtx{
		RawUri: uri,
	}

	uriParsed, err := url.Parse(ctx.RawUri)
	if err != nil {
		return nil, err
	}

	if uriParsed.Scheme == "" {
		uriParsed.Scheme = "file"
	}

	ctx.ScriptUri = uriParsed

	return ctx, nil
}

type scriptInstaller interface {
	prepare(ctx *scriptInstallCtx) error
	installMainFile(ctx *scriptInstallCtx) error
}

type assetInstallCtx struct {
	RawUri            string
	AssetUri          *url.URL
	InstallTargetFile string

	ScriptInstallCtx *scriptInstallCtx
}

func newAssetInstallCtx(assetUri string, scriptCtx *scriptInstallCtx) (*assetInstallCtx, error) {
	ctx := &assetInstallCtx{
		RawUri:           assetUri,
		ScriptInstallCtx: scriptCtx,
	}

	uriParsed, err := url.Parse(assetUri)
	if err != nil {
		return nil, err
	}

	ctx.AssetUri = uriParsed

	if ctx.IsRelative() {
		ctx.InstallTargetFile = ctx.ScriptInstallCtx.InstallTargetDir + "/" + uriParsed.Path
	} else {
		ctx.InstallTargetFile = ctx.ScriptInstallCtx.InstallTargetDir + "/" + pathext.Name(uriParsed.Path)
	}

	return ctx, nil
}

func (obj *assetInstallCtx) IsRelative() bool {
	return obj.AssetUri.Scheme == ""
}

type assetInstaller interface {
	installAsset(ctx *assetInstallCtx) error
}

var (
	assetInstallers = make(map[string]assetInstaller)
)

func getInstallProtocol(uriScheme string) string {
	switch uriScheme {
	case "":
	case "file":
		return FS_INSTALL_PROTOCOL
	case "http":
		return HTTP_INSTALL_PROTOCOL
	case "https":
		if strings.HasSuffix(uriScheme, ".git") {
			return GIT_INSTALL_PROTOCOL
		}

		return HTTP_INSTALL_PROTOCOL
	case "git":
		return GIT_INSTALL_PROTOCOL
	}

	return ""
}

func getScriptInstaller(ctx *scriptInstallCtx) (scriptInstaller, error) {
	scriptInstallProtocol := getInstallProtocol(ctx.ScriptUri.Scheme)
	var installer scriptInstaller
	switch scriptInstallProtocol {
	case FS_INSTALL_PROTOCOL:
		installer = &fsScriptInstaller{}
	case HTTP_INSTALL_PROTOCOL:
		installer = &httpScriptInstaller{}
	case GIT_INSTALL_PROTOCOL:
		installer = &gitScriptInstaller{}
	default:
		return nil, fmt.Errorf("unsupported install URI scheme '%s'", ctx.ScriptUri.Scheme)
	}

	return installer, nil
}

func getAssetInstaller(ctx *assetInstallCtx) (assetInstaller, error) {
	assetScheme := codeext.Tern(
		ctx.IsRelative(), ctx.ScriptInstallCtx.ScriptUri.Scheme, ctx.AssetUri.Scheme)

	assetProtocol := getInstallProtocol(assetScheme)
	installer, installerCacheHit := assetInstallers[assetProtocol]

	if installerCacheHit {
		return installer, nil
	}

	switch assetProtocol {
	case FS_INSTALL_PROTOCOL:
		installer = &fsAssetInstaller{}
	case HTTP_INSTALL_PROTOCOL:
		installer = &httpAssetInstaller{}
	case GIT_INSTALL_PROTOCOL:
		installer = &gitAssetInstaller{}
	default:
		return nil, fmt.Errorf("unsupported install URI scheme '%s'", ctx.AssetUri.Scheme)
	}

	assetInstallers[assetProtocol] = installer
	return installer, nil
}

func Install(uri string) (*scriptmeta.ScriptMetadata, error) {
	ctx, err := newScriptInstallCtx(uri)
	if err != nil {
		return nil, err
	}

	scriptInstlr, err := getScriptInstaller(ctx)
	if err != nil {
		return nil, err
	}

	err = scriptInstlr.prepare(ctx)
	if err != nil {
		return nil, err
	}

	if !pathext.Exists(ctx.InstallFromLocalFile) {
		return nil, fmt.Errorf("script not found at '%s'", ctx.InstallFromLocalFile)
	}

	meta, err := ParseMetadata(ctx.InstallFromLocalFile)
	if err != nil {
		return nil, err
	}
	ctx.Meta = meta

	installDir, err := xdg.DataFile(config.SCRIPT_HOME + "/" +
		ctx.Meta.InstallScriptIdDir())
	if err != nil {
		return nil, err
	}
	ctx.InstallTargetDir = installDir
	ctx.InstallTargetMainFile = installDir + "/" + ctx.Meta.Name + ctx.Meta.Ext

	err = scriptInstlr.installMainFile(ctx)
	if err != nil {
		return nil, err
	}

	for _, assetUri := range ctx.Meta.Assets {
		assetCtx, err := newAssetInstallCtx(assetUri, ctx)
		if err != nil {
			return nil, err
		}

		assetInstlr, err := getAssetInstaller(assetCtx)
		if err != nil {
			return nil, err
		}

		err = assetInstlr.installAsset(assetCtx)
		if err != nil {
			return nil, err
		}
	}

	return ctx.Meta, nil
}
