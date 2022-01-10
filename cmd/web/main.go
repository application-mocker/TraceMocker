package main

import (
	"TraceMocker/config"
	"TraceMocker/internal/task"
	"TraceMocker/internal/web"
	"TraceMocker/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func init() {

	flags()
	task.Init()
	utils.Logger.Info("Bootstrap application")

	task.InitProcessor()

	cronInstance := cron.New()
	cronInstance.Start()

	if config.Config.Application.ObjectClientConfig.Enable {
		_, err := cronInstance.AddFunc("@every 15s", func() {
			fmt.Println("Start sync ")
			task.ProcessorInstance.Sync()
		})
		if err != nil {
			panic(err)
		}
	}
}

func flags() {

	configPath := os.Getenv("CONF_PATH")
	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(file, config.Config); err != nil {
		panic(err)
	}

	NodeId := config.Config.Application.NodeId
	if id := os.Getenv("NODE_ID"); len(id) != 0 {
		NodeId = id
	}

	config.Config.Application.NodeId = NodeId
}

func main() {
	web.StartHttpServer()
}
