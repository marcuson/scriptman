package rewriter

import (
	"marcuson/scriptman/internal/script/internal/scan"
	"testing"

	"github.com/fluentassert/verify"
)

func TestSecRewriterBasicOk(t *testing.T) {
	rewriter := NewSecRewriter("run")

	result, err := rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:start run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "Code inside run"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:end run\n").
		Assert(t)
}

func TestRewriterProcessorCommonOk(t *testing.T) {
	rewriter := NewSecRewriter("run")

	result, err := rewriter.RewriteLine(&scan.LineScript{
		Text:      `#!/bin/env bash`,
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("#!/bin/env bash\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common l1"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common l1"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common l2"`,
		LineIndex: 5,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common l2"` + "\n").
		Assert(t)
}

func TestRewriterProcessorCommonAndSecOk(t *testing.T) {
	rewriter := NewSecRewriter("run")

	result, err := rewriter.RewriteLine(&scan.LineScript{
		Text:      `#!/bin/env bash`,
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("#!/bin/env bash\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common before"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common before"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:start run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 3,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "Code inside run"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  4,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:end run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common after"`,
		LineIndex: 5,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common after"` + "\n").
		Assert(t)
}

func TestRewriterProcessorCommonSkipOtherSecOk(t *testing.T) {
	rewriter := NewSecRewriter("run")

	result, err := rewriter.RewriteLine(&scan.LineScript{
		Text:      `#!/bin/env bash`,
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("#!/bin/env bash\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common before"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common before"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:start run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 3,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "Code inside run"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  4,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:end run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common between"`,
		LineIndex: 5,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common between"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start other",
		LineIndex:  6,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside other"`,
		LineIndex: 7,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end other",
		LineIndex:  8,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common after"`,
		LineIndex: 9,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common after"` + "\n").
		Assert(t)
}

func TestRewriterProcessorCommonMultiSecOk(t *testing.T) {
	rewriter := NewSecRewriter("run", "other")

	result, err := rewriter.RewriteLine(&scan.LineScript{
		Text:      `#!/bin/env bash`,
		LineIndex: 0,
		IsShebang: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("#!/bin/env bash\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common before"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common before"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:start run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 3,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "Code inside run"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  4,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:end run\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common between"`,
		LineIndex: 5,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common between"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start other",
		LineIndex:  6,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:start other\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside other"`,
		LineIndex: 7,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "Code inside other"` + "\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end other",
		LineIndex:  8,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal("# @scriptman sec:end other\n").
		Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "common after"`,
		LineIndex: 9,
	})
	verify.NoError(err).Require(t)
	verify.String(result).
		Equal(`echo "common after"` + "\n").
		Assert(t)
}
