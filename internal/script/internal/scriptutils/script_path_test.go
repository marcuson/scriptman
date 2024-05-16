package scriptutils

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestFindScriptPathOk(t *testing.T) {
	found, path := FindScriptPath(testdir + "/meta_ok_001.sh")
	verify.True(found).Assert(t)
	verify.String(path).Equal(testdir + "/meta_ok_001.sh").Require(t)
}

func TestFindScriptPathNotFound(t *testing.T) {
	found, _ := FindScriptPath(testdir + "/none.sh")
	verify.False(found).Assert(t)
}
