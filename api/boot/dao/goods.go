package dao

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
	"strconv"
)

func SelectGoods(name string, goods model.Goods) bool {
	result := global.MysqlDb.Model(&model.Goods{}).Where("name=?", name).First(&goods)
	if result.Error != nil {
		return false
	}
	return true
}

func ChangeGoodsQuantity(name string) bool {
	var goods model.Goods
	global.MysqlDb.Where("name=?", name).First(&goods)
	result := global.MysqlDb.Model(&model.Goods{}).Where("name=?", name).Update("quantity", goods.Quantity-1)
	if result.Error != nil {
		return false
	}
	return true
}

func CheckCommentPermission(username string, goodsId int64) bool {
	key := username + "-" + strconv.FormatInt(goodsId, 10)
	value, err := global.RedisDb.Get(global.Ctx, key).Result()
	if err != nil {
		//没有对应记录即未购买过该商品
		return false
	} else if value == "0" {
		return true
	} else {
		return false
	}
}
