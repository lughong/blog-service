package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/model"
	router "github.com/lughong/blog-service/internal/routers"
	"github.com/lughong/blog-service/pkg/file"
	"github.com/lughong/blog-service/pkg/logger"
	"github.com/lughong/blog-service/pkg/setting"
)

func init() {
	if err := interRootDir(); err != nil {
		log.Fatalf("interRootDir error. %v", err)
	}

	if err := setupSetting(); err != nil {
		log.Fatalf("setupSetting error. %v", err)
	}

	if err := setupDBEngine(); err != nil {
		log.Fatalf("setupDBEngine error. %v", err)
	}

	if err := setupLogger(); err != nil {
		log.Fatalf("setupLogger error. %v", err)
	}
}

// @title Blog-service example API
// @version 1.0
// @description This is a sample blog service.
// @termsOfService https://github.com/lughong/blog-service
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
	s := setting.NewSetting(global.RootDir)

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
			global.RootDir+"/"+global.AppSetting.LogSavePath,
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

func interRootDir() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	f := file.New()

	var inter func(d string) string
	inter = func(d string) string {
		if isExists := f.PathExists(d + "/configs"); !isExists {
			return d
		}

		return inter(filepath.Dir(d))
	}

	global.RootDir = inter(cwd)

	return nil
}
