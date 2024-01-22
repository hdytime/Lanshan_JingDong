package service

import (
	"Lanshan_JingDong/api/boot/dao"
	"Lanshan_JingDong/api/boot/middleware"
	"Lanshan_JingDong/api/boot/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//检查有无此用户
	flagInSelcet := dao.SelectUserInLogin(user.Username)
	if !flagInSelcet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doesn`t exist"})
		return
	}
	//判断密码是否正确
	flagInJudge := dao.JudgeUserInLogin(user.Username, user.Password)
	if !flagInJudge {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}
	//JWT认证
	token := middleware.NewToken(c, user.Username)
	c.JSON(http.StatusOK, gin.H{"massage": "Login successfully",
		"token": token})
}
