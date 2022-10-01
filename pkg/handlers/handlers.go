package handlers

import (
	"fmt"
	"net/http"
	// "time"

	"github.com/djedjethai/goRedis/pkg/config"
	"github.com/djedjethai/goRedis/pkg/internal/helpers"
	"github.com/djedjethai/goRedis/pkg/internal/models"
	// "github.com/gomodule/redigo/redis"
)

type Handlers struct {
	app   *config.AppConfig
	Json  helpers.Json
	token models.Token
}

func NewHandlers(a *config.AppConfig) *Handlers {
	return &Handlers{
		app: a,
		// pool:  a.Pool,
		Json:  helpers.Json{},
		token: models.Token{},
	}
}

func (h *Handlers) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ooookkk, lets work")
		conn := h.app.Pool.Get()
		defer conn.Close()
	}
}
