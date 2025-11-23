package handler

import (
	"errors"
	"livecode/internal/adapter/repository"
	"livecode/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService usecase.UserService
}

func NewUserHandler(userService usecase.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/users/:id", u.GetById)
}

func (u *UserHandler) GetById(c echo.Context) error {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := u.userService.GetByID(userId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}
