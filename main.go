package main

import (
	"shop-backend/api"
	"shop-backend/config"
	"shop-backend/utils/logger"
)

func main() {
	server := api.Init()

	err := server.Start(":" + config.Cfg.Port)
	if err != nil {
		logger.L.Fatalln(err)
	}
}
