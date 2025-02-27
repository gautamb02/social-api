package users

import (
	"database/sql"
)

type userRepository struct {
	sqlDB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepository{
		sqlDB: db,
	}
}

func (*userRepository) GetUserByID(id int) *User {

	return &User{
		ID:    id,
		Name:  "Hello",
		Email: "World@1.com",
	}
}
