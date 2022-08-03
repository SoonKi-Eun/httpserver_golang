package lib_module

import (
	"kbell/struct_module"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger
var Lum *lumberjack.Logger
var Config *struct_module.Config_ST

var log_inst_once sync.Once
var conf_inst_once sync.Once

func init() {

	Logger = GetLogInstance(true, "../logs/run.log")
	Config = GetConfigInstance(true, "../config/config.yaml")

	Logger.Info("lib_module package init >>> ")
}
