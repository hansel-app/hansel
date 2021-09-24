package repositories

import (
	"fmt"
	"time"

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

func (r *gemRepository) GetPendingCollectionByUser(userId int64) ([]gems.Gem, error) {
	sql, _, _ := qb.From("gems").Where(
		goqu.C("receiver_id").Eq(userId),
		goqu.C("received_at").IsNull(),
	).ToSQL()

	var gems []gems.Gem
	err := r.db.Select(&gems, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get gems pending collection for user with id %d: %w", userId, err)
	}

	return gems, nil
}

func (r *gemRepository) PickUpGem(id int64) error {
	sql, _, _ := qb.Update("gems").Where(goqu.C("id").Eq(id)).Set(
		goqu.Record{"received_at": time.Now()},
	).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("unable to pick up gem with id %d: %w", id, err)
	}

	return nil
}

func (r *gemRepository) Add(gem *gems.Gem) (int64, error) {
	// Necessary to use prepared statements here due to a bug in how goqu parses
	// bytes (used for the gem attachments)
	// More info here: https://github.com/doug-martin/goqu/issues/254
	sql, args, _ := qb.Insert("gems").Prepared(true).Rows(gem).Returning("id").ToSQL()

	var gemId int64
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}

	err = stmt.QueryRow(args...).Scan(&gemId)
	if err != nil {
		return 0, fmt.Errorf("unable to add gem: %w", err)
	}

	return gemId, nil
}

func (r *gemRepository) GetGemLogs(userId int64) (map[int64]([]gems.Gem), error) {
	gemsCreatedbySelf := qb.From("gems").Where(goqu.C("creator_id").Eq(userId))
	gemsSentToSelf := qb.From("gems").Where(goqu.C("receiver_id").Eq(userId))

	sql, _, _ := gemsCreatedbySelf.Union(gemsSentToSelf).ToSQL()

	rows, err := r.db.Queryx(sql)
	if err != nil {
		return nil, fmt.Errorf("unable to add gem: %w", err)
	}

	gemMap := make(map[int64][]gems.Gem)

	for rows.Next() {
		var gem gems.Gem
		err = rows.StructScan(&gem)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan row into gem: %s", err)
		}

		if gem.CreatorId != userId && gem.ReceiverId != userId {
			panic("Retrieved gem must be associated to the user")
		}

		if gem.CreatorId != userId {
			friendGems, ok := gemMap[gem.CreatorId]
			if !ok {
				gemMap[gem.CreatorId] = make([]gems.Gem, 0)
			}
			gemMap[gem.CreatorId] = append(friendGems, gem)
		}

		if gem.ReceiverId != userId {
			friendGems, ok := gemMap[gem.ReceiverId]
			if !ok {
				gemMap[gem.ReceiverId] = make([]gems.Gem, 0)
			}
			gemMap[gem.ReceiverId] = append(friendGems, gem)
		}
	}

	return gemMap, nil
}
