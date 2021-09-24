package users

import "time"

type AddFriendRequestStatus int64

const (
	SentFriendRequest AddFriendRequestStatus = 0
	AddedAsFriend     AddFriendRequestStatus = 1
)

type User struct {
	ID             int64  `db:"id" goqu:"skipinsert"`
	Username       string `db:"username"`
	DisplayName    string `db:"display_name"`
	HashedPassword string `db:"hashed_password"`
	Avatar         []byte `db:"avatar"`
}

type FriendRequest struct {
	Requester User
	CreatedAt time.Time `db:"created_at"`
}
