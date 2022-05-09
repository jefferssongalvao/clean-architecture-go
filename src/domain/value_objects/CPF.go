package value_objects

import (
	"errors"

	"github.com/paemuri/brdoc"
)

type CPF struct {
	number string
}

func CreateCPF(number string) (CPF, error) {
	cpf := CPF{number: number}
	if error := cpf.validate(); error != nil {
		return CPF{}, error
	}
	return cpf, nil
}

func (cpf *CPF) validate() error {
	if brdoc.IsCPF(cpf.number) {
		return errors.New("invalid cpf")
	}
	return nil
}

func (cpf *CPF) GetCpf() string {
	return cpf.number
}
