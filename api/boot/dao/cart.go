package dao

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func SettleCartInformation(c *gin.Context, name string, username string) {
	tx := global.MysqlDb.Begin()
	var cart model.Cart
	result := global.MysqlDb.Where("name=?", name).Delete(&cart)
	if result.Error != nil {
		tx.Rollback()
		global.Logger.Error("failed to settle cart information", zap.Error(result.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to settle cart information"})
		c.Abort()
		return
	}
	var goods model.Goods
	result = global.MysqlDb.Where("name=?", name).First(&goods)
	if result.Error != nil {
		tx.Rollback()
		global.Logger.Error("查找商品数据失败", zap.Error(result.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "购物失败"})
		c.Abort()
		return
	}
	goodsID := goods.GoodsID
	key := username + "-" + strconv.FormatInt(goodsID, 10)
	err := global.RedisDb.Set(global.Ctx, key, 0, 0).Err()
	if err != nil {
		tx.Rollback()
		global.Logger.Error("创建用户商品购买关系失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "购物失败"})
		c.Abort()
		return
	}
}

func DeductMoneyFromUser(c *gin.Context, name string, username string) {
	tx := global.MysqlDb.Begin()
	var goods model.Goods
	result := global.MysqlDb.Model(&model.Goods{}).Where("name=?", name).First(&goods)
	if result.Error != nil {
		tx.Rollback()
		global.Logger.Error("查找商品金额失败")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "购物失败"})
		return
	}
	price := goods.Price //商品的价格
	result = global.MysqlDb.Model(&model.User{}).Where("username=?", username).Update("money", gorm.Expr("money-?", price))
	if result.Error != nil {
		tx.Rollback()
		global.Logger.Error("修改用户金额失败", zap.Error(result.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "购物失败"})
		return
	}
	tx.Commit()
}

func GetCartInformation(c *gin.Context) model.Cart {
	var cart model.Cart
	tx := global.MysqlDb.Begin()
	result := global.MysqlDb.Model(model.Cart{}).First(&cart)
	if result.Error != nil {
		tx.Rollback()
		global.Logger.Error("查找购物车数据失败", zap.Error(result.Error))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "结算购物车失败"})
		c.Abort()
		return model.Cart{}
	}
	tx.Commit()
	return cart
}
