package models

type User struct {
	ID       string `redis:"id"`
	Username string `redis:"username"`
	Password string `redis:"password"`
}

type UserCredentials struct {
	Username string `redis:"username"`
	Password string `redis:"password"`
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
