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
go get github.com/sirupsen/logrus
go get gopkg.in/natefinch/lumberjack.v2
go get github.com/robfig/cron/v3
go get github.com/gookit/config/v2
go get github.com/gookit/config/v2/yaml
go get github.com/robfig/cron
go get github.com/labstack/echo
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
