package service

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func UploadGoodsInformation(c *gin.Context) {
	var goods model.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind"})
		return
	}
	tx := global.MysqlDb.Begin()
	if tx.Error != nil {
		global.Logger.Error("transaction start failed", zap.Error(tx.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "transaction start failed"})
		tx.Rollback()
		return
	}
	result := global.MysqlDb.Create(&goods)
	if result.Error != nil {
		tx.Rollback()
		global.Logger.Error("failed to upload goods information", zap.Error(result.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload goods information"})
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		global.Logger.Error("failed to commit transaction", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload goods information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "upload goods information successfully"})

}
