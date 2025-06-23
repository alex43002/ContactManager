#!/bin/bash

if [ $# -eq 0 ]; then
	echo "Usage: $0 <executable> [args...]"
	exit 1
fi

echo "Timing: $@"

time "$@"
