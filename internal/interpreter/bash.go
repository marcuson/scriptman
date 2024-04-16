package interpreter

import "text/template"

type bashInterpreter struct {
	getargsIntroTpl *template.Template
	getargsOutroTpl *template.Template
}

func newBashInterpreter() bashInterpreter {
	introTpl := "echo {{.Tokens 0}}\n" +
		"set 2>/dev/null | while read a; do [[ $a == *=* ]] || break; echo $a; done\n" +
		"echo {{.Tokens 1}}\n"
	getargsIntroTpl, err := template.New("getargsIntro").Parse(introTpl)
	if err != nil {
		panic(err)
	}

	outroTpl := "echo {{.Tokens 2}}\n" +
		"set 2>/dev/null | while read a; do [[ $a == *=* ]] || break; echo $a; done\n" +
		"echo {{.Tokens 3}}\n"
	getargsOutroTpl, err := template.New("getargsIntro").Parse(outroTpl)
	if err != nil {
		panic(err)
	}

	return bashInterpreter{
		getargsIntroTpl: getargsIntroTpl,
		getargsOutroTpl: getargsOutroTpl,
	}
}

var bashInter bashInterpreter = newBashInterpreter()

func (obj *bashInterpreter) GetargsFilterOutEnvVar(varName string) bool {
	return varName == "_"
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
