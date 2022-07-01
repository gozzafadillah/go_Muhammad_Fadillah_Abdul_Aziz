package controller_users

import (
	request_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/controllers/users/request"
	response_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/controllers/users/response"
	domain_users "github.com/gozzafadillah/24_Clean_and_Hexagonal_Architecture/praktikum/domain/users"
	"github.com/labstack/echo/v4"
)

type ControllerUser struct {
	UserBusiness domain_users.Business
}

func NewUsersController(uc domain_users.Business) ControllerUser {
	return ControllerUser{
		UserBusiness: uc,
	}
}

func (uc *ControllerUser) Register(ctx echo.Context) error {
	req := request_users.RequestJSONUsers{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	data, err := uc.UserBusiness.Register(request_users.ToDomainUsers(req))
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, echo.Map{
		"data": response_users.FromDomainUsers(data),
	})
}

func (uc *ControllerUser) GetAllUser(ctx echo.Context) error {
	data := uc.UserBusiness.GetUsers()
	sliceUsers := []response_users.ResponseJSONUsers{}

	if len(data) == 0 {
		return ctx.JSON(500, echo.Map{
			"error": "data empty",
		})
	}
	for _, value := range data {
		sliceUsers = append(sliceUsers, response_users.FromDomainUsers(value))
	}
	return ctx.JSON(200, echo.Map{
		"data": sliceUsers,
	})
}
