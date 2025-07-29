package service

import (
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) findByEmail(email string) (*domain.User, error) {
	// perform some DB operation
	// business logic
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) { // any = interface{}
	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	// generate token
	log.Println(user)

	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	return userInfo, err
}

func (s UserService) Login(email, password string) (string, error) {
	user, err := s.findByEmail(email)

	// compare password and generate token
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerificationCode(id uint, code int) error {

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]any, error) { // any = interface{}

	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]any, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]any, error) { // any = interface{}

	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (any, error) { // any = interface{}

	return nil, nil
}
