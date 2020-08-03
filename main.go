package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/model"
	router "github.com/lughong/blog-service/internal/routers"
	"github.com/lughong/blog-service/pkg/logger"
	"github.com/lughong/blog-service/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("setupSetting error. %s", err)
	}

	//if err := setupDBEngine(); err != nil {
	//	log.Fatalf("setupDBEngine error. %s", err)
	//}

	if err := setupLogger(); err != nil {
		log.Fatalf("setupLogger error. %s", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	r := router.NewRouter()

	srv := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	/*
		global.Logger.Debug("test debug")
		global.Logger.Debugf("test debug %s", "bbb")
		global.Logger.Info("test info")
		global.Logger.Infof("test info %s", "bbb")
		global.Logger.Error("test error")
		global.Logger.Errorf("test error %s", "bbb")
		global.Logger.Fatal("test fatal")
		global.Logger.Fatalf("test fatal %s", "bbb")
		global.Logger.Panic("test panic")
		global.Logger.Panicf("test panic %s", "bbb")
	*/

	srv.ListenAndServe()
}

func setupSetting() error {
	s := setting.NewSetting()

	if err := s.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}

	if err := s.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}

	if err := s.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	db, err := model.NewDBEngine()
	if err != nil {
		return err
	}

	global.DBEngine = db
	return nil
}

func setupLogger() error {
	l := &lumberjack.Logger{
		Filename: fmt.Sprintf(
			"%s/%s.%s",
			global.AppSetting.LogSavePath,
			global.AppSetting.LogFileName,
			global.AppSetting.LogFileExt,
		),
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	global.Logger = logger.NewLogger(
		l,
		"",
		log.Ldate|log.Ltime|log.Lshortfile,
	).WithCaller(2)

	return nil
}
