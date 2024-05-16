package processor

import (
	"marcuson/scriptman/internal/script/internal/scan"
	"testing"

	"github.com/fluentassert/verify"
)

func TestInterpreterShebangOk(t *testing.T) {
	processor := NewInterpreterProcessor()

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "#!/usr/bin/env bash",
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	inter := processor.Interpreter()
	verify.String(inter).Equal("bash").Assert(t)
}

func TestInterpreterMetadataOk(t *testing.T) {
	processor := NewInterpreterProcessor()

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman interpreter bash",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	inter := processor.Interpreter()
	verify.String(inter).Equal("bash").Assert(t)
}

func TestInterpreterShebangNoFirstLine(t *testing.T) {
	processor := NewInterpreterProcessor()

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "#!/usr/bin/env bash",
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	inter := processor.Interpreter()
	verify.String(inter).Equal("").Assert(t)
}

func TestInterpreterMetadataNoFirstLine(t *testing.T) {
	processor := NewInterpreterProcessor()

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman interpreter bash",
		LineIndex:  1,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	inter := processor.Interpreter()
	verify.String(inter).Equal("").Assert(t)
}
