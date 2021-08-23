package main

import (
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/setting"
	"fmt"
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/logging"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"git.extremevision.com.cn/yumen/push_gateway_tools/file"
)

func init() {
	setting.Setup()
	logging.Setup()
}

func main() {
	cron2 := cron.New()
	//执行定时任务（每10秒执行一次）
	err := cron2.AddFunc("*/1 * * * * *", cmd)
	if err != nil {
		fmt.Println(err)
	}
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

func cmd() {
	fmt.Println("每1s执行一次cron")
	fmt.Println(viper.GetString("INSTANCE_NAMESPACE"))
	fmt.Println(viper.GetString("DEPLOY_NAME"))
	fmt.Println(viper.GetString("MODEL_PATH"))
	fmt.Println(viper.GetString("PROMETHEUS_PUSH_GATEWAY_URL"))
	file.Monitor()
}
