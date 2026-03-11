package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
	Password          string `json:"password"`
}

func (user *User) BeforeCreate() error {
	if len(user.Password) > 0 {
		encPass, err := encryptString(user.Password)
		if err != nil {
			return err
		}
		user.EncryptedPassword = encPass
	}

	return nil
}

func (user *User) Validate() error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Length(6, 100)),
	)
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
