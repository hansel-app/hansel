package friend_requests

import (
	"github.com/hansel-app/hansel/internal/core/domain/users"
)


type Repository interface {
	Get(id int64) ([]users.User, error)
	Add(requester_id int64, receiver_id int64) error
	Remove(requester_id int64, receiver_id int64) error
}
