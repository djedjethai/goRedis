package handlers

import (
	"context"
	"fmt"
	// "strconv"
	"time"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

func (h *Handlers) getItem(ctx context.Context, iid string) models.Item {
	conn := h.app.Pool.Get()
	defer conn.Close()

	k := models.ItemsKey(iid)
	i := models.ItemAttr{}

	values, err := redis.Values(conn.Do("HGETALL", k))
	if err != nil {
		fmt.Println("err from redis get item")
	}

	redis.ScanStruct(values, &i)

	it := deSerializeItem(i)

	it.ID = iid

	return it
}

func (h *Handlers) getItems(iids []string) models.Item {

	return models.Item{}
}

func (h *Handlers) createItem(it models.Item) (string, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	id := uuid.New().String()

	k := models.ItemsKey(id)

	// it.CreatedAt = time.Now().Unix()
	it.CreatedAt = time.Now()
	it.EndingAt = time.Now()

	nv := serializeItem(it)

	fmt.Println("see nv: ", nv)

	_, err := conn.Do("HSET", redis.Args{}.Add(k).AddFlat(nv)...)
	if err != nil {
		return "", err
	}

	return id, nil
}

func unixTimestampsToTimeTime(t int64) time.Time {
	unixTimeUTC := time.Unix(t, 0)
	// unitTimeInRFC3339 :=unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format
	return unixTimeUTC

	// not the best even it get most like in so
	// i, _ := strconv.ParseInt(strconv.Itoa(int(t)), 10, 64)
	// tm := time.Unix(i, 0)
	// return tm
}

func deSerializeItem(i models.ItemAttr) models.Item {
	// convert unix timestamps to time.Time

	// Here i do not change time to string, let see ??
	return models.Item{
		ImageURL:         i.ImageURL,
		Description:      i.Description,
		Duration:         i.Duration,
		CreatedAt:        unixTimestampsToTimeTime(i.CreatedAt),
		EndingAt:         unixTimestampsToTimeTime(i.EndingAt),
		OwnerID:          i.OwnerID,
		HighestBidUserID: i.HighestBidUserID,
		Price:            i.Price,
		Views:            i.Views,
		Likes:            i.Likes,
		Bids:             i.Bids,
	}
}

func serializeItem(i models.Item) models.ItemAttr {
	// Here i do not change time to string, let see ??
	return models.ItemAttr{
		ImageURL:         i.ImageURL,
		Description:      i.Description,
		Duration:         i.Duration,
		CreatedAt:        i.CreatedAt.Unix(),
		EndingAt:         i.EndingAt.Unix(),
		OwnerID:          i.OwnerID,
		HighestBidUserID: i.HighestBidUserID,
		Price:            i.Price,
		Views:            i.Views,
		Likes:            i.Likes,
		Bids:             i.Bids,
	}
}
