package conf

import (
	"ecloud_computer_auto_boot/pkg/util"
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Config = &config{}

type config struct {
	Server server `yaml:"server" mapstructure:"server"`
	Secret secret `yaml:"secret" mapstructure:"secret"`
	Cron   cron   `yaml:"cron" mapstructure:"cron"`
}

type server struct {
	Debug    bool   `yaml:"debug" mapstructure:"debug"`
	LogLevel string `yaml:"log-level" mapstructure:"log-level"`
}

type secret struct {
	Type      string `yaml:"type" mapstructure:"type"`
	Username  string `yaml:"username" mapstructure:"username"`
	Password  string `yaml:"password" mapstructure:"password"`
	AccessKey string `yaml:"access-key" mapstructure:"access-key"`
	SecretKey string `yaml:"secret-key" mapstructure:"access-key"`
	PoolId    string `yaml:"pool-id" mapstructure:"pool-id"`
}

type cron struct {
	Duration int      `yaml:"duration" mapstructure:"duration"`
	Machines []string `yaml:"machines" mapstructure:"machines"`
}

func Init() {
	var isFileNotFound = false
	var err error

	viper.SetDefault("server.debug", false)
	viper.SetDefault("server.log-level", "info")
	viper.SetDefault("secret.type", "public")
	viper.SetDefault("secret.username", "")
	viper.SetDefault("secret.password", "")
	viper.SetDefault("cron.duration", 60)
	viper.SetDefault("cron.machines", []string{})

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	var configFileNotFoundError viper.ConfigFileNotFoundError
	err = viper.ReadInConfig()
	isFileNotFound = errors.As(err, &configFileNotFoundError)
	if err != nil && !isFileNotFound {
		util.Log().Error("[配置初始化] 配置文件读取失败: %s", err)
		os.Exit(1)
	} else if isFileNotFound {
		err := viper.SafeWriteConfig()
		if err != nil {
			util.Log().Error("[配置初始化] 默认配置文件写出失败: %s", err)
			os.Exit(1)
		}
		util.Log().Info("[配置初始化] 已生成默认配置文件。")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `__`))
	viper.AutomaticEnv()

	err = viper.Unmarshal(Config)
	if err != nil {
		util.Log().Error("[配置初始化] 配置文件解析失败, error: %v", err)
	}

	if Config.Server.Debug {
		Config.Server.LogLevel = "debug"
	}

	// 重设log等级
	if Config.Server.LogLevel != "" {
		util.GlobalLogger = nil
		util.BuildLogger(Config.Server.LogLevel)
	}
}
