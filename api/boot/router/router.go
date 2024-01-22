package router

import (
	"Lanshan_JingDong/api/boot/middleware"
	"Lanshan_JingDong/api/boot/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/register", service.UserRegister)
	r.POST("/login", service.UserLogin)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/showhomepage", service.ShowHomepage)
		UserRouter.POST("changehomepage", service.ChangeHomepage)
	}
	err := r.Run()
	if err != nil {
		fmt.Println("failed to run gin")
		return
	}

}
