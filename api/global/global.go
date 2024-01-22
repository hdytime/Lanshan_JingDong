package global

import (
	"Lanshan_JingDong/api/global/config"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	MysqlDb *gorm.DB
	RedisDb *redis.Client
	Ctx     = context.Background()
)
