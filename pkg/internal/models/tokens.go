package models

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	// "log"
	"time"
)

const (
	ScopeAuthentication = "authentication"
)

// Token is the type of authentications token
type Token struct {
	PlainText string    `json:"token"`
	Email     string    `json:"-"`
	Hash      []byte    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

// GenerateToken generates a token that last for ttl, and returns it
func (t *Token) GenerateToken(email string, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		Email:  email,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256(([]byte(token.PlainText)))
	token.Hash = hash[:]
	return token, nil
}

func (t *Token) InsertToken(tok *Token, u User) error {
	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// TODO ......
	// delete existing token(if token exist, otherwise nothing will be delete)
	// stmt := `delete from tokens where user_id = ?`
	// _, err := m.DB.ExecContext(ctx, stmt, u.ID)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println("the token expire: ", t.Expiry)

	// stmt = `insert into tokens (user_id, name, email, token_hash, created_at, updated_at, expiry)
	// 	values (?, ?, ?, ?, ?, ?, ?)`

	// _, err = m.DB.ExecContext(ctx, stmt,
	// 	u.ID,
	// 	u.LastName,
	// 	u.Email,
	// 	t.Hash,
	// 	time.Now(),
	// 	time.Now(),
	// 	t.Expiry,
	// )
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (t *Token) GetUserFromToken(token string) (*User, error) {
	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tokenHash := sha256.Sum256([]byte(token))
	var user User

	fmt.Println("the tok: ", tokenHash)

	// TODO .....
	// u will be an alias to the users table
	// query := `
	// 	select
	// 		u.id, u.first_name, u.last_name, u.email
	// 	from
	// 		users u
	// 		inner join tokens t on(u.id = t.user_id)
	// 	where
	// 		t.token_hash = ?
	// 		and t.expiry > ?
	// `

	// // QueryRowContext() as we know we will not get more than one token
	// // tokenHash is a slice and here i need an array
	// // the way to convert it is: tokenHash[:]
	// err := m.DB.QueryRowContext(ctx, query, tokenHash[:], time.Now()).Scan(
	// 	&user.ID,
	// 	&user.FirstName,
	// 	&user.LastName,
	// 	&user.Email,
	// )
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	return &user, nil
}
