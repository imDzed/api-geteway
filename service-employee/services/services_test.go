package service

import (
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	helpers "service-employee/helper"
	"service-employee/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockEmployeeRepository struct {
	mock.Mock
}

func (m *mockEmployeeRepository) Create(employee *model.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func TestCreateEmployee_Success(t *testing.T) {
	repoMock := new(mockEmployeeRepository)

	repoMock.On("Create", mock.Anything).Return(nil)

	employeeService := NewEmployeeServiceImpl(repoMock)

	employee := &model.Employee{
		Name: "John Doe",
	}

	err := employeeService.CreateEmployee(employee)

	assert.NoError(t, err)

	repoMock.AssertCalled(t, "Create", employee)
}

func TestCreateEmployee_InvalidData(t *testing.T) {
	repoMock := new(mockEmployeeRepository)

	employeeService := NewEmployeeServiceImpl(repoMock)

	employee := &model.Employee{
		Name: "",
	}

	err := employeeService.CreateEmployee(employee)

	assert.Error(t, err)

	webErr, ok := err.(*helpers.WebResponse)
	assert.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, webErr.Code)

	repoMock.AssertNotCalled(t, "Create")
}

func TestConnectUserService_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	employeeService := &employeeServiceImpl{}

	resp, err := employeeService.ConnectUserService(server.URL, "valid_access_token")

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestConnectUserService_InvalidToken(t *testing.T) {
	employeeService := &employeeServiceImpl{}

	resp, err := employeeService.ConnectUserService("http://localhost:3000", "")

	webErr, ok := err.(*helpers.WebResponse)
	assert.True(t, ok)
	assert.Equal(t, http.StatusUnauthorized, webErr.Code)
	assert.Equal(t, "401 unauthorized", webErr.Status)
	assert.Equal(t, "Invalid token: Access token missing", webErr.Message)

	assert.Nil(t, resp)
}

func TestConnectUserService_RequestError(t *testing.T) {
	employeeService := &employeeServiceImpl{}

	resp, err := employeeService.ConnectUserService("http://invalid-url:3000", "valid_access_token")

	webErr, ok := err.(*helpers.WebResponse)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, webErr.Code)
	assert.Equal(t, "Internal server error", webErr.Status)
	assert.Equal(t, "Failed to perform request", webErr.Message)

	assert.Nil(t, resp)
}
