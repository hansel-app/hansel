package repositories

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/hansel-app/hansel/internal/core/domain/gems"
	"github.com/jmoiron/sqlx"
)

type gemRepository struct {
	db sqlx.DB
}

func NewGemRepository(db *sqlx.DB) *gemRepository {
	return &gemRepository{
		db: *db,
	}
}

var qb = goqu.Dialect("psql")

func (r *gemRepository) Get(id int64) (*gems.Gem, error) {
	sql, _, _ := qb.From("gems").Where(goqu.C("id").Eq(id)).ToSQL()

	var gem gems.Gem
	err := r.db.Get(&gem, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get gem with id %d: %w", id, err)
	}

	return &gem, nil
}

func (r *gemRepository) GetAllForUser(userId int64) ([]gems.Gem, error) {
	// TODO: Add real logic.
	return []gems.Gem{}, nil
}

func (r *gemRepository) Add(message string) (*gems.Gem, error) {
	sql, _, _ := qb.Insert("gems").Rows(goqu.Record{
		"message": message,
	}).Returning(goqu.T("gems").All()).ToSQL()

	var gem gems.Gem
	err := r.db.Get(&gem, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to add gem: %w", err)
	}

	return &gem, nil
}

func (r *gemRepository) Update(gem *gems.Gem) (*gems.Gem, error) {
	// TODO: Add real logic.
	return gem, nil
}

func (r *gemRepository) Remove(id int64) error {
	// TODO: Add real logic.
	return nil
}
