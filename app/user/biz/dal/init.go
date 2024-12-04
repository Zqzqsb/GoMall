package dal

import (
	"zqzqsb.com/gomall/app/user/biz/dal/mysql"
	"zqzqsb.com/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
