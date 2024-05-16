package script

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/execext"
	"marcuson/scriptman/internal/utils/fsext"
	"marcuson/scriptman/internal/utils/hashext"
	"marcuson/scriptman/internal/utils/pathext"
	"os"
	"os/exec"
	"strings"

	"github.com/adrg/xdg"
)

func getGitRepoUrl(uri string) string {
	rawUriSplit := strings.Split(uri, ":")
	gitRepoUrl := strings.Join(rawUriSplit[:len(rawUriSplit)-1], ":")
	return gitRepoUrl
}

func getGitTmpDir(uri string) (string, error) {
	gitRepoUrl := getGitRepoUrl(uri)
	gitUrlHash := hashext.Md5Str(gitRepoUrl)

	tmpInstallPath, err := xdg.DataFile(config.TMP_GIT_INSTALL_DIR + "/" + gitUrlHash)
	if err != nil {
		return "", err
	}

	return tmpInstallPath, nil
}

func cloneRepo(repoUrl string, dest string) error {
	if pathext.Exists(dest) {
		err := os.RemoveAll(dest)
		if err != nil {
			return err
		}
	}

	gitRepoUrl := repoUrl
	gitBranch := ""

	if atIdx := strings.LastIndex(repoUrl, "@"); atIdx >= 0 {
		gitBranch = repoUrl[atIdx+1:]
		gitRepoUrl = repoUrl[:atIdx]
	}

	cloneCmdStr := "clone --depth 1"
	if gitBranch != "" {
		cloneCmdStr = cloneCmdStr + " --branch " + gitBranch
	}
	cloneCmdStr = cloneCmdStr + " " + gitRepoUrl + " " + dest
	cloneCmd := exec.Command("git", execext.StrToArgs(cloneCmdStr)...)
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stdin = os.Stdin

	return cloneCmd.Run()
}

type gitScriptInstaller struct{}

func (obj *gitScriptInstaller) prepare(ctx *scriptInstallCtx) error {
	gitRepoUrl := getGitRepoUrl(ctx.RawUri)
	tmpInstallDir, err := getGitTmpDir(ctx.RawUri)
	if err != nil {
		return err
	}

	if !ctx.TmpPathsCache.Has(tmpInstallDir) {
		err = cloneRepo(gitRepoUrl, tmpInstallDir)
		if err != nil {
			return err
		}

		ctx.TmpPathsCache.Register(tmpInstallDir)
	}

	rawUriSplit := strings.Split(ctx.RawUri, ":")
	mainFileGitPath := rawUriSplit[len(rawUriSplit)-1]
	ctx.InstallFromLocalFile = tmpInstallDir + "/" + mainFileGitPath

	return nil
}

func (obj *gitScriptInstaller) installMainFile(ctx *scriptInstallCtx) error {
	err := os.MkdirAll(ctx.InstallTargetDir, 0770)
	if err != nil {
		return err
	}

	_, err = fsext.CopyFile(ctx.InstallFromLocalFile, ctx.InstallTargetMainFile)
	return err
}

type gitAssetInstaller struct {
	fsAssetInstlr *fsAssetInstaller
}

func (obj *gitAssetInstaller) installAsset(ctx *assetInstallCtx) error {
	if ctx.IsRelative() {
		return obj.fsAssetInstlr.installAsset(ctx)
	}

	gitRepoUrl := getGitRepoUrl(ctx.RawUri)
	tmpInstallDir, err := getGitTmpDir(ctx.RawUri)
	if err != nil {
		return err
	}

	if !ctx.TmpPathsCache.Has(tmpInstallDir) &&
		!ctx.ScriptInstallCtx.TmpPathsCache.Has(tmpInstallDir) {

		err = cloneRepo(gitRepoUrl, tmpInstallDir)
		if err != nil {
			return err
		}

		ctx.TmpPathsCache.Register(tmpInstallDir)
	}

	fsUri := "file:" + strings.Replace(ctx.RawUri, gitRepoUrl+":", tmpInstallDir+"/", 1)
	fsCtx, err := newAssetInstallCtx(fsUri, ctx.ScriptInstallCtx)
	if err != nil {
		return err
	}
	return obj.fsAssetInstlr.installAsset(fsCtx)
}
