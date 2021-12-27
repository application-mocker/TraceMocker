package main

import (
	"TraceMocker/config"
	"TraceMocker/internal/web"
	"TraceMocker/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


func init() {

	flags()

	utils.Logger.Info("Bootstrap application")
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
}

func main() {
	web.StartHttpServer()
}
