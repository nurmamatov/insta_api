package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Services Services
	Lpe      ConfigP
}
type Services struct {
	PostServiceHost string
	PostServicePort int

	UserServiceHost string
	UserServicePort int

	CommentServiceHost string
	CommentServicePort int
}
type ConfigP struct {
	Enivorentment string
	LogLevel      string
	Port          string
}

func LoadConfig() Config {
	c := Config{}

	c.Lpe.Enivorentment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.Lpe.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.Lpe.Port = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	c.Services.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post_services"))
	c.Services.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 9000))

	c.Services.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user_service"))
	c.Services.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9002))

	c.Services.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "comment_service"))
	c.Services.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 9001))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
