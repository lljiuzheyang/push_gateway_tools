package setting

import (
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"log"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	if err = viper.BindEnv("INSTANCE_NAMESPACE"); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("DEPLOY_NAME"); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("MODEL_PATH"); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("PROMETHEUS_PUSH_GATEWAY_URL"); err != nil {
		panic(err)
	}
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
