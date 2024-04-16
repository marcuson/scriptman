package rewriter

import (
	"marcuson/scriptman/internal/script/internal/scan"
	"testing"

	"github.com/fluentassert/verify"
)

func TestGetargsInjectorRewriterBasicOk(t *testing.T) {
	rewriter := NewGetargsInjectorRewriter()
	rewriter.SetIntro("TEST INTRO")
	rewriter.SetOutro("TEST OUTRO")

	result, err := rewriter.RewriteBeforeAll()
	verify.NoError(err).Require(t)
	verify.String(result).Equal("TEST INTRO").Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)

	result, err = rewriter.RewriteAfterAll()
	verify.NoError(err).Require(t)
	verify.String(result).Equal("TEST OUTRO").Assert(t)
}
