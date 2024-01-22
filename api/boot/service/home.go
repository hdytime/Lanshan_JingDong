package service

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ShowHomepage(c *gin.Context) {
	username, _ := c.Get("username")
	var userinfo model.UserInfo
	result := global.MysqlDb.Where("username=?", username).First(&userinfo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user doesn`t exist",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": userinfo.Username,
		"gender":   userinfo.Gender,
		"sign":     userinfo.Sign,
	})
}

func ChangeHomepage(c *gin.Context) {
	var userinfo model.UserInfo
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to bind"})
		return
	}

	username, _ := c.Get("username")
	tx := global.MysqlDb.Begin()
	if tx.Error != nil {
		global.Logger.Error("transaction start failed", zap.Error(tx.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction start falied"})
		return
	}
	err = tx.Model(&model.UserInfo{}).Where("username=?", username).Updates(userinfo).Error
	if err != nil {
		tx.Rollback()
		global.Logger.Error("failed to update data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change user information"})
		return
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		global.Logger.Error("transaction commit falied", zap.Error(tx.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database commit error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user information updated successfully"})
}
