package lib_module

import (
	"sort"
)

func Create_EmptyStringBinaryList() []string {
	empty_list := []string{}
	return empty_list
}

func Insert_StringBinaryList(binary_list []string, target string, duplicate_allow ...bool) []string {

	var exist bool = false

	if len(duplicate_allow) == 0 {

		//default , not duplicate
		exist = Search_StringBinaryList(binary_list, target)

	} else {

		if !duplicate_allow[0] {
			//option = false , not duplicate
			exist = Search_StringBinaryList(binary_list, target)
		}
	}

	if !exist {

		//Logger.Info("Add new data : ", target)
		i := sort.SearchStrings(binary_list, target)
		binary_list = append(binary_list, "")
		copy(binary_list[i+1:], binary_list[i:])
		binary_list[i] = target

	} else {

		//Logger.Info("Already exist data : ", target)
	}

	return binary_list

}

func Search_StringBinaryList(binary_list []string, target string) bool {

	var res_flag = false

	i := sort.Search(len(binary_list), func(i int) bool { return target <= binary_list[i] })

	if i < len(binary_list) && binary_list[i] == target {
		//Logger.Info("Found %s at index %d in %v.\n", target, i, binary_list)
		res_flag = true
	} else {
		//Logger.Info("Did not find %s in %v.\n", target, binary_list)
		res_flag = false
	}
	return res_flag
}
