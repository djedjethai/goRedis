package handlers

// import (
// 	"errors"
// 	"net/http"
// 	// "fmt"
//
// 	"github.com/djedjethai/goRedis/pkg/internal"
// 	"github.com/djedjethai/goRedis/pkg/models"
// 	"github.com/gomodule/redigo/redis"
// 	"github.com/google/uuid"
// )
//
// func Signup() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// get username and password from post body
//
// 		// create a session id
// 		id := uuid.New()
//
// 		// save the session
// 		sess
// 	}
// }
//
// func getSession(sid string) (models.Session, error) {
// 	k := internal.SessionKey(sid)
// 	s := models.Session{}
//
// 	values, err := redis.Values(conn.Do("HGETALL", k))
// 	if err != nil {
// 		return s, errors.New("No session")
// 	}
//
// 	redis.ScanStruct(values, &s)
//
// 	return deserializeSession(sid, s), nil
// }
//
// func createSession(s models.Session) string {
// 	k := internal.SessionKey(s.ID)
//
// 	_, err := conn.Do("HSET", redis.Args{}.Add(k).AddFlat(serializeSession(s))...)
// 	// _, err = conn.Do("EXPIRE", k, expInSeconds)
// 	if err != nil {
// 		return ""
// 	}
//
// 	return s.ID
// }
//
// func serializeSession(s models.Session) models.SessionCredentials {
// 	return models.SessionCredentials{
// 		UserID:   s.UserID,
// 		Username: s.Username,
// 	}
// }
//
// func deserializeSession(sid string, s models.Session) models.Session {
// 	return models.Session{
// 		ID:       sid,
// 		UserID:   s.UserID,
// 		Username: s.Username,
// 	}
// }
