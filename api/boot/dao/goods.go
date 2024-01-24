package dao

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
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
