package interpreter

import (
	"slices"
)

type bashInterpreter struct {
}

func newBashInterpreter() bashInterpreter {
	return bashInterpreter{}
}

var bashInter bashInterpreter = newBashInterpreter()

func (obj *bashInterpreter) GetCommentStarter() string {
	return "#"
}

func (obj *bashInterpreter) GetargsFilterOutEnvVar(varName string) bool {
	filterOutVars := []string{"_", "LINES", "COLUMNS"}
	return slices.Contains(filterOutVars, varName)
}

func (obj *bashInterpreter) GetargsIntro(tokens []string) string {
	return "echo " + tokens[0] + "\n" +
		"set 2>/dev/null | while read a; do [[ $a == *=* ]] || break; echo $a; done\n" +
		"echo " + tokens[1] + "\n"
}

func (obj *bashInterpreter) GetargsOutro(tokens []string) string {
	return "echo " + tokens[2] + "\n" +
		"set 2>/dev/null | while read a; do [[ $a == *=* ]] || break; echo $a; done\n" +
		"echo " + tokens[3] + "\n"
}

func (obj *bashInterpreter) GetEnvVarInjectCode(key string, value string) string {
	return "export " + key + `="` + value + `"`
}
