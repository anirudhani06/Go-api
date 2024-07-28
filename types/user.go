package types

import (
	"fmt"
	"time"

	"github.com/anirudhani06/Go-api/utils"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LasttName string    `json:"last_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUserPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ValidateUser(user RegisterUserPayload) error {
	if !utils.IsValidEmail(user.Email) {
		return fmt.Errorf("please provide valid email")
	}

	if user.Password == "" {
		return fmt.Errorf("please provide the password")
	}
	if user.Password != user.ConfirmPassword {
		return fmt.Errorf("password not matching")
	}
	return nil
}
