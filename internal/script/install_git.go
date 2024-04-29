package script

import (
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/utils/execext"
	"marcuson/scriptman/internal/utils/fsext"
	"marcuson/scriptman/internal/utils/hashext"
	"os"
	"os/exec"
	"strings"

	"github.com/adrg/xdg"
)

func getGitTmpDir(uri string) (string, error) {
	uriSplit := strings.Split(uri, ":")
	baseUrl := uriSplit[1]
	gitUrlHash := hashext.Md5Str(baseUrl)

	tmpInstallPath, err := xdg.DataFile(config.TMP_GIT_INSTALL_DIR + "/" + gitUrlHash)
	if err != nil {
		return "", err
	}

	return tmpInstallPath, nil
}

type gitScriptInstaller struct{}

func (obj *gitScriptInstaller) prepare(ctx *scriptInstallCtx) error {
	tmpInstallDir, err := getGitTmpDir(ctx.RawUri)
	if err != nil {
		return err
	}

	rawUriSplit := strings.Split(ctx.RawUri, ":")
	mainFileGitPath := rawUriSplit[len(rawUriSplit)-1]
	ctx.InstallFromLocalFile = tmpInstallDir + "/" + mainFileGitPath

	gitRepoUrl := strings.Join(rawUriSplit[:len(rawUriSplit)-1], ":")
	gitBranch := ""

	if atIdx := strings.LastIndex(gitRepoUrl, "@"); atIdx >= 0 {
		gitBranch = gitRepoUrl[atIdx+1:]
		gitRepoUrl = gitRepoUrl[:atIdx]
	}

	cloneCmdStr := "clone --depth 1"
	if gitBranch != "" {
		cloneCmdStr = cloneCmdStr + " --branch " + gitBranch
	}
	cloneCmdStr = cloneCmdStr + " " + gitRepoUrl + " " + tmpInstallDir
	cloneCmd := exec.Command("git", execext.StrToArgs(cloneCmdStr)...)
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stdin = os.Stdin

	err = cloneCmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (obj *gitScriptInstaller) installMainFile(ctx *scriptInstallCtx) error {
	err := os.MkdirAll(ctx.InstallTargetDir, 0777) // FIXME: perm
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
	// FIXME: Clone repo if absolute first
	return obj.fsAssetInstlr.installAsset(ctx)
}
