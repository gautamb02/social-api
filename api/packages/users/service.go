package users

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
