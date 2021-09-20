package users

type User struct {
	ID             int64
	Username       string
	DisplayName    string
	Email          string
	HashedPassword string
}
