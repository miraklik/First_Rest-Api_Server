package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               int    `json:"ID"`
	Email            string `json:"email"`
	Password         string `json:"password,omitempty"`
	EncryptoPassword string `json:"-"`
}

func (u *User) Validate() bool {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIF(u.EncryptoPassword == "")), validation.Length(6, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptoString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptoPassword = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptoPassword), []byte(password)) == nil
}

func encryptoString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
