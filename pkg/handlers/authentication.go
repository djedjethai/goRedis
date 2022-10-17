package handlers

import (
	"context"
	// "errors"
	"fmt"
	"net/http"
	// "strconv"
	// "strings"
	"time"

	"github.com/djedjethai/goRedis/pkg/internal/models"
	"github.com/go-chi/chi/v5"
	// "github.com/gomodule/redigo/redis"
	// "golang.org/x/crypto/bcrypt"
)

func (h *Handlers) Signin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")

		var userInput struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		err := h.Json.ReadJson(w, r, &userInput)
		if err != nil {
			h.Json.BadRequest(w, r, err)
			return
		}

		// generate the token
		token, err := h.token.GenerateToken(userInput.Password, 24*time.Hour, models.ScopeAuthentication)
		if err != nil {
			h.Json.BadRequest(w, r, err)
			return
		}

		user := models.UserCredentials{}
		user.Username = userInput.Username
		user.Password = userInput.Password

		// TODO should save the user and the token in db
		// then when connect we check the token, and set the userID as key in redis
		// BUT HERE JUST USE USER as key

		// create the user and save in redis, the key is the token
		ctx := context.Background()
		//
		t, err := h.createUser(ctx, user, id)
		if err != nil {
			fmt.Println("Err create user: ", err)
		}
		fmt.Println("Token: ", t)

		payload := struct {
			Error   bool          `json:"error"`
			Message string        `json:"message"`
			Token   *models.Token `json:"authentication_token"`
		}{
			Error:   false,
			Message: fmt.Sprintf("token for %s created", user.Username),
			Token:   token,
		}

		err = h.Json.WriteJson(w, http.StatusOK, payload)
		if err != nil {
			h.Json.BadRequest(w, r, err)
			return
		}
	}
}

// TODO get the users data back....(from redis)
// func (h *Handlers) AuthenticateToken(r *http.Request) (*models.User, error) {
// 	authorizationHeader := r.Header.Get("Authorization")
// 	if authorizationHeader == "" {
// 		return nil, errors.New("No authorization header received")
// 	}
//
// 	headerParts := strings.Split(authorizationHeader, " ")
// 	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
// 		return nil, errors.New("No authorization header received")
// 	}
//
// 	token := headerParts[1]
// 	if len(token) != 26 {
// 		return nil, errors.New("Authorization token wrong size")
// 	}
//
// 	// TODO create the user + add token ?
// 	// get the user matching the token back from redis
//
// 	user := &models.User{}
//
// 	// get the user from the token's table
// 	// user, err := app.DB.GetUserFromToken(token)
// 	// if err != nil {
// 	// 	return nil, errors.New("No matching user found")
// 	// }
//
// 	// verif token validity
// 	return user, nil
// }

// func (h *Handlers) CheckAuthentication(w http.ResponseWriter, r *http.Request) {
// 	// validate the etoken, and get associated user
// 	user, err := h.AuthenticateToken(r)
// 	if err != nil {
// 		h.Json.InvalidCredentials(w)
// 		return
// 	}
//
// 	// valid user
// 	var payload struct {
// 		Error   bool   `json:"error"`
// 		Message string `json:"message"`
// 	}
// 	payload.Error = false
// 	payload.Message = fmt.Sprintf("authenticated user: %s", user.Username)
//
// 	h.Json.WriteJson(w, http.StatusOK, payload)
// }
