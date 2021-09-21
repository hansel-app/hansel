package users

type Repository interface {
	Get(id int64) (*User, error)
	GetByUsername(username string) (*User, error)
	Add(user *User) error
	Update(id int64, user *User) error
	Remove(id int64) error
	GetFriends(id int64) ([]*User, error)
	GetFriendRequests(id int64) ([]*User, error)
	SendFriendRequest() (int64, error)
	AcceptFriendRequest() (int64, error)
	DeclineFriendRequest() (int64, error)
}
