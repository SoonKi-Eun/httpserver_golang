package http_module

import (
	"kbell/lib_module"
	"kbell/struct_module"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger
var Config *struct_module.Config_ST

func init() {

	Logger = lib_module.GetLogInstance(true, "../logs/run.log")
	Config = lib_module.GetConfigInstance(true, "../config/config.yaml")
	Logger.Info("http_module package init >>> ")
}
