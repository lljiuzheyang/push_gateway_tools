package file

import (
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/file"
	file2 "git.extremevision.com.cn/yumen/push_gateway_tools/pkg/util/file"
	"encoding/json"
	"log"
	"git.extremevision.com.cn/yumen/push_gateway_tools/pkg/logging"
	"strings"
	"strconv"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/spf13/viper"
)

func Monitor() map[string]interface{} {
	var err error
	data := make(map[string]interface{})
	path := viper.GetString("MODEL_PATH")

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
	Collector(string(res))
	data["dir"] = string(res)
	return data
}

func Collector(resJson string) {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_model_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	prometheusPushGatewayUrl := viper.GetString("PROMETHEUS_PUSH_GATEWAY_URL")
	deployName := viper.GetString("DEPLOY_NAME")
	namespace := viper.GetString("INSTANCE_NAMESPACE")

	completionTime.SetToCurrentTime()
	if err := push.New(prometheusPushGatewayUrl, "db_backup_model").
		Collector(completionTime).
		Grouping("deploy_name", deployName).
		Grouping("namespace", namespace).
		Grouping("model_directory", resJson).
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
