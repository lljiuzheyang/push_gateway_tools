package file

import (
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/file"
	file2 "git.extremevision.com.cn/yumen/push_gateway_tools/pkg/util/file"
	"encoding/json"
	"log"
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/logging"
	"strings"
	"strconv"
)

func Monitor() map[string]interface{} {
	var err error
	data := make(map[string]interface{})
	path := "/Users/fsliu/Documents/work/company/jsj/data_set"

	allFile, err := file.GetAllFile(path)
	if err != nil {
		return data
	}
	node := file2.NewNode("/", "0", "/")
	for i := 0; i < len(allFile); i++ {
		filePath := strings.Replace(allFile[i], path, "", -1)
		node.Insert(filePath, strconv.Itoa(i))

	}
	res, err := json.Marshal(node)
	if err != nil {
		log.Fatal(err)
	}
	logging.Warn(string(res))
	data["dir"] = string(res)
	return data
}
