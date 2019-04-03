package main

import (
	"log"
	"os"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
)

const (
	logLevel   = "ERROR"
	configFile = "configs/krakend-group.json"
)

func main() {
	parser := config.NewParser()
	serviceConfig, err := parser.Parse(configFile)
	logger, err := logging.NewLogger(logLevel, os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	routerFactory := krakendgin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	routerFactory.New().Run(serviceConfig)
}
