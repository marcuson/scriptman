package scan

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var (
	reSplitLine *regexp.Regexp = regexp.MustCompile(`\s`)
)

type LineScript struct {
	Text      string
	LineIndex int32

	IsEmpty    bool
	IsShebang  bool
	IsComment  bool
	IsMetadata bool
}

func (obj *LineScript) LineSplit() []string {
	return reSplitLine.Split(obj.Text, -1)
}

type Scanner struct {
	scanner bufio.Scanner
	_line   LineScript
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		scanner: *bufio.NewScanner(r),
		_line:   LineScript{LineIndex: -1},
	}
}

func (obj *Scanner) Err() error {
	return obj.scanner.Err()
}

func (obj *Scanner) Scan() bool {
	isScanOk := obj.scanner.Scan()
	if err := obj.scanner.Err(); err != nil || !isScanOk {
		return false
	}

	obj._line.LineIndex += 1
	obj._line.Text = obj.scanner.Text()
	lineSplit := obj._line.LineSplit()

	obj._line.IsEmpty = len(lineSplit) <= 0 || lineSplit[0] == ""
	obj._line.IsShebang = strings.HasPrefix(lineSplit[0], "#!")
	obj._line.IsComment = strings.HasPrefix(lineSplit[0], "#")
	obj._line.IsMetadata = len(lineSplit) >= 2 && lineSplit[0] == "#" && lineSplit[1] == "@scriptman"

	return isScanOk
}

func (obj *Scanner) Line() *LineScript {
	return &obj._line
}
