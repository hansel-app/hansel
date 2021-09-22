package users

import "time"

type User struct {
	ID             int64  `db:"id"`
	Username       string `db:"username"`
	DisplayName    string `db:"display_name"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Avatar         []byte `db:"avatar"`
}

type FriendRelationshipStatus int64

const (
	Pending  FriendRelationshipStatus = 0
	Accepted FriendRelationshipStatus = 1
)

type FriendRelationship struct {
	Requester   User
	Receiver    User
	RequestedAt time.Time
	Status      FriendRelationshipStatus
}
