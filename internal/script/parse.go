package script

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type scriptParser struct {
	scanner   bufio.Scanner
	line      string
	lineIndex int64
	lineSplit []string

	isEmpty    bool
	isShebang  bool
	isComment  bool
	isMetadata bool

	reSplitLine *regexp.Regexp
}

func newScriptParser(r io.Reader) *scriptParser {
	return &scriptParser{
		scanner:     *bufio.NewScanner(r),
		lineIndex:   -1,
		reSplitLine: regexp.MustCompile(`\s`),
	}
}

func (obj *scriptParser) scanErr() error {
	return obj.scanner.Err()
}

func (obj *scriptParser) scan() bool {
	isScanOk := obj.scanner.Scan()
	if err := obj.scanner.Err(); err != nil || !isScanOk {
		return false
	}

	obj.lineIndex += 1
	obj.line = obj.scanner.Text()
	obj.lineSplit = obj.reSplitLine.Split(obj.line, -1)

	obj.isEmpty = len(obj.lineSplit) <= 0 || obj.lineSplit[0] == ""
	obj.isShebang = strings.HasPrefix(obj.lineSplit[0], "#!")
	obj.isComment = strings.HasPrefix(obj.lineSplit[0], "#")
	obj.isMetadata = len(obj.lineSplit) >= 2 && obj.lineSplit[0] == "#" && obj.lineSplit[1] == "@scriptman"

	return obj.isEmpty || obj.isComment
}

func (obj *scriptParser) parseShebang(meta *ScriptMetadata) error {
	interpreter := obj.lineSplit[len(obj.lineSplit)-1]
	interpreter = strings.Replace(interpreter, "#!", "", 1)
	meta.Interpreter = interpreter
	return nil
}

func (obj *scriptParser) parseMetadata(meta *ScriptMetadata) error {
	metaKey := obj.lineSplit[2]
	metaValue := obj.lineSplit[3]

	switch metaKey {
	case "namespace":
		meta.Namespace = metaValue
	case "name":
		meta.Name = metaValue
	case "interpreter":
		if meta.Interpreter == "" {
			meta.Interpreter = metaValue
		}
	default:
		return fmt.Errorf("unknown meta key: %s", metaKey)
	}

	return nil
}

func (obj *scriptParser) parse(meta *ScriptMetadata) error {
	switch {
	case obj.isShebang:
		return obj.parseShebang(meta)
	case obj.isMetadata:
		return obj.parseMetadata(meta)
	default:
		return nil
	}
}

func ParseMetadata(path string) (*ScriptMetadata, error) {
	var metadata ScriptMetadata

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	parser := newScriptParser(file)
	for parser.scan() {
		if err := parser.scanErr(); err != nil {
			return nil, err
		}

		if err = parser.parse(&metadata); err != nil {
			return nil, err
		}
	}

	metadata.FillMissingMetadata(path)
	return &metadata, nil
}
