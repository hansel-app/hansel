package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UseCases struct {
	repository Repository
}

func NewUseCases(repository Repository) *UseCases {
	return &UseCases{
		repository: repository,
	}
}

// AuthenticatePassword checks that the username/password combination exists for a user and returns
// the ID of the user if so.
func (u *UseCases) AuthenticatePassword(username string, password string) (int64, error) {
	user, err := u.repository.GetByUsername(username)
	if err != nil {
		return 0, fmt.Errorf("unable to find user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("incorrect password: %w", err)
	}

	return user.ID, nil
}

func (u *UseCases) GetFriends(user_id int64) ([]*User, error) {
	return u.repository.GetFriends(user_id)
}

func (u *UseCases) GetFriendRequests(user_id int64) ([]*User, error) {
	return u.repository.GetFriendRequests(user_id)
}

func (u *UseCases) AddFriendRequest(requester_id int64, receiver_id int64) (error) {
	return u.repository.AddFriendRequest(requester_id, receiver_id)
}


func (u *UseCases) AcceptFriendRequest(sender_id int64) (error) {
	return u.repository.AcceptFriendRequest(sender_id)
}

func (u *UseCases) DeclineFriendRequest(sender_id int64) (error) {
	return u.repository.DeclineFriendRequest(sender_id)
}
