package scriptmeta

import (
	"marcuson/scriptman/internal/config"
	"testing"

	"github.com/fluentassert/verify"
	"github.com/prashantv/gostub"
)

func TestScriptIdOk(t *testing.T) {
	m := NewScriptMetadata()
	m.Namespace = "test"
	m.Name = "script"

	id := m.ScriptId()
	verify.String(id).Equal("test-script").Assert(t)
}

func TestScriptIdDirOk(t *testing.T) {
	m := NewScriptMetadata()
	m.Namespace = "test"
	m.Name = "script"

	id := m.InstallScriptIdDir()
	verify.String(id).Equal("test/script").Assert(t)
}

func TestGetScriptPathFromIdOk(t *testing.T) {
	sDir := "/home/user/.local/share/" + config.SCRIPT_HOME + "/ns/name"
	xdgStub := gostub.StubFunc(&xdgSearchDataFile, sDir, nil)
	defer xdgStub.Reset()

	fgStub := gostub.StubFunc(&filepathGlob, []string{sDir + "/name.sh"}, nil)
	defer fgStub.Reset()

	f, p := GetScriptPathFromId("ns-name")
	verify.True(f).Assert(t)
	verify.String(p).Equal(sDir + "/name.sh").Assert(t)
}

func TestGetScriptPathFromIdMultiOk(t *testing.T) {
	sDir := "/home/user/.local/share/" + config.SCRIPT_HOME + "/ns/name"
	xdgStub := gostub.StubFunc(&xdgSearchDataFile, sDir, nil)
	defer xdgStub.Reset()

	fgStub := gostub.StubFunc(&filepathGlob, []string{sDir + "/name.sh", sDir + "/name.js"}, nil)
	defer fgStub.Reset()

	f, p := GetScriptPathFromId("ns-name")
	verify.True(f).Assert(t)
	verify.String(p).Equal(sDir + "/name.sh").Assert(t)
}

func TestGetScriptPathFromIdNotFound(t *testing.T) {
	sDir := "/home/user/.local/share/" + config.SCRIPT_HOME + "/ns/name"
	xdgStub := gostub.StubFunc(&xdgSearchDataFile, sDir, nil)
	defer xdgStub.Reset()

	fgStub := gostub.StubFunc(&filepathGlob, []string{}, nil)
	defer fgStub.Reset()

	f, _ := GetScriptPathFromId("ns-name")
	verify.False(f).Assert(t)
}
