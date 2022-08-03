package lib_module

import (
	"os/exec"
)

//ref exec : https://mingrammer.com/gobyexample/spawning-processes/
//ref exec : https://stackoverflow.com/questions/52280905/golang-exec-command-bash-command-not-working

func LoadCmd(cmd string) (string, bool) {

	// 시스템 명령어 수행 및 결과반환

	var flag_Res bool = true

	cmdRes := exec.Command("bash", "-c", cmd)

	cmdOut, err := cmdRes.Output()

	if err != nil {
		Logger.Error("error: %v", err)
		flag_Res = false
	}

	//Logger.Info("cmd : ", cmd, " , res : ", string(cmdOut))

	return string(cmdOut), flag_Res
}
