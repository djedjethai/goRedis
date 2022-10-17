package handlers

import (
	"context"
	"fmt"
	"net/http"
	// "time"

	"github.com/djedjethai/goRedis/pkg/config"
	"github.com/djedjethai/goRedis/pkg/internal/helpers"
	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/go-chi/chi/v5"
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

func (h *Handlers) CreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := models.Item{}

		err := h.Json.ReadJson(w, r, &item)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		str, err := h.createItem(item)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		fmt.Println("grrr: ", item)

		err = h.Json.WriteJson(w, http.StatusOK, str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) GetItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		ctx := context.Background()
		item := h.getItem(ctx, id)

		fmt.Println("see the item: ", item)

		h.Json.WriteJson(w, http.StatusOK, item)
	}
}

func (h *Handlers) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")

		ctx := context.Background()
		user := h.getUserByID(ctx, id)

		fmt.Println("see the user: ", user)

		h.Json.WriteJson(w, http.StatusOK, user)
	}
}

func (h *Handlers) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ooookkk, lets work")

	}
}
