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

	// mock method
	server.Any("/mock/http/status-code/code/:code", handler.SingleCodeHandler)
	server.Any("/mock/http/status-code/codes", handler.RandomCode)
	server.Any("/mock/http/response-size/size/:size", handler.SingleResponseSizeHandler)
	server.Any("/mock/http/ping", handler.PingHandler)
	server.Any("/mock/http/simple-request", handler.SimpleTransHandler)
	server.Any("/mock/http/trace-mock", handler.TraceHandler)

	server.GET("/tasks", handler.ListTask)
	server.POST("/task", handler.RegisterTask)
	server.DELETE("/task")

	utils.Logger.Infof("Start the http server with port: {%s}", config.Config.HttpServer.Port)
	err := server.Run(config.Config.HttpServer.Port)
	if err != nil {
		log.Fatal(err)
	}
}
