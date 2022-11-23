package main

import (
	"gin-mysql-jwt/controller"
	"gin-mysql-jwt/service"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	usersGroup := r.Group("/users")
	userService := service.NewUserService()
	jwtService := service.NewJwtService()
	controller.NewUserController(usersGroup, userService, jwtService)

}
