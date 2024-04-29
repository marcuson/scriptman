package rewriter

import (
	"fmt"
	"io"
	"marcuson/scriptman/internal/interpreter"
	scan "marcuson/scriptman/internal/script/internal/scan"
	"strings"

	"github.com/joho/godotenv"
)

type DotenvInjectorRewriter struct {
	envFileReader io.Reader
	interpreter   string
}

func NewDotenvInjectorRewriter(envFileReader io.Reader, interpreter string) *DotenvInjectorRewriter {
	return &DotenvInjectorRewriter{envFileReader: envFileReader, interpreter: interpreter}
}

func (obj *DotenvInjectorRewriter) RewriteBeforeAll() (string, error) {
	inter, interFound := interpreter.GetInterpreterInfo(obj.interpreter)
	if !interFound {
		return "", fmt.Errorf("interpreter '%s' not supported", obj.interpreter)
	}

	envMap, err := godotenv.Parse(obj.envFileReader)
	if err != nil {
		return "", err
	}

	injectLines := []string{"# SCRIPTMAN - LOADED ENV - START"}
	for k, v := range envMap {
		injectEnvLine := inter.GetEnvVarInjectCode(k, v)
		if injectEnvLine != "" {
			injectLines = append(injectLines, injectEnvLine)
		}
	}
	injectLines = append(injectLines, "# SCRIPTMAN - LOADED ENV - END", "")

	return strings.Join(injectLines, "\n"), nil
}

func (obj *DotenvInjectorRewriter) RewriteLine(line *scan.LineScript) (string, error) {
	return "", nil
}

func (obj *DotenvInjectorRewriter) RewriteAfterAll() (string, error) {
	return "", nil
}
