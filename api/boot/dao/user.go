package dao

import (
	"Lanshan_JingDong/api/boot/model"
	"Lanshan_JingDong/api/global"
)

func SelectUserInRegister(username string) bool {
	var existingUser model.User
	result := global.MysqlDb.Where("username=?", username).First(&existingUser)
	if result.Error == nil {
		//用户名已经存在
		return true
	} else {
		return false
	}
}

func SelectUserInLogin(username string) bool {
	var user model.User
	result := global.MysqlDb.Where("username=?", username).First(&user)
	if result.Error == nil {
		return true
	} else {
		return false
	}
}

func JudgeUserInLogin(username string, password string) bool {
	var user model.User
	global.MysqlDb.Where("username=?", username).First(&user)
	if password != user.Password {
		return false
	} else {
		return true
	}
}
