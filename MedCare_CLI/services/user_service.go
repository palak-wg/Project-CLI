package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
)

type UserService struct {
	repo interfaces.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *UserService) UpdateName(userID string, name string) error {
	return s.repo.UpdateName(userID, name)
}

func (s *UserService) UpdateAge(userID string, age string) error {
	return s.repo.UpdateAge(userID, age)
}

func (s *UserService) UpdateEmail(userID string, email string) error {
	return s.repo.UpdateEmail(userID, email)
}

func (s *UserService) UpdatePhoneNumber(userID string, phoneNumber string) error {
	return s.repo.UpdatePhoneNumber(userID, phoneNumber)
}

func (s *UserService) UpdateGender(userID string, gender string) error {
	return s.repo.UpdateGender(userID, gender)
}

func (s *UserService) UpdatePassword(userID string, newPassword string) error {
	return s.repo.UpdatePassword(userID, newPassword)
}
