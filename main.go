package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/model"
	router "github.com/lughong/blog-service/internal/routers"
	"github.com/lughong/blog-service/pkg/setting"
)

func init() {
	setupSetting()
	setupDatabase()
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	r := router.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() {
	s := setting.NewSetting()

	if err := s.ReadSection("Server", &global.ServerSetting); err != nil {
		panic(err)
	}

	if err := s.ReadSection("App", &global.AppSetting); err != nil {
		panic(err)
	}

	if err := s.ReadSection("Database", &global.DatabaseSetting); err != nil {
		panic(err)
	}
}

func setupDatabase() {
	db, err := model.NewDBEngine()
	if err != nil {
		panic(err)
	}

	global.DBEngine = db
}
