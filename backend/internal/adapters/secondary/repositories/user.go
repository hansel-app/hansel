package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/hansel-app/hansel/internal/constants"
	"github.com/hansel-app/hansel/internal/core/domain/users"
	"github.com/jmoiron/sqlx"
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

func (r *userRepository) GetUsersByIds(ids []int64) (map[int64]users.User, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no user ids supplied")
	}

	sql, _, _ := qb.From("users").Where(goqu.I("id").In(ids)).ToSQL()

	var selectedUsers []users.User
	err := r.db.Select(&selectedUsers, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve users with ids '%v': %w", ids, err)
	}

	userMap := make(map[int64]users.User)
	for _, user := range selectedUsers {
		userMap[user.ID] = user
	}

	return userMap, nil
}

func (r *userRepository) GetByUsername(username string) (*users.User, error) {
	sql, _, _ := qb.From(goqu.T("users")).Where(goqu.C("username").Eq(username)).ToSQL()

	var user users.User
	err := r.db.Get(&user, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get user with username '%s': %w", username, err)
	}

	return &user, nil
}

func (r *userRepository) SearchByUsername(searchQuery string) ([]users.User, error) {
	formattedSearchQuery := fmt.Sprintf("%%%s%%", searchQuery)
	sql, _, _ := qb.From("users").Where(goqu.C("username").ILike(formattedSearchQuery)).ToSQL()

	var matchingUsers []users.User
	err := r.db.Select(&matchingUsers, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get users with matching username: %w", err)
	}

	return matchingUsers, nil
}

func (r *userRepository) Add(user *users.User) (int64, error) {
	errorMsg := "unable to insert user: %w"

	tx := r.db.MustBegin()

	sql, args, _ := qb.Insert(goqu.T("users")).Prepared(true).Rows(user).Returning("id").ToSQL()

	var userId int64
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}

	err = stmt.QueryRow(args...).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf(errorMsg, err)
	}

	// Newly registered users are automatically friends with Gretel
	var firstUserID, secondUserID int64
	if constants.GretelUserId < userId {
		firstUserID = constants.GretelUserId
		secondUserID = userId
	} else {
		firstUserID = userId
		secondUserID = constants.GretelUserId
	}

	sql, _, _ = qb.
		Insert(goqu.T("friends")).
		Rows(goqu.Record{
			"first_user_id":  firstUserID,
			"second_user_id": secondUserID,
		}).
		ToSQL()

	_, err = tx.Exec(sql)
	if err != nil {
		return 0, fmt.Errorf(errorMsg, err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf(errorMsg, err)
	}

	return userId, nil
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

func (r *userRepository) GetPendingFriends(id int64) ([]*users.User, error) {
	sql, _, _ := qb.
		Select(goqu.I("users.*")).
		From("users").
		LeftOuterJoin(goqu.T("friend_requests"), goqu.On(goqu.Ex{
			"users.id": goqu.I("friend_requests.receiver_id"),
		})).
		Where(goqu.I("friend_requests.requester_id").Eq(id)).
		ToSQL()

	var pendingFriends []*users.User
	err := r.db.Select(&pendingFriends, sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get pending friends for user with id %d: %w", id, err)
	}

	return pendingFriends, nil
}

func (r *userRepository) GetFriendRequests(id int64) ([]*users.FriendRequest, error) {
	sql, _, _ := qb.
		Select(goqu.I("users.*"), goqu.I("friend_requests.created_at")).
		From("friend_requests").
		LeftOuterJoin(goqu.T("users"), goqu.On(goqu.Ex{
			"friend_requests.requester_id": goqu.I("users.id"),
		})).
		Where(goqu.C("receiver_id").Eq(id)).
		ToSQL()

	rows, err := r.db.Queryx(sql)
	if err != nil {
		return nil, fmt.Errorf("unable to get friend requests for user with id %d: %w", id, err)
	}

	type friendRequestWithRequester struct {
		users.FriendRequest
		users.User
	}

	var friendRequests []*users.FriendRequest
	for rows.Next() {
		var result friendRequestWithRequester
		err = rows.StructScan(&result)
		if err != nil {
			break
		}

		friendRequest := users.FriendRequest{
			Requester: users.User{
				ID:             result.ID,
				Username:       result.Username,
				DisplayName:    result.DisplayName,
				HashedPassword: result.HashedPassword,
				Avatar:         result.Avatar,
			},
			CreatedAt: result.CreatedAt,
		}
		friendRequests = append(friendRequests, &friendRequest)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to get friend requests for user with id %d: %w", id, err)
	}

	return friendRequests, nil
}

func (r *userRepository) checkIfUsersAreFriends(firstUserID int64, secondUserID int64) (bool, error) {
	sql, _, _ := qb.
		Select(goqu.L("COUNT(*)")).
		From(goqu.T("friends")).
		Where(goqu.Or(
			goqu.And(
				goqu.C("first_user_id").Eq(firstUserID),
				goqu.C("second_user_id").Eq(secondUserID),
			),
			goqu.And(
				goqu.C("first_user_id").Eq(secondUserID),
				goqu.C("second_user_id").Eq(firstUserID),
			),
		)).
		ToSQL()

	var count int64
	err := r.db.Get(&count, sql)
	if err != nil {
		return false, fmt.Errorf("unable to check if users are friends: %w", err)
	}

	return count == 1, nil
}

func (r *userRepository) AddFriendRequest(requesterID int64, receiverID int64) error {
	usersAreFriends, err := r.checkIfUsersAreFriends(requesterID, receiverID)
	if err != nil {
		return err
	}
	if usersAreFriends {
		return fmt.Errorf("requester with id %d and receiver with id %d are already friends", requesterID, receiverID)
	}

	sql, _, _ := qb.Insert("friend_requests").Rows(goqu.Record{
		"requester_id": requesterID,
		"receiver_id":  receiverID,
		"created_at":   time.Now(),
	}).ToSQL()

	_, err = r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf(
			"unable to add friend request for requester with id %d and receiver with id %d: %w",
			requesterID, receiverID, err,
		)
	}

	return nil
}

func (r *userRepository) AcceptFriendRequest(requesterID int64, receiverID int64) error {
	errorMsg := fmt.Sprintf(
		"unable to accept friend request for requester with id %d and receiver with id %d: %%w",
		requesterID, receiverID,
	)

	tx := r.db.MustBegin()

	var firstUserID, secondUserID int64
	if requesterID < receiverID {
		firstUserID = requesterID
		secondUserID = receiverID
	} else {
		firstUserID = receiverID
		secondUserID = requesterID
	}
	sql, _, _ := qb.
		Insert(goqu.T("friends")).
		Rows(goqu.Record{
			"first_user_id":  firstUserID,
			"second_user_id": secondUserID,
		}).
		ToSQL()
	_, err := tx.Exec(sql)
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}

	sql, _, _ = qb.
		From(goqu.T("friend_requests")).
		Where(goqu.C("requester_id").Eq(requesterID), goqu.C("receiver_id").Eq(receiverID)).
		Delete().
		ToSQL()
	res, err := tx.Exec(sql)
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("unable to get rows affected")
	}
	if rowsAffected == 0 {
		return errors.New("unable to accept non-existent friend request")
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}

	return nil
}

func (r *userRepository) DeclineFriendRequest(requesterID int64, receiverID int64) error {
	sql, _, _ := qb.
		From(goqu.T("friend_requests")).
		Where(goqu.C("requester_id").Eq(requesterID), goqu.C("receiver_id").Eq(receiverID)).
		Delete().
		ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf(
			"unable to decline friend request for requester with id %d and receiver with id %d: %w",
			requesterID, receiverID, err,
		)
	}

	return nil
}

func (r *userRepository) UpdateAvatar(id int64, newAvatar []byte) error {
	// Prepared statement to handle byte parsing for avatar.
	// More info here: https://github.com/doug-martin/goqu/issues/254
	sql, args, _ := qb.Update("users").Prepared(true).Where(goqu.C("id").Eq(id)).Set(goqu.Record{
		"avatar": newAvatar,
	}).ToSQL()

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("unable to update avatar for user with id %d: %w", id, err)
	}

	return nil
}

func (r *userRepository) UpdateDisplayName(id int64, newDisplayName string) error {
	sql, _, _ := qb.Update("users").Where(goqu.C("id").Eq(id)).Set(goqu.Record{
		"display_name": newDisplayName,
	}).ToSQL()

	_, err := r.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("unable to update display name for user with id %d: %w", id, err)
	}

	return nil
}
