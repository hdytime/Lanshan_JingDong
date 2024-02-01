package dao

import (
	"Lanshan_JingDong/api/global"
	"github.com/bwmarrin/snowflake"
)

func NewSnowflakeID() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		global.Logger.Error("new node failed")
		return 0
	}
	id := node.Generate()
	return int64(id)
}
