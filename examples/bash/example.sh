#!/usr/bin/env bash

# @sman namespace marcex
# @sman name exbash

echo "common code before"

# @sman sec:start getargs
echo -n 'set NEWENV value: '
read NEWENV
echo "getargs called, NEWENV set to $NEWENV"
# @sman sec:end getargs

echo "common code between"

# @sman sec:start run
echo "run called, NEWENV is: $NEWENV"
# @sman sec:end run

echo "common code after"