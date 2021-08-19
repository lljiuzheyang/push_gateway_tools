package main

import (
	"git.extremevision.com.cn/yumen/push_gateway_tools/file"
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/setting"
	"fmt"
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/logging"
)

func init() {
	setting.Setup()
	logging.Setup()
}


func main()  {
	data :=file.Monitor()
	fmt.Println(data)
}
