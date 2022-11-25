package handlers

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/gomodule/redigo/redis"
	// "github.com/google/uuid"
)

func (h *Handlers) createUser(ctx context.Context, uInfo models.UserCredentials, k string) (string, error) {
	conn := h.app.Pool.Get()
	defer conn.Close()

	// make sure the username is not already in the set of username(a set can not have twice the same element)
	// tmp is an int64, if 1 means the element is present in set
	// first arg is the key, the second is the value
	tmp, err := conn.Do("SISMEMBER", models.UsernamesUniqueKey(), uInfo.Username)
	if err != nil {
		return "", err
	}

	// case the username is already in the set
	if tmp == int64(1) {
		return "", errors.New("Username already in use")
	}

	// we may have more info in userCredential, doing like so we make sure we only set these
	k = models.UserIDKey(k)

	_, err = conn.Do("HSET", redis.Args{}.Add(k).AddFlat(serializeUser(uInfo))...)
	// _, err = conn.Do("EXPIRE", k, expInSeconds)
	if err != nil {
		return "", err
	}

	// add the username to the set which track the username list
	// res is an int64 if 1 means it has been added
	res, err := conn.Do("SADD", models.UsernamesUniqueKey(), uInfo.Username)
	if err != nil {
		return "", err
	}

	fmt.Println("see 3: ", reflect.TypeOf(res))

	// // TODO .... IMPLEMENT A SORTED SET .....
	// the userID(k) is hexadecimal format, so change it to decimal first
	// ad redis sorted set take int for value
	// decimalID := 1234
	// fmt.Println("userName: ", uInfo.Username, " - ", reflect.TypeOf(uInfo.Username))

	// // add the username with the id(k is the id, decimal format which should be int64)
	// res, err = conn.Do("ZADD", models.UserNamesKey(), decimalID, uInfo.Username)
	// if err != nil {
	// 	fmt.Println("wtf.... ", reflect.TypeOf(decimalID))
	// 	return "", err
	// }

	fmt.Println("oookkk")

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

// func (h *Handlers) getUserByUsername(ctx context.Context, uname string) models.User {
// 	k := models.UserIDKey(uid)
// 	u := models.User{}
// 	conn := h.app.Pool.Get()
// 	defer conn.Close()
//
// 	values, err := redis.Values(conn.Do("HGETALL", k))
// 	if err != nil {
// 		// w.Write([]byte(err.Error()))
// 		return u
// 	}
//
// 	redis.ScanStruct(values, &u)
//
// 	return deserializeUser(uid, u)
// }

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
