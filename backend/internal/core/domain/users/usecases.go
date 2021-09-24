package users

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UseCases struct {
	repository Repository
}

const hashingCost = 10

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

func (u *UseCases) Register(displayName string, username string, password string) (int64, error) {
	_, err := u.repository.GetByUsername(username)
	if err == nil {
		return 0, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost)
	if err != nil {
		return 0, fmt.Errorf("unable to generate hash from password: %w", err)
	}

	return u.repository.Add(&User{
		DisplayName:    displayName,
		Username:       username,
		HashedPassword: string(hashedPassword),
	})
}

func (u *UseCases) Get(userID int64) (*User, error) {
	return u.repository.Get(userID)
}

func (u *UseCases) GetUsersByIds(userIDs []int64) (map[int64]User, error) {
	return u.repository.GetUsersByIds(userIDs)
}

func (u *UseCases) SearchByUsername(searchQuery string) ([]User, error) {
	return u.repository.SearchByUsername(searchQuery)
}

func (u *UseCases) GetFriends(userID int64) ([]*User, error) {
	return u.repository.GetFriends(userID)
}

func (u *UseCases) GetFriendRequests(userID int64) ([]*FriendRequest, error) {
	return u.repository.GetFriendRequests(userID)
}

func (u *UseCases) AddFriendRequest(requesterID int64, receiverID int64) (AddFriendRequestStatus, error) {
	// Try to accept any incoming friend requests first.
	err := u.AcceptFriendRequest(receiverID, requesterID)
	if err == nil {
		return AddedAsFriend, nil
	}

	// If there are no incoming friend requests, send a friend request.
	err = u.repository.AddFriendRequest(requesterID, receiverID)
	if err != nil {
		return 0, err
	}

	return SentFriendRequest, nil
}

func (u *UseCases) AcceptFriendRequest(requesterID, receiverID int64) error {
	return u.repository.AcceptFriendRequest(requesterID, receiverID)
}

func (u *UseCases) DeclineFriendRequest(requesterID, receiverID int64) error {
	return u.repository.DeclineFriendRequest(requesterID, receiverID)
}

func (u *UseCases) UpdateAvatar(userID int64, newAvatar []byte) error {
	return u.repository.UpdateAvatar(userID, newAvatar)
}

func (u *UseCases) UpdateDisplayName(userID int64, newDisplayName string) error {
	return u.repository.UpdateDisplayName(userID, newDisplayName)
}
