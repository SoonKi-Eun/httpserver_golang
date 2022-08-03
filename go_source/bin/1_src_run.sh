#!/bin/bash

current_dir=`pwd`
cd $current_dir
cd ..
project_path=`pwd`

export GOPATH=$project_path

echo "========================================================"
echo "Project Path >> "
echo ": "$project_path
echo "GOROOT >>"
echo ": "$GOROOT
echo "GOPATH >>"
echo ": "$GOPATH
echo "========================================================"
echo "Proejct Go Run (ver. interpreter) >>"
cd $project_path/src
go run ./go_main.go $1 $2 $3

