package handler

import (
	"github.com/Sohaib03/GoTemplate/model"
	"github.com/Sohaib03/GoTemplate/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{Email: "Sohaib"}
	return render(c, user.Show(u))
}
