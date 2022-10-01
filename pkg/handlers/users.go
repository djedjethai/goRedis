package handlers

import (
	"context"
	// "fmt"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/gomodule/redigo/redis"
	// "github.com/google/uuid"
)

func (h *Handlers) createUser(ctx context.Context, uInfo models.UserCredentials, k string) (string, error) {
	// expInSeconds := 1000

	conn := h.app.Pool.Get()
	defer conn.Close()
	// id := uuid.New()
	// fmt.Println(id.String())

	// we may have more info in userCredential, doing like so we make sure we only set these
	// k := models.UserIDKey(id.String())

	_, err := conn.Do("HSET", redis.Args{}.Add(k).AddFlat(serializeUser(uInfo))...)
	// _, err = conn.Do("EXPIRE", k, expInSeconds)
	if err != nil {
		return "", err
	}
	return k, nil
}

func (h *Handlers) getUserByID(ctx context.Context, uid string) models.User {
	k := models.UserIDKey(uid)
	u := models.User{}
	conn := h.app.Pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", k))
	if err != nil {
		// w.Write([]byte(err.Error()))
		return u
	}

	redis.ScanStruct(values, &u)

	return deserializeUser(uid, u)
}

// make sure we use the right format, User may have more field than necessary
func serializeUser(uInfo models.UserCredentials) models.User {
	return models.User{
		Username: uInfo.Username,
		Password: uInfo.Password,
	}

}

// back from redis add the id back to the User
func deserializeUser(id string, user models.User) models.User {
	return models.User{
		ID:       id,
		Username: user.Username,
		Password: user.Password,
	}
}
