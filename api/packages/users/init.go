package users

import "database/sql"

func NewUserModule(db *sql.DB) *UserHandler {
	userRepo := NewUserRepo(db)
	userService := NewUserService(userRepo)
	return NewUserHandler(userService)
}
