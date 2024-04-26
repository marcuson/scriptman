#!/usr/bin/env bash

# @scriptman namespace marcex
# @scriptman name exbash

# @scriptman sec:start run
echo "run called, NEWENV is: $NEWENV, should be 'my_newenv_var' (if .env loaded correctly)"
echo "with envsubst command: $(envsubst <<< \"$NEWENV\")"
# @scriptman sec:end run