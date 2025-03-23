package auth

import (
	"errors"
	"explorekg/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Login(email string, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindUserByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredetials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredetials)
	}
	return existedUser.Email, nil
}

func (service *AuthService) Register(name, email, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindUserByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := user.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	_, err = service.UserRepository.CreateUser(&user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
