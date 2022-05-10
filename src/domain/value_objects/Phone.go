package value_objects

import (
	"errors"
	"fmt"
)

type Phone struct {
	ddd    string
	number string
}

func (phone *Phone) validate() error {
	if len([]rune(phone.ddd)) != 2 {
		return errors.New("invalid DDD")
	}

	if len([]rune(phone.number)) > 9 || len([]rune(phone.number)) < 8 {
		return errors.New("invalid number")
	}

	return nil
}

func CreatePhone(ddd, number string) (Phone, error) {
	var phone Phone
	phone.ddd = ddd
	phone.number = number

	if error := phone.validate(); error != nil {
		return Phone{}, error
	}

	return phone, nil
}

func (phone *Phone) GetDDD() string {
	return phone.ddd
}

func (phone *Phone) GetNumber() string {
	return phone.number
}

func (phone *Phone) GetPhone() Phone {
	return *phone
}

func (phone *Phone) FormatPhone() string {
	return fmt.Sprintf("(%s) %s", phone.ddd, phone.number)
}
