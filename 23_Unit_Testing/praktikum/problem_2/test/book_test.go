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

func InsertBook() error {
	books := models.Book{
		ID:        1,
		Title:     "Laskar Pelangi",
		Author:    "Giring Nidji",
		Publisher: "Mizan",
	}

	var err error
	if err = config.DB.Save(&books).Error; err != nil {
		return err
	}
	return nil
}

type TestCaseBook struct {
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

func TestGetBooks(t *testing.T) {

	testcase := TestCaseBook{

		Method:         http.MethodGet,
		Name:           "Get All Books",
		Path:           "/books",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success get all book",
	}

	e := InitEchoTestAPI()
	InsertBook()

	req := httptest.NewRequest(testcase.Method, testcase.Path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.GetBooksController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)

	}
}

func TestGetBook(t *testing.T) {

	testcase := TestCase{

		Method:         http.MethodGet,
		Name:           "Get book",
		Path:           "/",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success get book",
	}

	e := InitEchoTestAPI()
	InsertBook()

	req := httptest.NewRequest(testcase.Method, testcase.Path, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.GetBookController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestCreateBook(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPost,
		Name:           "Create book",
		Path:           "/books",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success add book",
	}

	e := InitEchoTestAPI()
	reqStr := `{
		"ID":        2,
		"title":     "Laskar Pelangi 2",
		"author":    "Giring",
		"publisher": "Mizan",
		}`
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.CreateBookController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestUpdateBook(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPut,
		Name:           "Update book",
		Path:           "/",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success Edit Book",
	}

	e := InitEchoTestAPI()
	InsertData()

	reqStr := `{
		"ID":        1,
		"title":     "Laskar Pelangi 3",
		"author":    "Giring",
		"publisher": "Mizan",
		}`
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.UpdateBookController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestDeleteBook(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodDelete,
		Name:           "Delete book",
		Path:           "/",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success Delete Book",
	}

	e := InitEchoTestAPI()
	InsertBook()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.DeleteBookController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}
}

// Api adalah kumpulan fungsi yang dapat digunakan oleh banyak orang Untuk melakukan proses / bisa diakses untuk mengoperasikan sebuah service.

// Rest Standarisasi untuk membuat memamnggil setiap api agar tidak pusing
