package dal

import (
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	"zqzqsb/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
