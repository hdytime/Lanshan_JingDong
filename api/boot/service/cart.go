package service

import (
	"Lanshan_JingDong/api/boot/dao"
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCart(c *gin.Context) {
	var cart model.Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind"})
		return
	}
	flag := dao.SelectCart(cart.Name)
	if !flag {
		tx := global.MysqlDb.Begin()
		if tx.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction"})
			return
		}
		result := global.MysqlDb.Create(&cart)
		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create cart"})
			return
		}
		tx = tx.Commit()
		if tx.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "add cart successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "cart has existed"})
	}
}

func SettleCart(c *gin.Context) {
	cart := dao.GetCartInformation(c)
	if cart.Name == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "你的购物车空空如也~"})
		c.Abort()
		return
	}
	price := dao.GetPriceFromGoods(c, cart.Name)
	totalPrice := price * float64(cart.Quantity)
	username, _ := c.Get("username")
	money := dao.SelectMoneyFromUser(c, username)
	if money < totalPrice {
		c.JSON(http.StatusOK, gin.H{"msg": "don`t have enough money"})
		return
	} else {
		dao.SettleCartInformation(c, cart.Name, username.(string))
		dao.DeductMoneyFromUser(c, cart.Name, username.(string))
		flag := dao.ChangeGoodsQuantity(cart.Name)
		if !flag {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change goods quantity"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "buy goods successfully"})
		}
	}
}
