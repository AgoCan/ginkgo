package config

// 配置文件导入yaml文件是configstruct.go
//
// 配置文件可以使用 -c 的参数
// https://github.com/go-yaml/yaml

import (
	"path"
	"runtime"

	"github.com/spf13/viper"

	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/file"
	"ginkgo/internal/pkg/jwt"
	"ginkgo/internal/pkg/middleware/log"
)

func New(loadPath string) *Config {
	var conf *Config
	viper.SetConfigFile(loadPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// Config 配置文件
type Config struct {
	Db   database.Options
	Log  log.Options
	Jwt  jwt.Options
	File file.Options
}

// 获取文件绝对路径
func GetCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
