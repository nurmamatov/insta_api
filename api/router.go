package api

import (
	v1 "tasks/Instagram_clone/insta_api/api/handlers/v1"
	"tasks/Instagram_clone/insta_api/config"
	"tasks/Instagram_clone/insta_api/pkg/logger"
	"tasks/Instagram_clone/insta_api/services"
	_ "tasks/Instagram_clone/insta_api/api/docs"


	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Conf            config.Config
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
}

// @BasePath /v1
// @version 1.0
// @description this is a user and task services api
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,

	})

	task := router.Group("/v1")
	task.POST("/register", handlerV1.RegisterUser)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	return router
}
