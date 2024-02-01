package router

import (
	"Lanshan_JingDong/api/boot/middleware"
	"Lanshan_JingDong/api/boot/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", service.UserRegister)
	r.POST("/login", service.UserLogin)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/showhomepage", service.ShowHomepage)
		UserRouter.POST("/changehomepage", service.ChangeHomepage)
		UserRouter.GET("/checkgoodsinformation", service.CheckGoodsInformation)
		UserRouter.POST("/addcart", service.AddCart)
		UserRouter.POST("/settlecart", service.SettleCart)
		UserRouter.GET("/searchgoods", service.SearchGoods)
		UserRouter.POST("/reviewgoods", service.ReviewGoods)
		UserRouter.GET("/checkgoodsreviews", service.CheckGoodsReviews)
	}
	SellerRouter := r.Group("/seller")
	{
		SellerRouter.Use(middleware.JWTAuthMiddleware())
		SellerRouter.POST("/uploadgoodsinformation", service.UploadGoodsInformation)
	}
	err := r.Run()
	if err != nil {
		fmt.Println("failed to run gin")
		return
	}
}
