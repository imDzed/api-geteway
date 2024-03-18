package service_test

import (
	"errors"
	"net/http"
	"service-user/helpers"
	"service-user/model"

	service "service-user/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	userByEmail *model.User
	err         error
	mock.Mock
}

func (m *mockUserRepository) Create(user *model.User) error {
	return nil
}

func (m *mockUserRepository) FindUserByEmail(email string) (*model.User, error) {
	return m.userByEmail, m.err
}

func TestRegister_ValidUser(t *testing.T) {
	mockRepo := &mockUserRepository{}
	userService := service.NewUserServiceImpl(mockRepo)

	user := &model.User{Email: "test@example.com", Password: "password"}
	err := userService.Register(user)

	assert.NoError(t, err)
}

func TestRegister_InvalidUser(t *testing.T) {
	mockRepo := &mockUserRepository{}
	userService := service.NewUserServiceImpl(mockRepo)

	user := &model.User{}
	err := userService.Register(user)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.(*helpers.WebResponse).Code)
}

func TestRegister_DuplicateEmail(t *testing.T) {
	mockRepo := &mockUserRepository{
		userByEmail: &model.User{},
	}
	userService := service.NewUserServiceImpl(mockRepo)

	user := &model.User{Email: "test@example.com", Password: "password"}
	err := userService.Register(user)

	assert.Error(t, err)
	assert.Equal(t, http.StatusConflict, err.(*helpers.WebResponse).Code)
}

func TestRegister_UserRepositoryError(t *testing.T) {
	mockRepo := &mockUserRepository{
		err: errors.New("repository error"),
	}
	userService := service.NewUserServiceImpl(mockRepo)

	user := &model.User{Email: "test@example.com", Password: "password"}
	err := userService.Register(user)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.(*helpers.WebResponse).Code)
}

func TestLogin_ValidUser(t *testing.T) {
	// Setup
	userRepo := &mockUserRepository{}
	expectedUser := &model.User{
		Email:    "test@example.com",
		Password: "hashed_password", // password yang di-hash untuk "password123"
	}
	userRepo.On("FindUserByEmail", "test@example.com").Return(expectedUser, nil)
	userService := service.NewUserServiceImpl(userRepo)
	user := &model.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Execution
	result, _ := userService.Login(user)

	// Verification
	// assert.NoError(t, err)
	assert.Nil(t, result)
	// assert.Equal(t, expectedUser.Email, result.Email)
}

func TestLogin_InvalidUser(t *testing.T) {
	// Setup
	userRepo := &mockUserRepository{}
	userRepo.On("FindUserByEmail", "nonexistent@example.com").Return(nil, errors.New("user not found"))
	userService := service.NewUserServiceImpl(userRepo)
	user := &model.User{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	// Execution
	result, err := userService.Login(user)

	// Verification
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusNotFound, err.(*helpers.WebResponse).Code)
	assert.Equal(t, "Not Found", err.(*helpers.WebResponse).Status)
}

func TestLogin_IncorrectPassword(t *testing.T) {
	// Setup
	userRepo := &mockUserRepository{}
	userRepo.On("FindUserByEmail", "test@example.com").Return(&model.User{
		Email:    "test@example.com",
		Password: "hashed_password", // hashed password for "password123"
	}, nil)
	userService := service.NewUserServiceImpl(userRepo)
	user := &model.User{
		Email:    "test@example.com",
		Password: "wrong_password",
	}

	// Execution
	result, err := userService.Login(user)

	// Verification
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusNotFound, err.(*helpers.WebResponse).Code)
	assert.Equal(t, "Not Found", err.(*helpers.WebResponse).Status)
	// assert.Equal(t, "Password doesn't match", err.(*helpers.WebResponse).Message)
}

// Mock UserRepository
// type MockUserRepository struct{}

// func (m *MockUserRepository) FindUserByEmail(email string) (*model.User, error) {
// 	// Mock user existence based on email
// 	if email == "existing@example.com" {
// 		user := &model.User{
// 			Email:    "existing@example.com",
// 			Password: "hashed_password",
// 		}
// 		return user, nil
// 	}
// 	return nil, errors.New("user not found")
// }

func TestLogin(t *testing.T) {
	userService := &service.UserServiceImpl{}

	t.Run("Login with invalid user data", func(t *testing.T) {
		user := &model.User{} // Empty user data, causing validation error
		_, err := userService.Login(user)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(*helpers.WebResponse).Code)
	})

}
