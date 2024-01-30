package service

import (
	"Lanshan_JingDong/api/boot/dao"
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userinfo := model.UserInfo{Username: user.Username}

	flag := dao.SelectUserInRegister(user.Username)
	if flag {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exist"})
		return
	}
	//执行事务
	tx := global.MysqlDb.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction start failed"})
		return
	}
	//密码加密
	salt := dao.GenerateSalt()
	passwordEncrypt, err := dao.PasswordEncrypt(user.Password, salt)
	if err != nil {
		global.Logger.Error("用户密码加密失败")
		return
	}
	user.Password = passwordEncrypt
	user.Salt = salt
	//用户默认给10000元
	user.Money = 10000

	result := global.MysqlDb.Create(&user)
	if result.Error != nil {
		global.Logger.Error("failed to register", zap.Error(result.Error))
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to register"})
		return
	}

	result = global.MysqlDb.Create(&userinfo)
	if result.Error != nil {
		global.Logger.Error("failed to register", zap.Error(result.Error))
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register"})
		return
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction commit failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "register successfully"})
}
