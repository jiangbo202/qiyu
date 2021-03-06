/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/06/17 10:02 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"jiangbo.com/blog_service/global"
	"jiangbo.com/blog_service/internal/model"
	"jiangbo.com/blog_service/internal/routers"
	"jiangbo.com/blog_service/pkg"
	"jiangbo.com/blog_service/pkg/logger"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {

	models := []interface{}{&model.Article{}, &model.Tag{}, &model.Category{}, &model.Auth{}}
	err := global.DBEngine.AutoMigrate(models...).Error
	if err != nil {
		panic("db迁移失败")
	}

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	//log.Printf("%v", global.ServerSetting)
	//global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := pkg.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
