#!/bin/bash

current_dir=`pwd`
cd $current_dir
cd ..
project_path=`pwd`
project_bin_name=cc_detector

export GOPATH=$project_path

echo "========================================================"
echo "Project Path >> "
echo ": "$project_path
echo "GOROOT >>"
echo ": "$GOROOT
echo "GOPATH >>"
echo ": "$GOPATH
echo "========================================================"
echo "Proejct Build(kbell) >>"
cd $project_path/src
go build .
ls -al
echo "========================================================"
echo "Move Binary(kbell) >> "
mv $project_path/src/kbell $project_path/bin/$project_bin_name
echo "========================================================"
