package processor

import (
	"marcuson/scriptman/internal/script/internal/scan"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"testing"

	"github.com/fluentassert/verify"
)

func TestMetadataShebangOk(t *testing.T) {
	processor := NewMetadataProcessor("")

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "#!/usr/bin/env bash",
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
}

func TestMetadataNamespaceOk(t *testing.T) {
	processor := NewMetadataProcessor("")

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman namespace marctest",
		LineIndex:  3,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.NoError(err).Require(t)
	verify.String(meta.Namespace).Equal("marctest").Assert(t)
}

func TestMetadataNameOk(t *testing.T) {
	processor := NewMetadataProcessor("")

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman name meta_ok_001",
		LineIndex:  4,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.NoError(err).Require(t)
	verify.String(meta.Name).Equal("meta_ok_001").Assert(t)
}

func TestMetadataSectionOk(t *testing.T) {
	processor := NewMetadataProcessor("")

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  6,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      `echo "test run_ok_001"`,
		LineIndex: 7,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  8,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.Map(meta.Sections).Len(1).Require(t)
	verify.Map(meta.Sections).
		ContainPair("run", &scriptmeta.ScriptSection{LineStart: 6, LineEnd: 8})
}

func TestMetadataAssetOk(t *testing.T) {
	processor := NewMetadataProcessor("")

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman asset info.txt",
		LineIndex:  1,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       `# @scriptman asset assets/**`,
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.Slice(meta.Assets).Len(2).Require(t)
	verify.Slice(meta.Assets).Contain("info.txt")
	verify.Slice(meta.Assets).Contain("assets/**")
}

func TestMetadataGetargsTplOk(t *testing.T) {
	processor := NewMetadataProcessor("")

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman getargs-tpl info.txt",
		LineIndex:  1,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.String(meta.GetargsTpl).Equal("info.txt").Require(t)
}

func TestMetadataFillMissingOk(t *testing.T) {
	scriptPath := testdir + "/meta_ok_001.sh"
	processor := NewMetadataProcessor(scriptPath)

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.String(meta.Namespace).Equal("_nons_").Assert(t)
	verify.String(meta.Name).Equal("meta_ok_001").Assert(t)
}

func TestMetadataHeadOnlyWithInstructionOk(t *testing.T) {
	scriptPath := testdir + "/meta_ok_001.sh"
	processor := NewMetadataProcessor(scriptPath)
	processor.SetHeadOnly(true)

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "#!/usr/bin/env bash",
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "",
		LineIndex: 1,
		IsEmpty:   true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "# comment #1",
		LineIndex: 2,
		IsComment: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman namespace example",
		LineIndex:  3,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "# comment #2",
		LineIndex: 4,
		IsComment: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "",
		LineIndex: 5,
		IsEmpty:   true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "first instruction",
		LineIndex: 6,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Assert(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "second instruction",
		LineIndex: 7,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Assert(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("example").Assert(t)
}

func TestMetadataHeadOnlyWithSectionOk(t *testing.T) {
	scriptPath := testdir + "/meta_ok_001.sh"
	processor := NewMetadataProcessor(scriptPath)
	processor.SetHeadOnly(true)

	err := processor.ProcessStart()
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "#!/usr/bin/env bash",
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "",
		LineIndex: 1,
		IsEmpty:   true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "# comment #1",
		LineIndex: 2,
		IsComment: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman namespace example",
		LineIndex:  3,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "# comment #2",
		LineIndex: 4,
		IsComment: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "",
		LineIndex: 5,
		IsEmpty:   true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "# @scriptman sec:start getargs",
		LineIndex: 6,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Assert(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      "# @scriptman sec:end getargs",
		LineIndex: 7,
	})
	verify.NoError(err).Require(t)
	verify.True(processor.IsProcessCompletedEarly()).Assert(t)

	err = processor.ProcessEnd()
	verify.NoError(err).Require(t)

	meta := processor.Metadata()
	verify.Obj(meta).NotEqual(nil).Assert(t)
	verify.String(meta.Interpreter).Equal("bash").Assert(t)
	verify.String(meta.Namespace).Equal("example").Assert(t)
}
