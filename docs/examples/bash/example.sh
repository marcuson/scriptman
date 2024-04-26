#!/usr/bin/env bash

# @scriptman namespace marcex
# @scriptman name exbash

echo "common code before"

# @scriptman sec:start getargs
echo -n 'set NEWENV value: '
read NEWENV
echo "getargs called, NEWENV set to $NEWENV"
# @scriptman sec:end getargs

echo "common code between"

# @scriptman sec:start run
echo "run called, NEWENV is: $NEWENV"
# @scriptman sec:end run

echo "common code after"