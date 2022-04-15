package main

import (
	"tasks/Instagram_clone/insta_api/api"
	"tasks/Instagram_clone/insta_api/config"
	"tasks/Instagram_clone/insta_api/pkg/logger"
	"tasks/Instagram_clone/insta_api/services"
)

func main() {

	cfg := config.LoadConfig()
	log := logger.New("develop", "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg.Services)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.NewRouter(api.Options{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.Lpe.Port); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
