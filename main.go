package main

import (
	"net/http"
	"time"

	"github.com/lughong/blog-service/global"
	router "github.com/lughong/blog-service/internal/routers"
	"github.com/lughong/blog-service/pkg/setting"
)

func init() {
	setupConfig()
}

func main() {
	r := router.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerConfig.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupConfig() {
	s := setting.NewSetting()

	if err := s.ReadSection("Server", &global.ServerConfig); err != nil {
		panic(err)
	}

	if err := s.ReadSection("App", &global.AppConfig); err != nil {
		panic(err)
	}

	if err := s.ReadSection("Database", &global.DatabaseConfig); err != nil {
		panic(err)
	}
}
