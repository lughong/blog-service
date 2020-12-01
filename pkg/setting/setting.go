package setting

import (
	"log"

	"github.com/lughong/blog-service/global"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() *Setting {
	vp := viper.New()
	vp.AddConfigPath("/etc/blog-service")
	vp.AddConfigPath("$HOME/.blog-service")
	vp.AddConfigPath(global.RootDir + "configs")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig error. %s", err)
	}

	return &Setting{
		vp,
	}
}
