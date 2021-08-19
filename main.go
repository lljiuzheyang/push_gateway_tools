package main

import (
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/setting"
	"fmt"
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/logging"
	"github.com/robfig/cron"
	"git.extremevision.com.cn/yumen/push_gateway_tools/file"
)

func init() {
	setting.Setup()
	logging.Setup()
}


func main()  {
	cron := cron.New()
	//执行定时任务（每10秒执行一次）
	err:= cron.AddFunc("*/10 * * * * *", cmd)
	if err!=nil{
		fmt.Println(err)
	}
	cron.Start()
	defer cron.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

func cmd()  {
	fmt.Println("每10s执行一次cron")
	file.Monitor()
}
