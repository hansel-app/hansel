package repositories

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/hansel-app/hansel/internal/core/domain/users"
)

type userRepository struct {
	db sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: *db,
	}
}

func (r *userRepository) Get(id int64) (*users.User, error) {
	sql, _, _ := qb.From("users").Where(goqu.C("id").Eq(id)).ToSQL()

	var user users.User
	err := r.db.Get(&user, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get user with id %d: %w", id, err)
	}

	return &user, nil
}

func (r *userRepository) Add(user users.User) error {
	sql, _, _ := qb.Insert("users").Rows(user).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}

	return nil
}

func (r *userRepository) Update(id int64, user users.User) error {
	sql, _, _ := qb.Update("users").Where(goqu.C("id").Eq(id)).Set(user).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("unable to update user with id %d: %w", id, err)
	}

	return nil
}

func (r *userRepository) Remove(id int64) error {
	sql, _, _ := qb.From("users").Where(goqu.C("id").Eq(id)).Delete().ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("unable to delete user with id %d: %w", id, err)
	}

	return nil
}
