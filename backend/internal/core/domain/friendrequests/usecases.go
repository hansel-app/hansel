package friend_requests

import (
	"github.com/hansel-app/hansel/internal/core/domain/users"
)

type UseCases struct {
	repository Repository
}

func NewUseCases(repository Repository) *UseCases {
	return &UseCases{
		repository: repository,
	}
}
func (u *UseCases) Add(requester_id int64, receiver_id int64) error {
	return u.repository.Add(requester_id, receiver_id)
}

func (u *UseCases) Get(id int64) ([]users.User, error) {
	return u.repository.Get(id)
}
