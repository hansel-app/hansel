package users

import "time"

type User struct {
	ID             int64  `db:"id"`
	Username       string `db:"username"`
	DisplayName    string `db:"display_name"`
	HashedPassword string `db:"hashed_password"`
	Avatar         []byte `db:"avatar"`
}

type FriendRequest struct {
	Requester User
	CreatedAt time.Time `db:"created_at"`
}
