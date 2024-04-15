package processor

import (
	"bytes"
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/scan"
	"marcuson/scriptman/internal/utils/codeext"
	"strconv"
	"testing"

	"github.com/fluentassert/verify"
	mock "github.com/stretchr/testify/mock"
)

func TestRewriterPassthroughOk(t *testing.T) {
	outbuf := new(bytes.Buffer)
	mockRew := rewriter.NewMockRewriter(t)
	mockRew.EXPECT().
		RewriteLine(mock.AnythingOfType("*scan.LineScript")).
		RunAndReturn(func(line *scan.LineScript) (string, error) {
			return line.Text + "\n", nil
		})
	processor := NewRewriterProcessor(outbuf, mockRew)

	err := processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	result := outbuf.String()
	verify.String(result).
		Equal("# @scriptman sec:start run\n" +
			`echo "Code inside run"` + "\n" +
			"# @scriptman sec:end run\n").
		Assert(t)
}

func TestRewriterTransformerOk(t *testing.T) {
	outbuf := new(bytes.Buffer)
	mockRew := rewriter.NewMockRewriter(t)
	mockRew.EXPECT().
		RewriteLine(mock.AnythingOfType("*scan.LineScript")).
		RunAndReturn(func(line *scan.LineScript) (string, error) {
			return strconv.Itoa(int(line.LineIndex)) + "\n", nil
		})
	processor := NewRewriterProcessor(outbuf, mockRew)

	err := processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	result := outbuf.String()
	verify.String(result).
		Equal("0\n1\n2\n").
		Assert(t)
}

func TestRewriterSkipOk(t *testing.T) {
	outbuf := new(bytes.Buffer)
	mockRew := rewriter.NewMockRewriter(t)
	mockRew.EXPECT().
		RewriteLine(mock.AnythingOfType("*scan.LineScript")).
		RunAndReturn(func(line *scan.LineScript) (string, error) {
			return "", nil
		})
	processor := NewRewriterProcessor(outbuf, mockRew)

	err := processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	result := outbuf.String()
	verify.String(result).
		Equal("").
		Assert(t)
}

func TestRewriterPassthroughAndSkipperOk(t *testing.T) {
	outbuf := new(bytes.Buffer)
	i := 0
	mockRew := rewriter.NewMockRewriter(t)
	mockRew.EXPECT().
		RewriteLine(mock.AnythingOfType("*scan.LineScript")).
		RunAndReturn(func(line *scan.LineScript) (string, error) {
			newline := codeext.Tern(i%2 == 0, line.Text+"\n", "")
			i += 1
			return newline, nil
		})
	processor := NewRewriterProcessor(outbuf, mockRew)

	err := processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:start run",
		LineIndex:  0,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:      `echo "Code inside run"`,
		LineIndex: 1,
	})
	verify.NoError(err).Require(t)

	err = processor.ProcessLine(&scan.LineScript{
		Text:       "# @scriptman sec:end run",
		LineIndex:  2,
		IsMetadata: true,
	})
	verify.NoError(err).Require(t)

	result := outbuf.String()
	verify.String(result).
		Equal("# @scriptman sec:start run\n" +
			"# @scriptman sec:end run\n").
		Assert(t)
}
