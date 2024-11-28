package dal

import (
	"zqzqsb/gomall/app/frontend/biz/dal/mysql"
	"zqzqsb/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
