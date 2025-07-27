package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {
}

func (s UserService) findByEmail(email string) (*domain.User, error) {
	// perform some DB operation
	// business logic

	return nil, nil
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) { // any = interface{}
	log.Println(input)

	return "this-is-my-token-as-of-end", nil
}

func (s UserService) Login(input any) (string, error) {

	return "", nil
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
