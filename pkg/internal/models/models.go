package models

import (
	"time"
)

type User struct {
	ID       string `redis:"id"`
	Username string `redis:"username"`
	Password string `redis:"password"`
}

type UserCredentials struct {
	Username string `redis:"username"`
	Password string `redis:"password"`
}

type Item struct {
	ID               string    `redis:"id" json:"id"`
	ImageURL         string    `redis:"image_url" json:"image_url"`
	Description      string    `redis:"description" json:"description"`
	Duration         float64   `redis:"duration" json:"duration"`
	CreatedAt        time.Time `redis:"created_at" json:"created_at"`
	EndingAt         time.Time `redis:"ending_at" json:"ending_at"`
	OwnerID          string    `redis:"owner_id" json:"owner_id"`
	HighestBidUserID string    `redis:"highest_bid_user_id" json:"highest_bid_user_id"`
	Price            float64   `redis:"price" json:"price"`
	Views            int       `redis:"views" json:"views"`
	Likes            int       `redis:"likes" json:"likes"`
	Bids             int       `redis:"bids" json:"bids"`
}

type ItemAttr struct {
	ImageURL         string  `redis:"image_url" json:"image_url"`
	Description      string  `redis:"description" json:"description"`
	Duration         float64 `redis:"duration" json:"duration"`
	CreatedAt        int64   `redis:"created_at" json:"created_at"`
	EndingAt         int64   `redis:"ending_at" json:"ending_at"`
	OwnerID          string  `redis:"owner_id" json:"owner_id"`
	HighestBidUserID string  `redis:"highest_bid_user_id" json:"highest_bid_user_id"`
	Price            float64 `redis:"price" json:"price"`
	Views            int     `redis:"views" json:"views"`
	Likes            int     `redis:"likes" json:"likes"`
	Bids             int     `redis:"bids" json:"bids"`
}

// type Session struct {
// 	ID       string `redis:"id"`
// 	UserID   string `redis:"user_id"`
// 	Username string `redis:"username"`
// }
//
// type SessionCredentials struct {
// 	UserID   string `redis:"user_id"`
// 	Username string `redis:"username"`
// }
