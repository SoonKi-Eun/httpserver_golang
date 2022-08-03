package lib_module

import (
	"fmt"
	"os"
)

func AWS_GetInstMetric_SaveJsonFile(json_data string, project_path string, json_file string) {

	// json_file 에 json_data 저장 , overwrite
	file_path := fmt.Sprintf("%s/tmp/%s.json", project_path, json_file)
	fd, err := os.Create(file_path)

	if err != nil {
		panic(err)
	}
	defer fd.Close()

	fmt.Fprint(fd, json_data)

}

func FileAppend(file_path string, data string) {

	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		Logger.Error("failed opening file :", err)
	}

	defer file.Close()

	bytes, err := file.WriteString(data)

	if err != nil {
		Logger.Error("failed writing to file (", bytes, "): ", err)
	}
}

func FileRemove(file_path string) {

	e := os.Remove(file_path)

	if e != nil {
		Logger.Error("file remove fail : ", e)
	}

}
