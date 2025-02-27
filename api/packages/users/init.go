package users

import "database/sql"

func NewUserModule(db *sql.DB) *UserHandler {
	userRepo := NewUserRepo(db)             // Initialize repository
	userService := NewUserService(userRepo) // Initialize service
	return NewUserHandler(userService)      // Initialize handler
}
