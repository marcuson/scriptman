package script

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestParseOk(t *testing.T) {
	meta, err := ParseMetadata("./testdata/meta_ok_001.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marcuson/test").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_001").Assert(t)
}

func TestParseOkNoShebang(t *testing.T) {
	meta, err := ParseMetadata("./testdata/meta_ok_002.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marcuson/test").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_002").Assert(t)
}

func TestParseOkShebangAndInterpreter(t *testing.T) {
	meta, err := ParseMetadata("./testdata/meta_ok_003.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marcuson/test").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_003").Assert(t)
}
