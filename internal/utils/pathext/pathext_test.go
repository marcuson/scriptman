package pathext

import (
	"testing"

	"github.com/fluentassert/verify"
)

const (
	testsrcdir = "./_testdata"
)

func TestExists(t *testing.T) {
	path := testsrcdir + "/test.txt"
	exists := Exists(path)
	verify.True(exists).Assert(t)
}

func TestNotExists(t *testing.T) {
	path := testsrcdir + "/none.txt"
	exists := Exists(path)
	verify.False(exists).Assert(t)
}

func TestName(t *testing.T) {
	path := testsrcdir + "/test.txt"
	n := Name(path)
	verify.String(n).Equal("test").Assert(t)
}
