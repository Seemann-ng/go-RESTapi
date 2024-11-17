package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Validate ...
func (user *User) Validate() error {
	return validation.ValidateStruct(
		user,
		validation.Field(
			&user.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&user.Password,
			validation.By(requiredIf(user.EncryptedPassword == "")),
			validation.Length(6, 30),
		),
	)
}

// BeforeCreate ...
func (user *User) BeforeCreate() error {
	if len(user.Password) > 0 {
		enc, err := encryptString(user.Password)
		if err != nil {
			return err
		}

		user.EncryptedPassword = enc
	}

	return nil
}

// Sanitize ...
func (user *User) Sanitize() {
	user.Password = ""
}

// ComparePassword ...
func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) == nil
}

// encryptString ...
func encryptString(str string) (string, error) {
	arr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(arr), nil
}
