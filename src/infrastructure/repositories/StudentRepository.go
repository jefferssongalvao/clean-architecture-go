package repositories

import (
	"clean-architecture/src/domain/entities/student"
	"clean-architecture/src/domain/value_objects"
	"database/sql"
)

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *studentRepository {
	return &studentRepository{db}
}

func (studentRepository *studentRepository) Add(student student.Student) error {
	statement, error := studentRepository.db.Prepare("INSERT INTO students (cpf, name, email) values (?, ?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	result, error := statement.Exec(student.Cpf(), student.Name(), student.Email())
	if error != nil {
		return error
	}

	lastId, error := result.LastInsertId()
	if error != nil {
		return nil
	}

	statement, error = studentRepository.db.Prepare("INSERT INTO phones (ddd, number, student_id) values (?, ?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	for _, phone := range student.Phones() {
		_, error := statement.Exec(phone.GetDDD(), phone.GetNumber(), lastId)
		if error != nil {
			return error
		}
	}

	return nil
}

func (studentRepository *studentRepository) SearchForCpf(cpf value_objects.CPF) (student.Student, error) {
	return student.Student{}, nil
}

func (studentRepository *studentRepository) SearchAll() []student.Student {
	return []student.Student{}
}
