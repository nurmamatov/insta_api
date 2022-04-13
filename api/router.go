package api

import (
	_ "tasks/Instagram_clone/insta_api/api/docs"
	v1 "tasks/Instagram_clone/insta_api/api/handlers/v1"
	"tasks/Instagram_clone/insta_api/config"
	"tasks/Instagram_clone/insta_api/pkg/logger"
	"tasks/Instagram_clone/insta_api/services"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @BasePath /v1
// @version 1.0
// @description this is a user, post and comment services api

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
	task.GET("/login", handlerV1.Login)
	task.PUT("/user/update", handlerV1.Update)
	task.PUT("/user/delete", handlerV1.Delete)
	task.GET("/search/user", handlerV1.Search)
	task.GET("/user/get", handlerV1.GetUserPosts)

	task.POST("/post/create", handlerV1.CreatePost)
	task.GET("/post/get", handlerV1.GetPost)
	task.PUT("/post/update", handlerV1.UpdatePost)
	task.PUT("/post/delete", handlerV1.DeletePost)
	task.GET("/list/user/posts", handlerV1.ListUserPosts)
	task.POST("/post/like", handlerV1.Like)
	task.DELETE("/post/like", handlerV1.DeleteLike)

	task.POST("/comment/create", handlerV1.CreateComment)
	task.GET("/comment/get", handlerV1.GetComment)
	task.PUT("/comment/update", handlerV1.UpdateComment)
	task.PUT("/comment/delete", handlerV1.DeleteComment)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	return router
}
