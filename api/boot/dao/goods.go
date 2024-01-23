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
