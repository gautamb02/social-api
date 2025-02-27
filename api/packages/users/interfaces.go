package users

type UserService interface {
	GetUserByID(id int) *User
}

type UserRepository interface {
	GetUserByID(id int) *User
}
