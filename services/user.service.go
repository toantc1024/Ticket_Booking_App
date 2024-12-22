package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"tickets/models"
	"tickets/repositories"
)

type UserService interface {
	CreateUser(user *models.CreateUserSchema) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, user *models.UpdateUserSchema) (*models.User, error)
	DeleteUser(id string) error
	Login(email, password string) (*models.User, error)
	Logout(userID string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *models.CreateUserSchema) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
		Role:     models.UserRole,
	}

	err = s.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id string, user *models.UpdateUserSchema) (*models.User, error) {
	existingUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existingUser.Password = string(hashedPassword)
	}

	err = s.repo.UpdateUser(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("Email hoặc mật khẩu chưa đúng")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Email hoặc mật khẩu chưa đúng")
	}

	// TODO: Sẽ triển khai JWT tại đây

	return user, nil
}

func (s *userService) Logout(userID string) error {
	// TODO: Sẽ triển khai JWT tại đây

	return nil
}
