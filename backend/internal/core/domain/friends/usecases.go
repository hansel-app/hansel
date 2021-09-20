package friends

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
func (u *UseCases) Add(friend1_id int64, friend2_id int64) error {
	return u.repository.Add(friend1_id, friend2_id)
}

func (u *UseCases) Get(id int64) ([]users.User, error) {
	return u.repository.Get(id)
}
