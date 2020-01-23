#!/bin/bash
SOURCE=${SOURCE:-$1}
SOURCE=${SOURCE:-sources.yml}
ENVT=${ENVT:-env-tester}
if [ -e "$SOURCE" ]; then
	$ENVT source execute --source $SOURCE
else
	echo "$SOURCE does not exists"
fi
