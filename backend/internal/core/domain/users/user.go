package users

import "time"

type User struct {
	ID             int64
	Username       string
	DisplayName    string
	Email          string
	HashedPassword string
	Avatar         []byte
}

type FriendRelationship struct {
	Requester   User
	Receiver    User
	RequestedAt time.Time
}
