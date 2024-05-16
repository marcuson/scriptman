package run

import (
	"testing"

	"github.com/fluentassert/verify"
)

const (
	testdir = "./../../_testdata"
)

func TestRunWithoutHooks(t *testing.T) {
	ctx, err := RunWithHooks(testdir+"/meta_ok_001.sh", RunHooks{})
	verify.NoError(err).Assert(t)
	verify.String(ctx.ScriptPath).Equal(testdir + "/meta_ok_001.sh").Require(t)
	verify.String(ctx.TmpScriptPath).Equal("../../_testdata/__run-meta_ok_001.sh").Require(t)
	verify.String(ctx.NormTmpScriptPath).Equal("../../_testdata/__run-meta_ok_001.sh").Require(t)
	verify.Slice(ctx.Rewriters).Len(0).Require(t)
	verify.Map(ctx.Props).Len(0).Require(t)
}
