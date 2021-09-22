package users

import "time"

type User struct {
	ID             int64  `db:"id"`
	Username       string `db:"username"`
	DisplayName    string
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Avatar         []byte
}

type FriendRelationship struct {
	Requester   User
	Receiver    User
	RequestedAt time.Time
}
