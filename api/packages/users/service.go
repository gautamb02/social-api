package users

import "github.com/gautamb02/social-api/shared/utils"

type userService struct {
	userRepo UserRepository
}

func NewUserService(repo UserRepository) UserService { // Pass interface, not pointer
	return &userService{
		userRepo: repo,
	}
}

// FindUserByID retrieves a user from the repository
func (s *userService) GetUserByID(id int) *User {
	return s.userRepo.GetUserByID(1)
}

func (s *userService) CreateUser(user User) (int64, error) {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashed
	return s.userRepo.CreateUser(user)
}
