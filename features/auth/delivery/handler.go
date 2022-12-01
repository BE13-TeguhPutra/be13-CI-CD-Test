package delivery

import (
	"be13/clean-arch/features/auth"
	"be13/clean-arch/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authServices auth.ServiceInterface
}

func NewAuth(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthDelivery{
		authServices: service,
	}

	e.POST("/login", handler.login)

}

func (delivery *AuthDelivery) login(c echo.Context) error {
	authInput := AuthRequest{}
	errBind := c.Bind(&authInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	token, err := delivery.authServices.Login(authInput.Email, authInput.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}
	data := map[string]any{
		"token": token,
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", data))

}
