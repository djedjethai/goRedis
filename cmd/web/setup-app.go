package main

import (
	"flag"

	"github.com/djedjethai/goRedis/pkg/config"
	"github.com/djedjethai/goRedis/pkg/handlers"
	"log"
	"os"
)

func setupApp() (*string, error) {

	insecurePort := flag.String("port", ":4000", "port to listen on")

	// create info log -- just writes to Stdout
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	a := config.AppConfig{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Version:  goRedVersion,
	}

	// TODO verif with watcher
	// rPool := &redis.Pool{
	// 	MaxIdle:   80,
	// 	MaxActive: 12000, // max number of connections
	// 	Dial: func() (redis.Conn, error) {
	// 		c, err := redis.Dial("tcp", ":6379")
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		return c, err
	// 	},
	// }

	// conn := rPool.Get()
	// defer conn.Close()

	// a.Conn = conn
	app = a

	hand := handlers.NewHandlers(&app)
	handler = hand

	return insecurePort, nil
}
