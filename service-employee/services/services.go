package service

import (
	"net/http"
	"service-employee/model"

	helpers "service-employee/helper"
	"service-employee/repo"
)

type EmployeeService interface {
	CreateEmployee(employee *model.Employee) error
	ConnectUserService(user_uri string, access_token string) (*http.Response, error)
}

type employeeServiceImpl struct {
	employeeRepository repo.EmployeeRepository
}

func NewEmployeeServiceImpl(repository repo.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{employeeRepository: repository}
}

func (employeeService *employeeServiceImpl) CreateEmployee(employee *model.Employee) error {
	if err := employee.Validate(); err != nil {
		return &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    employee,
		}
	}

	return employeeService.employeeRepository.Create(employee)
}

func (employeeService *employeeServiceImpl) ConnectUserService(user_uri string, access_token string) (*http.Response, error) {

	if len(access_token) == 0 {
		return nil, &helpers.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "401 unauthorized",
			Message: "Invalid token: Access token missing",
		}
	}

	req, err := http.NewRequest("GET", user_uri+"/auth", nil)
	if err != nil {
		return nil, &helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal server error",
			Message: "Failed to create request",
		}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, &helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal server error",
			Message: "Failed to perform request",
		}
	}
	return resp, nil
}
