package lib_module

import (
	"io/ioutil"
	"kbell/struct_module"
	"runtime"

	"gopkg.in/yaml.v2"
)

func LoadConfigration(out_flag bool, path string) *struct_module.Config_ST {

	config := struct_module.Config_ST{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		Logger.Error("error: %v", err)
	}

	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		Logger.Error("error: %v", err)
	}

	if out_flag {
		ShowConfigration(&config)
	}

	return &config
}

func GetConfigInstance(out_flag bool, path string) *struct_module.Config_ST {

	conf_inst_once.Do(func() {

		Config = LoadConfigration(out_flag, path)
	})

	return Config
}

func ShowConfigration(c *struct_module.Config_ST) {

	Logger.Info("==============================================================")
	Logger.Info("project >>>>>>>>>>")
	Logger.Info(" total_cpu               : ", runtime.NumCPU())
	Logger.Info(" porject_name            : ", c.Project_Config.Project_name)
	Logger.Info(" porject_path            : ", c.Project_Config.Project_path)
	Logger.Info("http    >>>>>>>>>>")
	Logger.Info(" port                    : ", c.HttpServer_Config.Port)
	Logger.Info(" timeout                 : ", c.HttpServer_Config.TimeOut)
	Logger.Info(" ssl_active              : ", c.HttpServer_Config.SSL_Active)
	Logger.Info(" private_key_path        : ", c.HttpServer_Config.PrivateKey_path)
	Logger.Info(" crt_certificate_path    : ", c.HttpServer_Config.CrtCert_path)
	Logger.Info("==============================================================")
}
