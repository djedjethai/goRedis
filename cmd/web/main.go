package main

import (
	// "encoding/gob"
	"log"
	"net/http"
	"os"
	// "reflect"
	"time"

	"github.com/djedjethai/goRedis/pkg/config"
	"github.com/djedjethai/goRedis/pkg/handlers"
	"github.com/gomodule/redigo/redis"
)

const goRedVersion = "1.0.4"
const redisConnection = "127.0.0.1:6379"
const staggerDelay = time.Second * 2
const maxWorkerPoolSize = 5
const maxJobMaxWorkers = 5

var app config.AppConfig
var handler *handlers.Handlers

var infoLog *log.Logger
var errorLog *log.Logger

func init() {
	// gob.Register()
	_ = os.Setenv("TZ", "UTC")
}

func main() {

	insecurePort, err := setupApp()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	rPool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
		// Dial: func() (redis.Conn, error) {
		// 	c, err := redis.Dial("tcp", ":6379")
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// 	return c, err
		// },
	}

	// conn := rPool.Get()
	// defer conn.Close()

	app.Pool = rPool

	// create http server
	srv := &http.Server{
		Addr:     *insecurePort,
		ErrorLog: errorLog,
		// Handler:           routes(app),
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("Starting HTTP server on port %s....", *insecurePort)

	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}
