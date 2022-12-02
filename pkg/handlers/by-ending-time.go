package handlers

import (
	"context"
	"net/http"
	// "encoding/json"
	"fmt"
	// "strings"

	// "strconv"
	// "time"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/gomodule/redigo/redis"
)

func (h *Handlers) ItemsByEndingTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()

		itc := struct {
			Order  string `json:"order"`
			Offset int    `json:"offset"`
			Count  int    `json:"count"`
		}{}

		err := h.Json.ReadJson(w, r, &itc)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		str, err := h.itemsByEndingTime(ctx, itc.Order, itc.Offset, itc.Count)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		err = h.Json.WriteJson(w, http.StatusOK, str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) itemsByEndingTime(ctx context.Context, order string, offset, count int) ([]models.Item, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	var items = []models.Item{}

	// get(from a sorted set dataStructure) the 2 first items with score from time.Now()(or 0 for 2nd query) until +inf
	// values, err := redis.Values(conn.Do("ZRANGE", models.ItemsByEndingAtKey(), time.Now().Unix(), "+inf", "BYSCORE", "LIMIT", 0, 2))
	values, err := redis.Values(conn.Do("ZRANGE", models.ItemsByEndingAtKey(), 0, "+inf", "BYSCORE", "LIMIT", 0, 2))
	if err != nil {
		fmt.Println("Err ZRANGE in ItemsByEndingTime: ", err)
		return items, err
	}

	var ids []string
	for _, v := range values {
		ids = append(ids, string(v.([]byte)))
	}

	items, err = h.getItems(ids)
	if err != nil {
		fmt.Println("err get items from ItemsByEndingTime: ", err)
	}

	// conn.Send("multi")
	return items, nil
}
