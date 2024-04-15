package rewriter

import (
	scan "marcuson/scriptman/internal/script/internal/scan"
)

type GetargsInjectorRewriter struct {
	intro string
	outro string
}

func NewGetargsInjectorRewriter() *GetargsInjectorRewriter {
	return &GetargsInjectorRewriter{}
}

func (obj *GetargsInjectorRewriter) SetIntro(intro string) {
	obj.intro = intro
}

func (obj *GetargsInjectorRewriter) SetOutro(outro string) {
	obj.outro = outro
}

func (obj *GetargsInjectorRewriter) RewriteBeforeAll() (string, error) {
	return obj.intro, nil
}

func (obj *GetargsInjectorRewriter) RewriteLine(line *scan.LineScript) (string, error) {
	return "", nil
}

func (obj *GetargsInjectorRewriter) RewriteAfterAll() (string, error) {
	return obj.outro, nil
}
