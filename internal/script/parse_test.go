package script

import (
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"testing"

	"github.com/fluentassert/verify"
)

func TestParseOk(t *testing.T) {
	meta, err := ParseMetadata(testdir + "/meta_ok_001.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marctest").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_001").Assert(t)
	verify.Map(meta.Sections).Empty()
}

func TestParseOkNoShebang(t *testing.T) {
	meta, err := ParseMetadata(testdir + "/meta_ok_002.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marctest").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_002").Assert(t)
	verify.Map(meta.Sections).Empty()
}

func TestParseOkShebangAndInterpreter(t *testing.T) {
	meta, err := ParseMetadata(testdir + "/meta_ok_003.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marctest").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_003").Assert(t)
	verify.Map(meta.Sections).Empty()
}

func TestParseOkWithRun(t *testing.T) {
	meta, err := ParseMetadata(testdir + "/run_ok_001.sh")

	verify.NoError(err).Require(t)
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("marctest").Assert(t)
	verify.String(meta.Name).Equal("run_ok_001").Assert(t)

	verify.Map(meta.Sections).Len(1).Require(t)
	verify.Map(meta.Sections).ContainPair("run", &scriptmeta.ScriptSection{LineStart: 5, LineEnd: 7})
}

func TestParseInterpreterOnlyOk(t *testing.T) {
	inter, err := ParseInterpreter(testdir + "/meta_ok_001.sh")

	verify.NoError(err).Require(t)
	verify.String(inter).Equal("bash").Assert(t)
}
