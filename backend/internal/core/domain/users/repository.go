package users

type Repository interface {
	Get(id int64) (*User, error)
	GetUsersByIds(ids []int64) (map[int64]User, error)
	GetByUsername(username string) (*User, error)
	SearchByUsername(searchQuery string) ([]User, error)
	Add(user *User) (int64, error)
	Update(id int64, user *User) error
	Remove(id int64) error
	GetFriends(id int64) ([]*User, error)
	GetPendingFriends(id int64) ([]*User, error)
	GetFriendRequests(id int64) ([]*FriendRequest, error)
	AddFriendRequest(requesterID int64, receiverID int64) error
	AcceptFriendRequest(requesterID int64, receiverID int64) error
	DeclineFriendRequest(requesterID int64, receiverID int64) error
	UpdateAvatar(id int64, avatar []byte) error
	UpdateDisplayName(id int64, displayName string) error
}
