package repositories

import "github.com/hansel-app/hansel/internal/core/domain/gems"

type GemRepository struct{}

// TODO: Add an actual database.

func (r *GemRepository) Get(id int64) (*gems.Gem, error) {
	return &gems.Gem{
		ID:      id,
		Message: "Hello World!",
	}, nil
}

func (r *GemRepository) GetAllForUser(userId int64) ([]gems.Gem, error) {
	return []gems.Gem{}, nil
}

func (r *GemRepository) Add(gem *gems.Gem) (*gems.Gem, error) {
	return gem, nil
}

func (r *GemRepository) Update(gem *gems.Gem) (*gems.Gem, error) {
	return gem, nil
}

func (r *GemRepository) Remove(id int64) error {
	return nil
}
