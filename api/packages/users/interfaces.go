package users

type UserService interface {
	GetUserByID(id int) *User
	CreateUser(user User) (int64, error)
}

type UserRepository interface {
	GetUserByID(id int) *User
	CreateUser(user User) (int64, error)
}
