package business_users_test

import (
	"errors"
	"os"
	"testing"

	middlewares "github.com/gozzafadillah/27_Docker/praktikum/app/middleware"
	business_users "github.com/gozzafadillah/27_Docker/praktikum/businesses/users"
	domain_users "github.com/gozzafadillah/27_Docker/praktikum/domain/users"
	RepoUsersMock "github.com/gozzafadillah/27_Docker/praktikum/domain/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userBusiness   domain_users.Business
	userDomainUser domain_users.Users
	userRepo       RepoUsersMock.Repository
)

func TestMain(m *testing.M) {
	userBusiness = business_users.NewUsersBusiness(&userRepo, &middlewares.ConfigJwt{})
	userDomainUser = domain_users.Users{
		ID:       1,
		Email:    "gozzafadillah@gmail.com",
		Password: "123456",
	}
	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T) {
	t.Run("get users", func(t *testing.T) {
		userRepo.On("GetAllUser").Return([]domain_users.Users{userDomainUser, {ID: 2, Email: "Kiana@gmail.com", Password: "123456"}}).Once()
		res := userBusiness.GetUsers()

		assert.Equal(t, userDomainUser.Email, res[0].Email)
	})
}

func TestRegister(t *testing.T) {
	t.Run("success register", func(t *testing.T) {
		userRepo.On("Store", mock.Anything).Return(userDomainUser, nil).Once()
		res, err := userBusiness.Register(userDomainUser)
		assert.NoError(t, err)
		assert.Equal(t, userDomainUser.Email, res.Email)
	})
	t.Run("failed register", func(t *testing.T) {
		userRepo.On("Store", mock.Anything).Return(domain_users.Users{}, errors.New("register failed")).Once()
		res, err := userBusiness.Register(userDomainUser)
		assert.Error(t, err)
		assert.Equal(t, domain_users.Users{}, res)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success login", func(t *testing.T) {
		userRepo.On("CheckEmailPassword", mock.Anything, mock.Anything).Return(userDomainUser, nil).Once()
		res, err := userBusiness.Login(userDomainUser.Email, userDomainUser.Password)
		assert.NoError(t, err)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc", res)
	})
	t.Run("failed login", func(t *testing.T) {
		userRepo.On("CheckEmailPassword", mock.Anything, mock.Anything).Return(domain_users.Users{}, errors.New("email password miss macth")).Once()
		res, err := userBusiness.Login(userDomainUser.Email, userDomainUser.Password)
		assert.Error(t, err)
		assert.Equal(t, "", res)
	})
}

// go test ./praktikum/businesses/users/usecase_test.go -coverpkg=./praktikum/businesses/users/...
