package conf

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	FlagName = "config"
)

func InitViper(fname, ftype string, fpaths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(fname)
	v.SetConfigType(ftype)
	for _, p := range fpaths {
		v.AddConfigPath(p)
	}
	return v
}

func EnableEnv(prefix string, allowEmpty bool, v *viper.Viper) {
	v.AutomaticEnv()
	v.SetEnvPrefix(prefix)
	v.AllowEmptyEnv(allowEmpty)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func BindConfigFlag(fs *pflag.FlagSet, v *viper.Viper) {
	var fullfile string
	fs.StringVar(&fullfile, FlagName, "", "config file, e.g. -config={$filepath}/config.yaml")
	_ = v.BindPFlags(fs)
}
