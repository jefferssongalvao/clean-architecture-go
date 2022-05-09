package student

import (
	"clean-architecture/src/domain/entities/student"
	"clean-architecture/src/domain/value_objects"
)

type StudentRepository struct {
}

func (studentRepository *StudentRepository) Add(student student.Student) error {
	return nil
}

func (studentRepository *StudentRepository) SearchForCpf(cpf value_objects.CPF) (student.Student, error) {
	return student.Student{}, nil
}

func (studentRepository *StudentRepository) SearchAll() []student.Student {
	return []student.Student{}
}
