package dal

import (
	"zqzqsb.com/gomall/demo/hex/biz/dal/mysql"
	"zqzqsb.com/gomall/demo/hex/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
