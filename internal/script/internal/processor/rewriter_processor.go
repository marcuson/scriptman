package processor

import (
	"io"
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/scan"
)

type RewriterProcessor struct {
	writer    io.StringWriter
	rewriters []rewriter.Rewriter

	isBeforeAllDone bool
}

func NewRewriterProcessor(w io.StringWriter, rewriters ...rewriter.Rewriter) *RewriterProcessor {
	return &RewriterProcessor{
		writer:    w,
		rewriters: rewriters,
	}
}

func (obj *RewriterProcessor) rewriteString(str string) error {
	if str == "" {
		return nil
	}

	_, err := obj.writer.WriteString(str)
	return err
}

func (obj *RewriterProcessor) processBeforeAll() error {
	for _, p := range obj.rewriters {
		newline, err := p.RewriteBeforeAll()

		if err != nil {
			return err
		}

		if err = obj.rewriteString(newline); err != nil {
			return err
		}
	}

	return nil
}

func (obj *RewriterProcessor) ProcessStart() error {
	return nil
}

func (obj *RewriterProcessor) ProcessLine(line *scan.LineScript) error {
	if !obj.isBeforeAllDone && !line.IsShebang {
		obj.processBeforeAll()
		obj.isBeforeAllDone = true
	}

	for _, p := range obj.rewriters {
		newline, err := p.RewriteLine(line)

		if err != nil {
			return err
		}

		if err = obj.rewriteString(newline); err != nil {
			return err
		}
	}

	return nil
}

func (obj *RewriterProcessor) ProcessEnd() error {
	for _, p := range obj.rewriters {
		newline, err := p.RewriteAfterAll()

		if err != nil {
			return err
		}

		if err = obj.rewriteString(newline); err != nil {
			return err
		}
	}

	return nil
}
