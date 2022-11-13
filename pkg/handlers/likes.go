package handlers

import (
	"fmt"
	"net/http"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	// "github.com/gomodule/redigo/redis"
)

type Like struct {
	ItemID string `json:"item_id"`
	UserID string `json:"user_id"`
}

func (h *Handlers) LikeItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lk := Like{}

		err := h.Json.ReadJson(w, r, &lk)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		str, err := h.likeItem(lk.ItemID, lk.UserID)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		err = h.Json.WriteJson(w, http.StatusOK, str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) UnlikeItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lk := Like{}

		err := h.Json.ReadJson(w, r, &lk)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		str, err := h.unlikeItem(lk.ItemID, lk.UserID)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		err = h.Json.WriteJson(w, http.StatusOK, str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) UserLikesItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lk := Like{}

		err := h.Json.ReadJson(w, r, &lk)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		str, err := h.userLikesItem(lk.ItemID, lk.UserID)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		err = h.Json.WriteJson(w, http.StatusOK, str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) likeItem(itemID, userID string) (int64, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	res, err := conn.Do("SADD", models.UserLikesKey(userID), itemID)
	if err != nil {
		return 0, err
	}

	fmt.Println("the res likeItem: ", res)

	// only if the item has been liked in the set of userid's likes
	if res == int64(1) {
		// increment by one the like of the item(as there is no other way to compte the total like/item)
		resInc, err := conn.Do("HINCRBYFLOAT", models.ItemsKey(itemID), "price", 1)
		if err != nil {
			return 0, err
		}

		fmt.Println("see the resInc: ", resInc)
	}

	return res.(int64), nil
}

func (h *Handlers) unlikeItem(itemID, userID string) (int64, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	// remove a key from a set
	res, err := conn.Do("SREM", models.UserLikesKey(userID), itemID)
	if err != nil {
		return 0, err
	}

	fmt.Println("the res unlikeItem: ", res)

	return res.(int64), nil

}

func (h *Handlers) userLikesItem(itemID, userID string) (int64, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	res, err := conn.Do("SISMEMBER", models.UserLikesKey(userID), itemID)
	if err != nil {
		return 0, err
	}

	fmt.Println("the res likeItem: ", res)

	return res.(int64), nil
}
