##  pushGatewayTools 

### 介绍
prometheus文件目录监听

### 特性
- 基于 go mod 的依赖管理(国内源可使用：https://goproxy.cn/)

### 依赖
- 开发语言:`Golang ^1.16`

- 依赖软件包:
  ```go
  github.com/prometheus/client_golang v1.11.0
  github.com/robfig/cron v1.2.0
  github.com/spf13/viper v1.8.1
  golang.org/x/sys v0.0.0-20210819072135-bce67f096156 // indirect
  gopkg.in/ini.v1 v1.62.0
  ```

### 快速开始
```shell
$ git clone git@git.extremevision.com.cn:yumen/push_gateway_tools.git

$ cd push_gateway_tools

# 初始化环境变量
export INSTANCE_NAMESPACE="fsliu" ;export DEPLOY_NAME="fsliu123";export MODEL_PATH="/Users/fsliu/go/src/git.extremevision.com.cn/yumen/data_center/cmd" ;export PROMETHEUS_PUSH_GATEWAY_URL="http://192.168.67.62:9091"
# 下载依赖包
$ go mod vendor

# OR 使用go命令运行
$ go run main.go
```

### 
### 部署
```shell
docker build -t push_gateway_tools_go:v1 .

docker run -itd -v /data/go/runtime:/runtime  push_gateway_tools_go:v1
```

### 环境变量说明

| 字段                        | 类型   | 描述              | 例子                           |
| --------------------------- | ------ | ----------------- | ------------------------------ |
| INSTANCE_NAMESPACE          | string | 实例的命名空间    | platform                       |
| DEPLOY_NAME                 | string | deploy name       | platform-15-instance-51-ov-ide |
| MODEL_PATH                  | string | 需要监听的目录    | /home/data/model               |
| PROMETHEUS_PUSH_GATEWAY_URL | string | Push gateway 服务 | http://pushgateway:9091        |



