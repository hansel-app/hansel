package friends

import (
	"github.com/hansel-app/hansel/internal/core/domain/users"
)


type Repository interface {
	Get(id int64) ([]users.User, error)
	Add(friend1_id int64, friend2_id int64) error
	Remove(id int64) error
}
