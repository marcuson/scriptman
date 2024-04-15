package scriptmeta

import (
	"testing"

	"github.com/fluentassert/verify"
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
