package dao

import (
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Create(c *gin.Context, s interface{}, errorMsg string, successMsg string, code int) {
	result := global.MysqlDb.Create(&s)
	if result.Error != nil {
		global.Logger.Error(errorMsg, zap.Error(result.Error))
		c.JSON(code, gin.H{"error": errorMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": successMsg})
}
