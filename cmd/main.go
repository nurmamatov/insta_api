package main

import (
	"tasks/Instagram_clone/insta_api/api"
	"tasks/Instagram_clone/insta_api/config"
	"tasks/Instagram_clone/insta_api/pkg/logger"
	"tasks/Instagram_clone/insta_api/services"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	serviceManager.UserService()
	if err := server.Run(cfg.Port); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
