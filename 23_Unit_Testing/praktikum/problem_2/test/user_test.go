package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/config"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/controller"
	"github.com/gozzafadillah/23_Unit_Testing/praktikum/problem_2/models"
	"github.com/labstack/echo/v4"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	config.InitMigrateTest()
	e := echo.New()
	return e
}

func InsertData() error {
	user := models.User{
		ID:       1,
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

type TestCase struct {
	Method         string
	Name           string
	Path           string
	ExpectStatus   int
	ExpectResponse string
}

/*
===> CARA TEST NYA YA
 go test ./... -v -coverpkg=./controller/...,./lib/...,./model/... -coverprofile=cover.out && go tool cover -html=cover.out
*/

func TestLoginUser(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPost,
		Name:           "Login User",
		Path:           "/login",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success login",
	}

	e := InitEchoTestAPI()
	InsertData()
	reqStr := `{
		"id": 1,
		"name": "Alta",
		"email": "alta@gmail.com",
		"password": "123"
	}`

	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.LoginUserController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)

	}
}

func TestGetUsers(t *testing.T) {

	testcases := []TestCase{
		{
			Method:         http.MethodGet,
			Name:           "Get All Users",
			Path:           "/Users",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "users",
		},
	}

	e := InitEchoTestAPI()
	InsertData()

	for _, testcase := range testcases {
		req := httptest.NewRequest(testcase.Method, testcase.Path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)
		if assert.NoError(t, controller.GetUsersController(c)) {
			assert.Equal(t, testcase.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
		}

	}
}

func TestGetUser(t *testing.T) {

	testcase := TestCase{

		Method:         http.MethodGet,
		Name:           "Get User",
		Path:           "/",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success get user",
	}

	e := InitEchoTestAPI()
	InsertData()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.GetUserController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestCreateUser(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPost,
		Name:           "Create user",
		Path:           "/Users",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success add user",
	}

	e := InitEchoTestAPI()
	reqStr := `{
		"id": 2,
		"name":     "Alta",
		"password": "123",
		"email":    "alta@gmail.com"
		}`
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"
	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.CreateUserController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestUpdateUser(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPut,
		Name:           "Update user",
		Path:           "/",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success Edit User",
	}

	e := InitEchoTestAPI()
	InsertData()

	reqStr := `{
		"id":       1,
		"name":     "Alterra",
		"password": "123",
		"email":    "alta@gmail.com"
		}`
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.UpdateUserController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestDeleteUser(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodDelete,
		Name:           "Delete User",
		Path:           "/",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success Delete User",
	}

	e := InitEchoTestAPI()
	InsertData()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.DeleteUserController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}
}

// Api adalah kumpulan fungsi yang dapat digunakan oleh banyak orang Untuk melakukan proses / bisa diakses untuk mengoperasikan sebuah service.

// Rest Standarisasi untuk membuat memamnggil setiap api agar tidak pusing
