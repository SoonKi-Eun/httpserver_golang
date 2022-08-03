package main

import (
	"kbell/http_module"
	"kbell/lib_module"
	"kbell/struct_module"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger
var Config *struct_module.Config_ST

func main() {
	proc_init()
	http_module.HttpServerStart()
}

func proc_init() {

	// Max cpu conf set

	if Config.Project_Config.Max_Cpu_Use > 0 {
		runtime.GOMAXPROCS(Config.Project_Config.Max_Cpu_Use)
	} else {
		Logger.Warn("Config max_cpu_use(", Config.Project_Config.Max_Cpu_Use, ") wrong , force setup max_cpu_use = 1")
		runtime.GOMAXPROCS(1)
	}
}

func init() {

	Logger = lib_module.GetLogInstance(true, "../logs/run.log")
	Config = lib_module.GetConfigInstance(true, "../config/config.yaml")

	Logger.Info("main package init >>> ")
}
