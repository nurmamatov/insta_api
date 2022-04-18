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

type Options struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @BasePath /v1
// @version 1.0
// @description this is a user, post and comment services api

func NewRouter(option Options) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	urls := router.Group("/v1")
	urls.POST("/register", handlerV1.RegisterUser)
	urls.PUT("/login", handlerV1.Login)
	urls.PUT("/user/update", handlerV1.Update)
	urls.PUT("/user/delete", handlerV1.Delete)
	urls.GET("/search/:username", handlerV1.Search)
	urls.GET("/user/:username", handlerV1.GetUserPosts)
	urls.PUT("/user/password", handlerV1.ChangePassword)

	urls.POST("/post/create", handlerV1.CreatePost)
	urls.PUT("/post/get", handlerV1.GetPost)
	urls.PUT("/post/update", handlerV1.UpdatePost)
	urls.PUT("/post/delete", handlerV1.DeletePost)
	urls.GET("/list/:id/posts", handlerV1.ListUserPosts)
	urls.POST("/post/like", handlerV1.Like)
	urls.DELETE("/post/like", handlerV1.DeleteLike)

	urls.POST("/comment/create", handlerV1.CreateComment)
	urls.PUT("/comment/update", handlerV1.UpdateComment)
	urls.PUT("/comment/delete", handlerV1.DeleteComment)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	return router
}
