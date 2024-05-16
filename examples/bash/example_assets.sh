#!/usr/bin/env bash

# @scriptman namespace marcex
# @scriptman asset assets/**

# @scriptman sec:start run
SCRIPT_DIR_PATH=$(readlink -f "$0" | xargs dirname)

echo "Showing content of asset 'assets/info.txt':"
cat "$SCRIPT_DIR_PATH/assets/info.txt"
# @scriptman sec:end run