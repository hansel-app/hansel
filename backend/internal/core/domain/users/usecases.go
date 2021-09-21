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

func (u *UseCases) Get(userID int64) (*User, error) {
	return u.repository.Get(userID)
}

func (u *UseCases) GetFriends(userID int64) ([]*User, error) {
	return u.repository.GetFriends(userID)
}

func (u *UseCases) GetFriendRequests(userID int64) ([]*FriendRelationship, error) {
	return u.repository.GetFriendRequests(userID)
}

func (u *UseCases) AddFriendRequest(requesterID int64, receiverID int64) error {
	return u.repository.AddFriendRequest(requesterID, receiverID)
}

func (u *UseCases) AcceptFriendRequest(requesterID, receiverID int64) error {
	return u.repository.AcceptFriendRequest(requesterID, receiverID)
}

func (u *UseCases) DeclineFriendRequest(requesterID, receiverID int64) error {
	return u.repository.DeclineFriendRequest(requesterID, receiverID)
}
