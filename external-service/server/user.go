package server

import (
	"fmt"

	"github.com/facuellarg/users/external-service/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerApp struct {
	UserController controller.UserControllerI
	port           int
	engine         *echo.Echo
}

// NewServerApp creates a new NewServerApp
func NewServerApp(userController controller.UserControllerI, port int) *ServerApp {
	return &ServerApp{
		UserController: userController,
		port:           port,
		engine:         echo.New(),
	}
}

func (s *ServerApp) Run() error {
	s.RegisterRoutes()
	s.engine.Use(middleware.Logger())
	s.engine.Use(middleware.CORS())

	return s.engine.Start(fmt.Sprintf(":%d", s.port))
}

func (s *ServerApp) RegisterRoutes() {
	fmt.Println("Registering routes")
	s.engine.POST("/contacts", s.UserController.CreateUser)
	s.engine.GET("/contacts/:id", s.UserController.GetUser)
	s.engine.DELETE("/contacts/:id", s.UserController.DeleteUserById)
	s.engine.PUT("/contacts/:id", s.UserController.UpdateUser)
}
