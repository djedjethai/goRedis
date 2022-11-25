package handlers

import (
	"context"
	// "encoding/json"
	"fmt"
	"strings"

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

	it := deSerializeItem(iid, i)

	return it
}

// a1388cd2-017f-4634-80bc-2c82823c38b4, 0c8b2ae4-fbfc-4a72-b0bf-403b9c195859
func (h *Handlers) getItems(iids []string) ([]models.Item, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	var items = []models.Item{}

	// Initialize Pipeline
	conn.Send("MULTI")

	// Send writes the command to the connection's output buffer
	var s string
	for _, v := range iids {
		s = models.ItemsKey(strings.TrimSpace(v))
		conn.Send("HGETALL", s)
	}

	// Execute the Pipeline
	pipe_prox, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return items, err
	}

	// Use the generic redis.Values to break down this outer structure first
	// then we can parse the inner slices.
	ita := models.ItemAttr{}
	for i, r := range pipe_prox {
		// s, _ := redis.Strings(r, nil)
		s, _ := redis.Values(r, nil)

		redis.ScanStruct(s, &ita)

		it := deSerializeItem(iids[i], ita)

		items = append(items, it)
	}

	return items, nil
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

	// TODO pipelining these 2 cmd
	// _, err := conn.Do("HSET", redis.Args{}.Add(k).AddFlat(nv)...)
	// if err != nil {
	// 	return "", err
	// }
	// // set the number of views in the sortedSet registering it
	// _, err = conn.Do("ZADD", models.ItemsByViewsKey(), 0, id)
	// if err != nil {
	// 	return "", err
	// }

	// the pipeline (6ff52659-8259-4871-8a74-91a64edf543c)
	conn.Send("MULTI")
	conn.Send("HSET", redis.Args{}.Add(k).AddFlat(nv)...)
	conn.Send("ZADD", models.ItemsByViewsKey(), 0, id)
	pipe_prox, err := conn.Do("EXEC")
	if err != nil {
		return "", err
	}

	fmt.Println("seeee the pipe_prox: ", pipe_prox)

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

func deSerializeItem(id string, i models.ItemAttr) models.Item {
	// convert unix timestamps to time.Time

	// Here i do not change time to string, let see ??
	return models.Item{
		ID:               id,
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
