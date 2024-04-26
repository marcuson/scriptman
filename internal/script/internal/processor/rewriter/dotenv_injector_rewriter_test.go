package rewriter

import (
	scan "marcuson/scriptman/internal/script/internal/scan"
	"strings"
	"testing"

	"github.com/fluentassert/verify"
)

func TestDotenvInjectorRewriterBashOk(t *testing.T) {
	envContent := "ENVPROP1=env1\n" +
		"#comment to ignore\n" +
		"ENVPROP2=env2\n"
	reader := strings.NewReader(envContent)

	rewriter := NewDotenvInjectorRewriter(reader, "bash")

	result, err := rewriter.RewriteBeforeAll()
	verify.NoError(err).Require(t)
	verify.String(result).Equal("export ENVPROP1=env1\nexport ENVPROP2=env2").Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  1,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 2,
	})
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)

	result, err = rewriter.RewriteLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  3,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)

	result, err = rewriter.RewriteAfterAll()
	verify.NoError(err).Require(t)
	verify.String(result).Equal("").Assert(t)
}
