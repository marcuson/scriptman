package rewriter

import (
	"marcuson/scriptman/internal/script/internal/scan"
	"slices"
)

type SecRewriter struct {
	shouldWrite     bool
	sections        []string
	isInsideSection bool
}

func NewSecRewriter(sections ...string) *SecRewriter {
	return &SecRewriter{
		shouldWrite:     true,
		sections:        sections,
		isInsideSection: false,
	}
}

func (obj *SecRewriter) RewriteBeforeAll() (string, error) {
	return "", nil
}

func (obj *SecRewriter) RewriteLine(line *scan.LineScript) (string, error) {
	var err error
	if line.IsMetadata {
		err = obj.checkMetadataLine(line)
	} else {
		err = obj.checkNonMetadataLine()
	}
	if err != nil {
		return "", err
	}

	if !obj.shouldWrite {
		return "", nil
	}

	return line.Text + "\n", nil
}

func (obj *SecRewriter) checkMetadataLine(line *scan.LineScript) error {
	lineSplit := line.LineSplit()

	metaKey := lineSplit[2]
	metaValue := lineSplit[3]

	switch metaKey {
	case "sec:start":
		obj.shouldWrite = slices.Contains(obj.sections, metaValue)
		obj.isInsideSection = true
	case "sec:end":
		obj.isInsideSection = false
	}

	return nil
}

func (obj *SecRewriter) checkNonMetadataLine() error {
	if obj.isInsideSection {
		return nil
	}

	obj.shouldWrite = true
	return nil
}

func (obj *SecRewriter) RewriteAfterAll() (string, error) {
	return "", nil
}
