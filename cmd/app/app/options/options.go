package options

import (
	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/middleware/log"
	"ginkgo/internal/server"
)

const (
	_defaultConfigFile = "config/config.yaml"
)

type AppOptions struct {
	ConfFile string
	Config   *config.Config
}

func NewAppOptions() *AppOptions {
	o := &AppOptions{}
	return o
}

func (o *AppOptions) NewServer() (*server.Server, error) {
	s := server.NewServer()
	o.loadConfig(o.ConfFile)
	s.Config = o.Config
	s.Log = log.NewClient(s.Config.Log)
	s.DB = database.NewClient(&s.Config.Db)
	return s, nil
}

func (o *AppOptions) loadConfig(configFile string) {
	if configFile == "" {
		configFile = _defaultConfigFile
	}
	o.Config = config.New(configFile)
	if o.Config.Log.LogDirector == "" {
		o.Config.Log.LogDirector = "./log"
	}
	o.Config.Log.LogErrorFilePath = o.Config.Log.LogDirector + "/" + o.Config.Log.LogErrorFilename
	o.Config.Log.LogInfoFilePath = o.Config.Log.LogDirector + "/" + o.Config.Log.LogInfoFilename

}
