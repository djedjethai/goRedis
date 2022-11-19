package main

import (
	// "fmt"

	// "github.com/djedjethai/goRedis/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	// "github.com/tsawler/gowatcher/pkg/config"
	"net/http"
)

func routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// login
	mux.Get("/", handler.Test())

	mux.Post("/signin/{id}", handler.Signin())
	mux.Get("/user/{id}", handler.GetUser())

	mux.Route("/api", func(mux chi.Router) {
		// mux.Use(auth)
		mux.Post("/createitem", handler.CreateItem())
		mux.Get("/getitem/{id}", handler.GetItem())
		mux.Post("/getitems", handler.GetItems())
		mux.Post("/likeitem", handler.LikeItem())
		mux.Post("/unlikeitem", handler.UnlikeItem())
		mux.Post("/userlikesitem", handler.UserLikesItem())
		mux.Post("/likeditems", handler.LikedItems())
		mux.Post("/commonlikeditems", handler.CommonLikedItems())
	})

	return mux
}
