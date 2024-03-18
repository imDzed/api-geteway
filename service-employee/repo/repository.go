package repo

import (
	"service-employee/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(user *model.Employee) error
}

type employeeRepositoryImplement struct {
	db *gorm.DB
}

func New(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImplement{
		db: db,
	}
}

func (e *employeeRepositoryImplement) Create(user *model.Employee) error {

	var requestBody model.Employee
	requestBody.ID = string(uuid.New().String())
	if err := e.db.Create(requestBody).Error; err != nil {
		return err
	}
	return nil
}
