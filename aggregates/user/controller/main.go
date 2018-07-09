package userController

import (
	"github.com/amiralii/goEchoExample/config/response"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	Repo "github.com/amiralii/goEchoExample/aggregates/user/repository"
	"github.com/labstack/echo"
)

func Signup(e echo.Context) error {
	u := new(user.User)
	if err := e.Bind(u); err != nil {
		return response.Error(err.Error(), 500)
	}
	if err := e.Validate(u); err != nil {
		return response.Error(err.Error(), 422)
	}
	err := Repo.Signup(*u)
	if err != nil {
		return response.Error(err.Error(), 500)
	}
	return response.Created(e, "user created")
}

func Signin(e echo.Context) error {
	u := new(user.User)
	if err := e.Bind(u); err != nil {
		return response.Error(err.Error(), 500)
	}
	data, err := Repo.FindUser(u.Username, u.Password)
	if err != nil {
		return response.Error(err.Error(), 500)
	}
	return response.Ok(e, data)
}
