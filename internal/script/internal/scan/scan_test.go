package scan

import (
	"os"
	"testing"

	"github.com/fluentassert/verify"
)

const (
	testdir = "../../_testdata"
)

func TestScanOk(t *testing.T) {
	file, err := os.Open(testdir + "/meta_ok_001.sh")
	if err != nil {
		t.Fatal("unable to open test file")
	}
	defer file.Close()

	scanner := NewScanner(file)

	lineOk := scanner.Scan()
	line := scanner.Line()
	verify.True(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
	verify.Number(line.LineIndex).Equal(0).Assert(t)
	verify.String(line.Text).Equal("#!/usr/bin/env bash").Assert(t)
	verify.True(line.IsShebang).Assert(t)

	lineOk = scanner.Scan()
	line = scanner.Line()
	verify.True(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
	verify.Number(line.LineIndex).Equal(1).Assert(t)
	verify.String(line.Text).Equal("").Assert(t)

	lineOk = scanner.Scan()
	line = scanner.Line()
	verify.True(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
	verify.Number(line.LineIndex).Equal(2).Assert(t)
	verify.String(line.Text).Equal("# @scriptman namespace marctest").Assert(t)
	verify.True(line.IsMetadata).Assert(t)

	lineOk = scanner.Scan()
	line = scanner.Line()
	verify.True(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
	verify.Number(line.LineIndex).Equal(3).Assert(t)
	verify.String(line.Text).Equal("# @scriptman name meta_ok_001").Assert(t)
	verify.True(line.IsMetadata).Assert(t)

	lineOk = scanner.Scan()
	line = scanner.Line()
	verify.True(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
	verify.Number(line.LineIndex).Equal(4).Assert(t)
	verify.String(line.Text).Equal("").Assert(t)

	lineOk = scanner.Scan()
	line = scanner.Line()
	verify.True(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
	verify.Number(line.LineIndex).Equal(5).Assert(t)
	verify.String(line.Text).Equal(`echo "test meta_ok_001"`).Assert(t)

	lineOk = scanner.Scan()
	verify.False(lineOk).Require(t)
	verify.NoError(scanner.Err()).Require(t)
}
