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
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "securePassword123",
		ID:        1,
	}
}

func (r *userRepository) CreateUser(user User) (int64, error) {
	result, err := r.sqlDB.Exec(CreateUserQuery, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertedID, nil
}
