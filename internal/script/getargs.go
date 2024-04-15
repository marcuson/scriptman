package script

import (
	"bytes"
	"fmt"
	"io"
	"marcuson/scriptman/internal/config"
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/run"
	"marcuson/scriptman/internal/utils/codeext"
	"os"
	"slices"
	"strings"
)

var (
	getargsE0Tokens = []string{
		"__e0:s__",
		"__e0:e__",
	}
	getargsE1Tokens = []string{
		"__e1:s__",
		"__e1:e__",
	}
	getargsTokens = append(getargsE0Tokens, getargsE1Tokens...)
)

type getargsStdout struct {
	stdout        io.Writer
	e0            *bytes.Buffer
	e1            *bytes.Buffer
	currentWriter io.Writer

	tokenEnv         string
	tokenIndex       int
	tokenSearch      string
	tokenSearchBytes []byte
	isOutsideEnv     bool
}

func newGetargsStdout() *getargsStdout {
	ret := &getargsStdout{
		stdout:     os.Stdout,
		e0:         new(bytes.Buffer),
		e1:         new(bytes.Buffer),
		tokenIndex: -1,
	}

	ret.currentWriter = ret.stdout
	ret.setNextTokenSearch()
	return ret
}

func (obj *getargsStdout) setTokenSearch(token string) {
	obj.tokenSearch = token
	obj.tokenSearchBytes = []byte(obj.tokenSearch)
	splitToken := strings.Split(strings.Trim(token, "_"), ":")
	obj.tokenEnv = splitToken[0]
	obj.isOutsideEnv = splitToken[1] == "s"
}

func (obj *getargsStdout) setNextTokenSearch() {
	if obj.tokenIndex >= len(getargsTokens)-1 {
		return
	}

	obj.tokenIndex += 1
	obj.setTokenSearch(getargsTokens[obj.tokenIndex])
}

func (obj *getargsStdout) Write(p []byte) (n int, err error) {
	from := 0

	for {
		idx := bytes.Index(p[from:], obj.tokenSearchBytes)
		if idx < 0 {
			break
		}

		if obj.isOutsideEnv {
			_, err = obj.stdout.Write(p[from : from+idx])
			from = idx + len(obj.tokenSearchBytes)
			obj.currentWriter = codeext.Tern(obj.tokenEnv == "e0", obj.e0, obj.e1)
		} else {
			_, err = obj.currentWriter.Write(p[from : from+idx])
			from = idx + len(obj.tokenSearchBytes)
			obj.currentWriter = obj.stdout
		}

		if err != nil {
			return 0, err
		}

		obj.setNextTokenSearch()
	}

	for ; from <= len(p)-1; from += 1 {
		if p[from] != '\r' && p[from] != '\n' {
			break
		}
	}

	if from <= len(p)-1 {
		if _, err = obj.currentWriter.Write(p[from:]); err != nil {
			return 0, err
		}
	}

	return len(p), nil
}

func envToMap(envStr string) map[string]string {
	envLines := slices.DeleteFunc(
		strings.Split(envStr, "\n"), func(s string) bool { return s == "" })
	m := make(map[string]string, len(envLines))
	for _, l := range envLines {
		kv := strings.Split(l, "=")
		if kv[0] == "_" { // FIXME: Bash only
			continue
		}
		m[kv[0]] = kv[1]
	}
	return m
}

func envDiff(map1, map2 map[string]string) map[string]string {
	res := make(map[string]string)

	for k, v1 := range map1 {
		v2, found := map2[k]
		if !found {
			res[k] = v1
			continue
		}

		if v1 != v2 {
			res[k] = v1
		}
	}

	return res
}

func getargsPreRun(ctx *run.RunCtx) error {
	if _, found := ctx.Meta.Sections[config.GETARGS_SECTION]; !found {
		return fmt.Errorf("section '%s' is not present in the script", config.GETARGS_SECTION)
	}

	getargsStdout := newGetargsStdout()
	ctx.InterpreterCmd.Stdout = getargsStdout
	ctx.Props["getargs_env"] = getargsStdout

	return nil
}

func getargsPostRun(ctx *run.RunCtx) error {
	gaWriter := ctx.Props["getargs_env"].(*getargsStdout)
	e0Str := gaWriter.e0.String()
	e1Str := gaWriter.e1.String()

	e0Map := envToMap(e0Str)
	e1Map := envToMap(e1Str)

	mapDiff := envDiff(e1Map, e0Map)
	ctx.Props["getargs_diff"] = mapDiff

	return nil
}

func Getargs(idOrPath string, out string) error {
	getargsAugmenter := rewriter.NewGetargsInjectorRewriter()
	// FIXME: This is valid only for bash!!!
	getargsAugmenter.SetIntro("echo " + getargsE0Tokens[0] + "\n" +
		"set 2>/dev/null | while read a; do [[ $a == *=* ]] || break; echo $a; done\n" +
		"echo " + getargsE0Tokens[1] + "\n")
	getargsAugmenter.SetOutro("echo " + getargsE1Tokens[0] + "\n" +
		"set 2>/dev/null | while read a; do [[ $a == *=* ]] || break; echo $a; done\n" +
		"echo " + getargsE1Tokens[1] + "\n")
	secRewriter := rewriter.NewSecRewriter(config.GETARGS_SECTION)

	hooks := run.RunHooks{
		PreRun:  getargsPreRun,
		PostRun: getargsPostRun,
	}

	ctx, err := run.RunWithHooks(idOrPath, hooks, getargsAugmenter, secRewriter)
	if err != nil {
		return err
	}

	envDelta := ctx.Props["getargs_diff"].(map[string]string)
	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for k, v := range envDelta {
		_, err := outFile.WriteString(k + "=" + v + "\n")
		if err != nil {
			return err
		}
	}

	return outFile.Sync()
}
