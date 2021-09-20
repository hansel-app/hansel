package repositories

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/hansel-app/hansel/internal/core/domain/users"
)

type friendRepository struct {
	db sqlx.DB
}

func NewFriendRepository(db *sqlx.DB) *friendRepository {
	return &friendRepository{
		db: *db,
	}
}

func (r *friendRepository) Get(id int64) ([]*users.User, error) {
	// TODO: Make this SQL statement correct.
	// need to join with Users ?
	sql, _, _ := qb.From("friends").Where(goqu.Or(
		goqu.C("friend1_id").Eq(id),
		goqu.C("friend2_id").Eq(id),
	)).ToSQL()

	var friends []*users.User
	err := r.db.Select(&friends, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get friends of user with id %d: %w", id, err)
	}

	return friends, nil
}
