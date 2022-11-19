package handlers

import (
	"fmt"
	"net/http"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/gomodule/redigo/redis"
)

type Like struct {
	ItemID string `json:"item_id"`
	UserID string `json:"user_id"`
}

type CommonLike struct {
	UserID1 string `json:"user_id1"`
	UserID2 string `json:"user_id2"`
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

func (h *Handlers) LikedItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lk := Like{}

		err := h.Json.ReadJson(w, r, &lk)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		arrItems, err := h.likedItems(lk.UserID)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		err = h.Json.WriteJson(w, http.StatusOK, arrItems)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) CommonLikedItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cl := CommonLike{}

		err := h.Json.ReadJson(w, r, &cl)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		its, err := h.commonLikedItems(cl.UserID1, cl.UserID2)
		if err != nil {
			h.Json.BadRequest(w, r, err)
		}

		err = h.Json.WriteJson(w, http.StatusOK, its)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *Handlers) commonLikedItems(userID1, userID2 string) ([]models.Item, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	// return elements which are common to each set(the liked items' id) of user1 and the set of user2
	values, err := redis.Values(conn.Do("SINTER", models.UserLikesKey(userID1), models.UserLikesKey(userID2)))
	if err != nil {
		return []models.Item{}, err
	}

	res := []string{}
	for _, r := range values {
		fmt.Println(string(r.([]byte)))
		res = append(res, string(r.([]byte)))
	}

	return h.getItems(res)
}

func (h *Handlers) likedItems(userID string) ([]models.Item, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	// get all itemID the user likes from the set
	values, err := redis.Values(conn.Do("SMEMBERS", models.UserLikesKey(userID)))
	if err != nil {
		return []models.Item{}, err
	}

	res := []string{}
	for _, r := range values {
		res = append(res, string(r.([]byte)))
	}

	return h.getItems(res)

}

func (h *Handlers) likeItem(itemID, userID string) (int64, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	fmt.Println("before the res likeItem: ", userID, " - ", itemID)

	// add 1 itemID to the set of userID
	res, err := conn.Do("SADD", models.UserLikesKey(userID), itemID)
	if err != nil {
		return 0, err
	}

	// only if the item has been liked in the set of userid's likes
	if res == int64(1) {
		// increment by one the like of the item(as there is no other way to compte the total like/item)
		_, err := conn.Do("HINCRBYFLOAT", models.ItemsKey(itemID), "likes", 1)
		if err != nil {
			return 0, err
		}

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
	// then remove the like from the ItemAttr
	if res == int64(1) {
		resInc, err := conn.Do("HINCRBYFLOAT", models.ItemsKey(itemID), "likes", -1)
		if err != nil {
			return 0, err
		}

		fmt.Println("see the resDec: ", resInc)

	}

	return res.(int64), nil

}

func (h *Handlers) userLikesItem(itemID, userID string) (int64, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	res, err := conn.Do("SISMEMBER", models.UserLikesKey(userID), itemID)
	if err != nil {
		return 0, err
	}

	fmt.Println("the res userLikesItem: ", res)

	return res.(int64), nil
}
