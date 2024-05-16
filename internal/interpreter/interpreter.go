package interpreter

import "fmt"

type InterpreterInfo interface {
	GetNormalizedScriptPath(scriptPath string) string
	GetargsFilterOutEnvVar(varName string) bool
	GetargsIntro(tokens []string) string
	GetargsOutro(tokens []string) string
	GetEnvVarInjectCode(key string, value string) string
	GetCommentStarter() string
}

var (
	Interpreters map[string]InterpreterInfo = map[string]InterpreterInfo{
		"bash": &bashInter,
	}
)

func GetInterpreterInfo(interpreter string) (InterpreterInfo, error) {
	info, found := Interpreters[interpreter]
	if !found {
		return nil, fmt.Errorf("unsupported interpreter %s", interpreter)
	}
	return info, nil
}

func IsSupported(interpreter string) bool {
	_, found := Interpreters[interpreter]
	return found
}
