package parser_module

import (
	"kbell/lib_module"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {

	Logger = lib_module.GetLogInstance(true, "../logs/run.log")
	Logger.Info("parser_module package init >>> ")
}
