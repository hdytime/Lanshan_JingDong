package dao

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SelectCart(name string) bool {
	var cart model.Cart
	result := global.MysqlDb.Where("name=?", name).First(&cart)
	if result.Error != nil {
		return false
	}
	return true
}

func GetPriceFromGoods(c *gin.Context, name string) (price float64) {
	var goods model.Goods
	flag := SelectGoods(name, goods)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cart doesn`t exist"})
		return
	} else {
		global.MysqlDb.Where("name=?", name).First(&goods)
		return goods.Price
	}
}

func SettleCartInformation(c *gin.Context, name string) {
	var cart model.Cart
	result := global.MysqlDb.Where("name=?", name).Delete(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to settle cart information"})
		return
	}
}
