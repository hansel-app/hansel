package users

type User struct {
	ID             int64
	Username       string
	Email          string
	HashedPassword string
}
