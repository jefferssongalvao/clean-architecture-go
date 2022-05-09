package student

import (
	"clean-architecture/src/domain/value_objects"
	"errors"
)

type Student struct {
	cpf    value_objects.CPF
	name   string
	email  value_objects.Email
	phones []value_objects.Phone
}

func (student *Student) WithCpfNameEmail(cpf, name, email string) (Student, error) {
	var error error
	// set CPF
	if student.cpf, error = value_objects.CreateCPF(cpf); error != nil {
		return Student{}, error
	}

	// set NAME
	student.name = name

	// set Email
	if student.email, error = value_objects.CreateEmail(email); error != nil {
		return Student{}, error
	}

	if error := student.validate(); error != nil {
		return Student{}, error
	}

	return *student, nil
}

func (student *Student) AddPhone(ddd, number string) error {
	phone, error := value_objects.CreatePhone(ddd, number)
	if error != nil {
		return error
	}

	student.phones = append(student.phones, phone)

	return nil
}

func (student *Student) Cpf() string {
	return student.cpf.GetCpf()
}

func (student *Student) Name() string {
	return student.name
}

func (student *Student) Email() string {
	return student.email.GetEmail()
}

func (student *Student) Phones() []string {
	var phones []string
	for _, phone := range student.phones {
		phones = append(phones, phone.GetPhone())
	}
	return phones
}

func (student *Student) validate() error {
	if student.name == "" {
		return errors.New("invalid name")
	}
	return nil
}
