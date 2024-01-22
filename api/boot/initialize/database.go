package initialize

import (
	"Lanshan_JingDong/api/global"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func SetupDatabase() {
	setupMysql()
	setupRedis()
}

func setupMysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/京东?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database:%v", err)
	}
	global.MysqlDb = db
	global.Logger.Info("init mysql success")
}

func setupRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 地址
		Password: "",               // 密码，没有就是空字符串
		DB:       0,                // 使用默认的DB
	})
	_, err := rdb.Ping(global.Ctx).Result()
	if err != nil {
		log.Panicf("Redis ping failed" + err.Error())
	}
	global.RedisDb = rdb
	global.Logger.Info("init redis success")
}
