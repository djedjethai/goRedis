package handlers

import (
	// "context"
	"fmt"
	"net/http"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	// "github.com/gomodule/redigo/redis"
	// "github.com/google/uuid"
)

type Views struct {
	ItemID string `json:"item_id"`
	UserID string `json:"user_id"`
}

// type Response struct {
// 	ID string `json:"id"`
// }

func (h *Handlers) IncrementViews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vs := Views{}

		err := h.Json.ReadJson(w, r, &vs)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		str, err := h.incrementViews(vs.ItemID, vs.UserID)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		vs.UserID = str

		err = h.Json.WriteJson(w, http.StatusOK, vs)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) incrementViews(itemID, userID string) (str string, err error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	conn.Send("MULTI")
	conn.Send("HINCRBY", models.ItemsKey(itemID), "views", 1)
	conn.Send("ZINCRBY", models.ItemsByViewsKey(), 1, itemID)
	pipe_prox, err := conn.Do("EXEC")
	if err != nil {
		return "", err
	}

	fmt.Println("see the res pipeline incrementViews: ", pipe_prox)
	return itemID, nil
}
