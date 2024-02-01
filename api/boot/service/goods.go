package service

import (
	"Lanshan_JingDong/api/boot/dao"
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	flag := dao.SelectGoods(goods.Name, goods)

	if !flag {
		//没有对应商品
		goods.GoodsID = dao.NewSnowflakeID()
		goods.Reviews = 0
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
	} else {
		//有对应商品
		result := global.MysqlDb.Model(&model.Goods{}).Where("name=?", goods.Name).Updates(&goods)
		if result.Error != nil {
			tx.Rollback()
			global.Logger.Error("failed to update goods information", zap.Error(result.Error))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update goods information"})
			return
		}
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
			global.Logger.Error("failed to commit transaction", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update goods information"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "update goods information successfully"})
	}
}

func CheckGoodsInformation(c *gin.Context) {
	var goods []model.Goods
	//err := c.ShouldBindJSON(&goods)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind"})
	//	return
	//}
	//flag := dao.SelectGoods(goods.Name, goods)
	//if !flag {
	//	//没有对应商品
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "goods doesn`t exist"})
	//	return
	//} else {
	//有对应商品
	//global.MysqlDb.Where("name=?", goods.Name).First(&goods)
	//c.JSON(http.StatusOK, gin.H{
	//	"name":     goods.Name,
	//	"price":    goods.Price,
	//	"quantity": goods.Quantity,
	//})
	//}
	result := global.MysqlDb.Find(&goods)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "empty goods information"})
		return
	}
	c.JSON(http.StatusOK, goods)
}

func SearchGoods(c *gin.Context) {
	var goods model.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind"})
		return
	}

	result := global.MysqlDb.Where("name=?", goods.Name).First(&goods)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "没有此商品"})
		return
	}

	c.JSON(http.StatusOK, goods)
}

func ReviewGoods(c *gin.Context) {
	tx := global.MysqlDb.Begin()
	var goods model.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind"})
		return
	}

	var existingGoods model.Goods
	result := global.MysqlDb.Where("name=?", goods.Name).First(&existingGoods)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询数据库失败"})
		return
	}
	username, _ := c.Get("username")
	reviews := model.Reviews{
		GoodsID:  existingGoods.GoodsID,
		Username: username.(string),
		Content:  goods.Content,
		Name:     existingGoods.Name,
	}

	flag := dao.CheckCommentPermission(username.(string), existingGoods.GoodsID)
	if flag {
		result := global.MysqlDb.Model(&model.Reviews{}).Create(&reviews)
		if result.Error != nil {
			tx.Rollback()
			global.Logger.Error("添加评论到数据库失败")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败"})
			return
		}

		result = global.MysqlDb.Model(&model.Goods{}).Where("name=?", goods.Name).Update("reviews", gorm.Expr("reviews+?", 1))
		if result.Error != nil {
			tx.Rollback()
			global.Logger.Error("更新评论数加一失败", zap.Error(result.Error))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败"})
			return
		}

		result = global.MysqlDb.Where("name=?", goods.Name).First(&existingGoods)

		key := username.(string) + "-" + strconv.FormatInt(existingGoods.GoodsID, 10)
		err = global.RedisDb.Set(global.Ctx, key, 1, 0).Err()
		if err != nil {
			tx.Rollback()
			global.Logger.Error("用户商品对应关系写入redis失败")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "评论成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "你未购买该商品或已经评论过该商品"})
		return
	}
}

func CheckGoodsReviews(c *gin.Context) {
	var reviews []model.Reviews
	var goods model.Goods
	err := c.ShouldBindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind"})
		return
	}
	tx := global.MysqlDb.Begin()
	result := global.MysqlDb.Model(&model.Reviews{}).Where("name=?", goods.Name).Find(&reviews)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询商品评论失败"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}
