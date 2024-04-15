package run

import (
	"fmt"
	"marcuson/scriptman/internal/script/internal/handle"
	"marcuson/scriptman/internal/script/internal/processor"
	"marcuson/scriptman/internal/script/internal/processor/rewriter"
	"marcuson/scriptman/internal/script/internal/scriptmeta"
	"marcuson/scriptman/internal/utils/execext"
	"marcuson/scriptman/internal/utils/pathext"
	"os"
	"os/exec"
	"path"
)

type RunHookFn = func(ctx *RunCtx) error

type RunHooks struct {
	PreRun  RunHookFn
	PostRun RunHookFn
}

type RunCtx struct {
	ScriptPath string

	Meta              *scriptmeta.ScriptMetadata
	TmpScriptPath     string
	NormTmpScriptPath string
	InterpreterCmd    *exec.Cmd

	Props map[string]interface{}
}

func prepCmdBash(ctx *RunCtx) {
	ctx.InterpreterCmd.Args =
		append(ctx.InterpreterCmd.Args, execext.StrToArgs(ctx.NormTmpScriptPath)...)
	ctx.InterpreterCmd.Stderr = os.Stderr
	ctx.InterpreterCmd.Stdout = os.Stdout
	ctx.InterpreterCmd.Stdin = os.Stdin
}

func getTempScriptPath(scriptPath string) string {
	scriptBase := path.Base(scriptPath)
	scriptDir := path.Dir(scriptPath)
	return path.Join(scriptDir, "__run-"+scriptBase)
}

func createCtx(scriptPath string) *RunCtx {
	ctx := &RunCtx{
		ScriptPath:    scriptPath,
		Meta:          nil,
		TmpScriptPath: "",
		Props:         make(map[string]interface{}),
	}
	ctx.TmpScriptPath = getTempScriptPath(scriptPath)

	return ctx
}

func setupRun(scriptPath string, rewriters ...rewriter.Rewriter) (*RunCtx, error) {
	ctx := createCtx(scriptPath)

	scriptFile, err := os.Open(scriptPath)
	if err != nil {
		return nil, err
	}
	defer scriptFile.Close()

	tmpScriptFile, err := os.Create(ctx.TmpScriptPath)
	if err != nil {
		return nil, err
	}
	defer tmpScriptFile.Close()

	err = tmpScriptFile.Chmod(0770)
	if err != nil {
		return nil, err
	}

	metaProc := processor.NewMetadataProcessor(ctx.ScriptPath)
	rewriterProc := processor.NewRewriterProcessor(tmpScriptFile, rewriters...)
	handler := handle.NewHandler(scriptFile, metaProc, rewriterProc)
	err = handler.Handle()
	if err != nil {
		return nil, err
	}

	ctx.Meta = metaProc.Metadata()

	ctx.NormTmpScriptPath = getNormalizedScriptPath(ctx.TmpScriptPath, ctx.Meta.Interpreter)
	ctx.InterpreterCmd = exec.Command(ctx.Meta.Interpreter)

	return ctx, nil
}

func teardownRun(ctx *RunCtx) error {
	return os.Remove(ctx.TmpScriptPath)
}

func callHookIfPresent(hook RunHookFn, ctx *RunCtx) error {
	if hook == nil {
		return nil
	}

	return hook(ctx)
}

func RunWithHooks(idOrPath string, hooks RunHooks, rewriters ...rewriter.Rewriter) (*RunCtx, error) {
	instFound, scriptPath := scriptmeta.GetScriptPathFromId(idOrPath)
	if !instFound {
		scriptPath = idOrPath
	}
	if !pathext.Exists(scriptPath) {
		return nil, fmt.Errorf("cannot find script '%s' by id or path", idOrPath)
	}

	ctx, err := setupRun(scriptPath, rewriters...)
	if err != nil {
		return nil, err
	}
	defer teardownRun(ctx)

	switch ctx.Meta.Interpreter {
	case "bash":
		prepCmdBash(ctx)
	default:
		return nil, fmt.Errorf("unsupported interpreter %s", ctx.Meta.Interpreter)
	}

	if err = callHookIfPresent(hooks.PreRun, ctx); err != nil {
		return nil, err
	}

	if err = ctx.InterpreterCmd.Run(); err != nil {
		return nil, err
	}

	if err = callHookIfPresent(hooks.PostRun, ctx); err != nil {
		return nil, err
	}

	return ctx, nil
}
