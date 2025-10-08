package usercenter

import (
	"github.com/chhz0/usercenter-go/internal/pkg/conf"
	"github.com/chhz0/usercenter-go/internal/pkg/options"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	appName     = "usercenter"
	defaultMode = "dev"

	defaultConfigName = "usercenter"
	defaultConfigType = "yaml"
	defaultConfigPath = "./configs"
	defaultEnvPrefix  = "USERCENTER"
)

type Config struct {
	App  string               `json:"app" mapstructure:"app"`
	Mode string               `json:"mode" mapstructure:"mode"`
	HTTP *options.HTTPOptions `json:"http" mapstructure:"http"`

	V *viper.Viper
}

func (c *Config) BindFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.Mode, "mode", "dev", "server mode")
}

func (c *Config) Server() error {
	// todo: server choice
	httpSrv := newGinHTTPServer(c.HTTP)
	_ = httpSrv.ListenAndServe()

	return nil
}

func (c *Config) Load() {
	conf.EnableEnv("USERCENTER", true, c.V)

	if ffile, ok := c.V.Get(conf.FlagName).(string); ok && ffile != "" {
		c.V.SetConfigFile(ffile)
	}

	if err := c.V.ReadInConfig(); err != nil {
		zap.L().Fatal("read config error", zap.Error(err))
	}

	if err := c.V.Unmarshal(c); err != nil {
		zap.L().Fatal("unmarshal config error", zap.Error(err))
	}

	go func() {
		c.V.WatchConfig()
		c.V.OnConfigChange(func(in fsnotify.Event) {
			zap.L().Warn("config file changed")
			if err := c.V.Unmarshal(c); err != nil {
				zap.L().Fatal("read config error", zap.Error(err))
			}
		})
	}()
}

func newConf() *Config {
	return &Config{
		App:  appName,
		Mode: defaultMode,
		HTTP: options.NewHTTPOptions(),

		V: conf.InitViper(defaultConfigName, defaultConfigType, defaultConfigPath),
	}
}
