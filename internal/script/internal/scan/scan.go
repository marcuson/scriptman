package scan

import (
	"bufio"
	"fmt"
	"io"
	"marcuson/scriptman/internal/interpreter"
	"regexp"
	"slices"
	"strings"
)

var (
	RE_SPLIT_LINE                     *regexp.Regexp = regexp.MustCompile(`\s`)
	VALID_FIRST_LINE_COMMENT_STARTERS                = []string{"#", "//"}
	VALID_METADATA_STARTERS                          = []string{"@scriptman", "@sman"}
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
	return RE_SPLIT_LINE.Split(obj.Text, -1)
}

type Scanner struct {
	scanner        bufio.Scanner
	interpreter    string
	commentStarter string
	scriptErr      error
	_line          LineScript
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		scanner: *bufio.NewScanner(r),
		_line:   LineScript{LineIndex: -1},
	}
}

func (obj *Scanner) Err() error {
	if obj.scriptErr != nil {
		return obj.scriptErr
	}
	return obj.scanner.Err()
}

func (obj *Scanner) SetInterpreter(inter string) {
	obj.interpreter = inter
	interInfo, err := interpreter.GetInterpreterInfo(inter)
	if err == nil {
		obj.commentStarter = interInfo.GetCommentStarter()
	}
}

func (obj *Scanner) Interpreter() string {
	return obj.interpreter
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
	obj._line.IsShebang = obj._line.LineIndex == 0 && strings.HasPrefix(lineSplit[0], "#!")

	switch {
	case obj._line.LineIndex == 0:
		obj._line.IsComment = slices.ContainsFunc(VALID_FIRST_LINE_COMMENT_STARTERS,
			func(s string) bool { return strings.HasPrefix(lineSplit[0], s) })
	case obj._line.LineIndex > 0 && obj.commentStarter == "":
		obj.scriptErr = fmt.Errorf("scan second line or more without interpreter set")
		return false
	default:
		obj._line.IsComment = strings.HasPrefix(lineSplit[0], obj.commentStarter)
	}

	obj._line.IsMetadata = len(lineSplit) >= 2 && obj._line.IsComment &&
		slices.Index(VALID_METADATA_STARTERS, lineSplit[1]) >= 0

	return isScanOk
}

func (obj *Scanner) Line() *LineScript {
	return &obj._line
}
