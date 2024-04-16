package interpreter

type InterpreterInfo interface {
	GetNormalizedScriptPath(scriptPath string) string
	GetargsFilterOutEnvVar(varName string) bool
	GetargsIntro(tokens []string) string
	GetargsOutro(tokens []string) string
}

var (
	Interpreters map[string]InterpreterInfo = map[string]InterpreterInfo{
		"bash": &bashInter,
	}
)

func GetInterpreterInfo(interpreter string) (InterpreterInfo, bool) {
	info, found := Interpreters[interpreter]
	return info, found
}

func IsSupported(interpreter string) bool {
	_, found := Interpreters[interpreter]
	return found
}
