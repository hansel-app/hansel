package users

type Repository interface {
	Get(id int64) (*User, error)
	Add(user *User) error
	Update(id int64, user *User) error
	Remove(id int64) error
}
