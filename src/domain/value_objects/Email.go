package value_objects

import (
	"errors"

	"github.com/badoux/checkmail"
)

type Email struct {
	address string
}

func (email *Email) validate() error {
	if error := checkmail.ValidateFormat(email.address); error != nil {
		return errors.New("e-mail invalid")
	}
	return nil
}

func CreateEmail(address string) (Email, error) {
	var email Email
	email.address = address
	if error := email.validate(); error != nil {
		return Email{}, error
	}
	return email, nil
}

func (email *Email) GetEmail() string {
	return email.address
}
