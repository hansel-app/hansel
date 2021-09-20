package repositories

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type friendRepository struct {
	db sqlx.DB
}

func NewFriendRepository(db *sqlx.DB) *friendRepository {
	return &friendRepository{
		db: *db,
	}
}

func (r *friendRepository) Get(id int64) ([]int64, error) {
	sql, _, _ := qb.From("friends").Where(goqu.Or(
		goqu.C("friend1_id").Eq(id),
		goqu.C("friend2_id").Eq(id),
	)).ToSQL()

	var friends []int64
	err := r.db.Select(&friends, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get friends of user with id %d: %w", id, err)
	}

	return friends, nil
}
