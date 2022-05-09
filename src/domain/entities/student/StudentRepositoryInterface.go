package student

import "clean-architecture/src/domain/value_objects"

type StudentRepositoryInterface interface {
	Add(student Student) error
	SearchForCpf(cpf value_objects.CPF) (Student, error)
	SearchAll() []Student
}
