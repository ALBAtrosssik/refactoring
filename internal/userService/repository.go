package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, updates User) (User, error) {
	var user User

	result := r.db.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	if updates.Email != "" {
		user.Email = updates.Email
	}

	if updates.Password != "" {
		user.Password = updates.Password
	}

	result = r.db.Save(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	err := r.db.Delete(&user, id).Error
	return err
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(task User) (User, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return User{}, result.Error
	}

	return task, nil
}
