package execext

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestStrToArgsSpaces(t *testing.T) {
	cmd := "ls -la ."
	res := StrToArgs(cmd)
	verify.Slice(res).Len(3).Require(t)
	verify.Slice(res).Equivalent([]string{"ls", "-la", "."}).Require(t)
}

func TestStrToArgsMultiSpaces(t *testing.T) {
	cmd := "ls    -la     ."
	res := StrToArgs(cmd)
	verify.Slice(res).Len(3).Require(t)
	verify.Slice(res).Equivalent([]string{"ls", "-la", "."}).Require(t)
}

func TestStrToArgsTabs(t *testing.T) {
	cmd := "ls	-la	."
	res := StrToArgs(cmd)
	verify.Slice(res).Len(3).Require(t)
	verify.Slice(res).Equivalent([]string{"ls", "-la", "."}).Require(t)
}

func TestStrToArgsSpacesAndTabs(t *testing.T) {
	cmd := "ls	-la ."
	res := StrToArgs(cmd)
	verify.Slice(res).Len(3).Require(t)
	verify.Slice(res).Equivalent([]string{"ls", "-la", "."}).Require(t)
}
