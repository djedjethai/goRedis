package config

import (
	"log"

	// "github.com/djedjethai/goRedis/pkg/internal"
	"github.com/gomodule/redigo/redis"
)

type AppConfig struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Version  string
	Pool     *redis.Pool
}
