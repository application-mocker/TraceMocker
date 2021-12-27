package web

import (
	"TraceMocker/config"
	"TraceMocker/internal/web/handler"
	"TraceMocker/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func innerInit() {

	// set mode
	utils.Logger.Infof("Init HttpServer with mode: {%s}", config.Config.Application.Mode)
	if config.Config.Application.Mode == utils.ApplicationTestMode {
		gin.SetMode(gin.TestMode)
	} else if config.Config.Application.Mode == utils.ApplicationReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func StartHttpServer() {

	innerInit()

	server := gin.Default()
	utils.Logger.Trace("Add routes...")

	server.Any("/trace-mock/simple", handler.SimpleHandler)

	utils.Logger.Infof("Start the http server with port: {%s}", config.Config.HttpServer.Port)
	err := server.Run(config.Config.HttpServer.Port)
	if err != nil {
		log.Fatal(err)
	}
}
