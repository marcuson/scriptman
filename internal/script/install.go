package script

import (
	"fmt"
	"net/url"
	"strings"

	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/pathext"

	"github.com/adrg/xdg"
)

const (
	FS_INSTALL_PROTOCOL   = "fs"
	HTTP_INSTALL_PROTOCOL = "http"
	GIT_INSTALL_PROTOCOL  = "git"
)

func getInstallProtocol(uriScheme string, rawUri string, fallback ...string) string {
	switch uriScheme {
	case "":
	case "file":
		return FS_INSTALL_PROTOCOL
	case "http":
		return HTTP_INSTALL_PROTOCOL
	case "https":
		rawSplit := strings.Split(rawUri, ":")
		var baseUrl string
		if atIdx := strings.LastIndex(rawSplit[1], "@"); atIdx >= 0 {
			baseUrl = rawSplit[1][:atIdx]
		} else {
			baseUrl = rawSplit[1]
		}

		if strings.HasSuffix(baseUrl, ".git") {
			return GIT_INSTALL_PROTOCOL
		}

		return HTTP_INSTALL_PROTOCOL
	case "git":
		return GIT_INSTALL_PROTOCOL
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	return ""
}

type tmpPathCacher struct {
	tmpPaths map[string]bool
}

func newTmpPathCacher() *tmpPathCacher {
	return &tmpPathCacher{
		tmpPaths: make(map[string]bool),
	}
}

func (obj *tmpPathCacher) Has(path string) bool {
	_, found := obj.tmpPaths[path]
	return found
}

func (obj *tmpPathCacher) Register(path string) {
	obj.tmpPaths[path] = true
}

type scriptInstallCtx struct {
	RawUri          string
	ScriptUri       *url.URL
	InstallProtocol string

	InstallFromLocalFile  string
	Meta                  *scriptmeta.ScriptMetadata
	InstallTargetMainFile string
	InstallTargetDir      string

	TmpPathsCache *tmpPathCacher
}

func newScriptInstallCtx(uri string) (*scriptInstallCtx, error) {
	ctx := &scriptInstallCtx{
		RawUri:        uri,
		TmpPathsCache: newTmpPathCacher(),
	}

	uriParsed, err := url.Parse(ctx.RawUri)
	if err != nil {
		return nil, err
	}

	if uriParsed.Scheme == "" {
		uriParsed.Scheme = "file"
	}

	ctx.ScriptUri = uriParsed
	ctx.InstallProtocol = getInstallProtocol(uriParsed.Scheme, uri)

	return ctx, nil
}

type scriptInstaller interface {
	prepare(ctx *scriptInstallCtx) error
	installMainFile(ctx *scriptInstallCtx) error
}

type assetInstallCtx struct {
	RawUri            string
	AssetUri          *url.URL
	InstallProtocol   string
	InstallTargetFile string

	TmpPathsCache    *tmpPathCacher
	ScriptInstallCtx *scriptInstallCtx
}

func newAssetInstallCtx(assetUri string, scriptCtx *scriptInstallCtx) (*assetInstallCtx, error) {
	ctx := &assetInstallCtx{
		RawUri:           assetUri,
		ScriptInstallCtx: scriptCtx,
		TmpPathsCache:    newTmpPathCacher(),
	}

	uriParsed, err := url.Parse(assetUri)
	if err != nil {
		return nil, err
	}

	ctx.AssetUri = uriParsed
	ctx.InstallProtocol = getInstallProtocol(uriParsed.Scheme, assetUri, scriptCtx.InstallProtocol)

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

func getScriptInstaller(ctx *scriptInstallCtx) (scriptInstaller, error) {
	scriptInstallProtocol := ctx.InstallProtocol

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
	assetProtocol := ctx.InstallProtocol
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
		installer = &gitAssetInstaller{
			fsAssetInstlr: &fsAssetInstaller{},
		}
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

	meta, err := ParseMetadataHeaderOnly(ctx.InstallFromLocalFile)
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
