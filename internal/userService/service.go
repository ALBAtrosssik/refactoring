package userService

type UserService struct {
	userRepo UserRepository
}

func NewUserService(repo *userRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.userRepo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.userRepo.DeleteUserByID(id)
}
