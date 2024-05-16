package fsext

import (
	"os"
	"testing"

	"github.com/fluentassert/verify"
)

const (
	testsrcdir = "./_testdata"
	targetdir  = "./../../../_test/fsext"
)

func setup() error {
	return os.MkdirAll(targetdir, 0777)
}

func teardown() error {
	return os.RemoveAll(targetdir)
}

func TestCopyOk(t *testing.T) {
	err := setup()
	verify.NoError(err).Assert(t)
	defer teardown()

	srcPath := testsrcdir + "/test.txt"
	targetPath := targetdir + "/test.txt"
	b, err := CopyFile(srcPath, targetPath)
	verify.NoError(err).Assert(t)
	verify.Number(b).Equal(4).Assert(t)
}

func TestCopyNonExistentDirOk(t *testing.T) {
	err := setup()
	verify.NoError(err).Assert(t)
	defer teardown()

	srcPath := testsrcdir + "/test.txt"
	targetPath := targetdir + "/nested/test.txt"
	b, err := CopyFile(srcPath, targetPath)
	verify.NoError(err).Assert(t)
	verify.Number(b).Equal(4).Assert(t)
}
