package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Enivorentment    string

	PostServiceHost  string
	PostServicePort  int

	UserServiceHost  string
	UserServicePort  int
	
	CommentServiceHost string
	CommentServicePort int

	LogLevel string
	Port string

}

func Load() Config {
	c := Config{}

	c.Enivorentment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.Port = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	
	c.PostServiceHost = cast.ToString(getOrReturnDefault("Post_SERVICE_HOST", "post-services"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("Post_SERVICE_PORT", 9000))
	
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "assignee-services"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9002))
	
	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "comment_service"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 9001))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
