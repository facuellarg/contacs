package main

import (
	"log"

	"github.com/facuellarg/users/external-service/server"
	"github.com/facuellarg/users/interface/controller"
	"github.com/facuellarg/users/interface/repository"
	"github.com/facuellarg/users/use-case/service"
)

func main() {
	userRepository := repository.GetInstance()
	userService := service.NewUserService(userRepository)
	userController := controller.NewController(userService)
	app := server.NewServerApp(userController, 8080)
	log.Fatal(app.Run())
}
