package script

import (
	"fmt"
	"marcuson/scriptman/internal/utils/pathext"
	"os"
	"os/exec"
	"regexp"
)

func runCmdBash(interpreterCmd *exec.Cmd, scriptPath string) error {
	cmd := "-c 'source " + scriptPath + " && [[ $(type -t runScript) ]] " +
		"&& runScript || exit 100'"
	cmdSplit := regexp.MustCompile(`\s`).Split(cmd, -1)
	interpreterCmd.Args = append(interpreterCmd.Args, cmdSplit...)
	interpreterCmd.Stderr = os.Stderr
	interpreterCmd.Stdout = os.Stdout
	interpreterCmd.Stdin = os.Stdin

	err := interpreterCmd.Run()

	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); !ok || exiterr.ExitCode() != 100 {
			return err
		}
	}

	return nil
}

func Run(idOrPath string) error {
	instFound, scriptPath := GetScriptPathFromId(idOrPath)
	if !instFound {
		scriptPath = idOrPath
	}
	if !pathext.Exists(scriptPath) {
		return fmt.Errorf("cannot find script '%s' by id or path", idOrPath)
	}

	meta, err := ParseMetadata(scriptPath)
	if err != nil {
		return err
	}

	scriptPathNormalized := getNormalizedScriptPath(scriptPath, meta.Interpreter)
	interpreterCmd := exec.Command(meta.Interpreter)

	switch meta.Interpreter {
	case "bash":
		runCmdBash(interpreterCmd, scriptPathNormalized)
	default:
		return fmt.Errorf("unsupported interpreter %s", meta.Interpreter)
	}

	return nil
}
