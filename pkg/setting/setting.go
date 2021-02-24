package setting

import (
	"log"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(rootPath string) *Setting {
	vp := viper.New()
	vp.AddConfigPath("/etc/blog-service")
	vp.AddConfigPath("$HOME/.blog-service")
	vp.AddConfigPath(rootPath + "/configs")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig error. %s", err)
	}

	return &Setting{
		vp,
	}
}
