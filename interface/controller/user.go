package controller

import (
	"strconv"

	"github.com/facuellarg/users/interface/service"
	"github.com/facuellarg/users/models"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserService service.UserServiceI
}

func NewController(userService service.UserServiceI) *Controller {
	return &Controller{UserService: userService}
}

func (c *Controller) CreateUser(ctx echo.Context) error {
	newUser := &models.User{}
	if err := ctx.Bind(newUser); err != nil {
		return err
	}
	// userCreated, err := s.UserController.CreateUser(newUser)
	if user, err := c.UserService.GetUserById(newUser.Id); err != nil {
		return proccessError(err)
	} else if user != nil {
		return echo.NewHTTPError(409, "user already exists")
	}

	userCreated, err := c.UserService.CreateUser(newUser)
	if err != nil {
		return proccessError(err)
	}

	return ctx.JSON(201, userCreated)
}

func (c *Controller) GetUser(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(400, "Invalid id")
	}

	user, err := c.UserService.GetUserById(id)
	if err != nil {
		return proccessError(err)
	}

	if user == nil {
		return echo.NewHTTPError(404, "user not found")
	}

	return ctx.JSON(200, user)
}

func (c *Controller) DeleteUserById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(400, "Invalid id")
	}

	if user, err := c.UserService.GetUserById(id); err != nil {
		return proccessError(err)
	} else if user == nil {
		return echo.NewHTTPError(404, "user not found")
	}

	deleted, err := c.UserService.DeleteUserById(id)
	if err != nil {
		return proccessError(err)
	}

	if deleted {
		return ctx.NoContent(204)
	}

	return echo.NewHTTPError(500, "something went wrong, user not deleted")
}

func (c *Controller) UpdateUser(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(400, "Invalid id")
	}

	user := &models.User{}
	if err := ctx.Bind(user); err != nil {
		return err
	}
	user.Id = id

	oldUser, err := c.UserService.GetUserById(user.Id)
	if err != nil {
		return proccessError(err)
	}
	if oldUser == nil {
		return echo.NewHTTPError(404, "user not found")
	}

	userUpdated, err := c.UserService.UpdateUser(user)
	if err != nil {
		return proccessError(err)
	}

	if userUpdated != nil {
		return ctx.JSON(200, userUpdated)
	}

	return echo.NewHTTPError(500, "something went wrong, user not updated")
}

func proccessError(err error) error {
	if codeError, ok := err.(models.Error); ok {
		return echo.NewHTTPError(codeError.Code, codeError.Message)
	}
	return echo.NewHTTPError(500, err.Error())
}
