package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"fist_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verivied  bool      `json:"verified"`
	UserType  string    `json:"user_type"`
}
