package repositories

import (
	"fmt"
	"time"

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

func (r *userRepository) GetByUsername(username string) (*users.User, error) {
	sql, _, _ := qb.From("users").Where(goqu.C("username").Eq(username)).ToSQL()

	var user users.User
	err := r.db.Get(&user, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get user with username '%s': %w", username, err)
	}

	return &user, nil
}

func (r *userRepository) Add(user *users.User) error {
	sql, _, _ := qb.Insert("users").Rows(user).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}

	return nil
}

func (r *userRepository) Update(id int64, user *users.User) error {
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

func (r *userRepository) GetFriends(id int64) ([]*users.User, error) {
	firstSubquery := qb.
		Select(goqu.I("users.*")).
		From("users").
		LeftOuterJoin(goqu.T("friends"), goqu.On(goqu.Ex{
			"users.id": goqu.I("friends.first_user_id"),
		})).
		Where(goqu.I("friends.second_user_id").Eq(id))

	secondSubquery := qb.
		Select(goqu.I("users.*")).
		From("users").
		LeftOuterJoin(goqu.T("friends"), goqu.On(goqu.Ex{
			"users.id": goqu.I("friends.second_user_id"),
		})).
		Where(goqu.I("friends.first_user_id").Eq(id))

	sql, _, _ := firstSubquery.Union(secondSubquery).ToSQL()

	var friends []*users.User
	err := r.db.Select(&friends, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get friends for user with id %d: %w", id, err)
	}

	return friends, nil
}

func (r *userRepository) GetFriendRequests(id int64) ([]*users.FriendRelationship, error) {
	// TODO: Make this SQL statement correct
	// make status an enum somewhere + join with Users
	// + return datetime
	sql, _, _ := qb.From("friends").Where(goqu.And(
		goqu.C("receiver_id").Eq(id),
		goqu.C("status").Eq("PENDING"),
	)).ToSQL()

	var friendRequests []*users.FriendRelationship
	err := r.db.Select(&friendRequests, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get friend requests for user with id %d: %w", id, err)
	}

	return friendRequests, nil
}

func (r *userRepository) AddFriendRequest(requesterID int64, receiverID int64) error {
	sql, _, _ := qb.Insert("friends").Rows(goqu.Record{
		"requester_id": requesterID,
		"receiver_id":  receiverID,
		"requested_at": time.Now(),
	}).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf(
			"unable to add friend request for requester with id %d and receiver with id %d: %w",
			requesterID, receiverID, err,
		)
	}

	return nil
}

func (r *userRepository) AcceptFriendRequest(requesterID int64, receiverID int64) error {
	sql, _, _ := goqu.From("friends").Where(goqu.And(
		goqu.C("requester_id").Eq(requesterID),
		goqu.C("receiver_id").Eq(receiverID),
		goqu.C("status").Eq("PENDING"),
	)).Update().Set(goqu.Record{
		"status":        "FRIENDS",
		"friends_since": time.Now(),
	}).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf(
			"unable to decline friend request for requester with id %d and receiver with id %d: %w",
			requesterID, receiverID, err,
		)
	}

	return nil
}

func (r *userRepository) DeclineFriendRequest(requesterID int64, receiverID int64) error {
	sql, _, _ := goqu.From("friends").Where(goqu.And(
		goqu.C("requester_id").Eq(requesterID),
		goqu.C("receiver_id").Eq(receiverID),
		goqu.C("status").Eq("PENDING"),
	)).Delete().ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf(
			"unable to decline friend request for requester with id %d and receiver with id %d: %w",
			requesterID, receiverID, err,
		)
	}

	return nil
}
