package mw

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"zqzqsb.com/gomall/app/user/conf"
)

func InitSession(h *server.Hertz) {
	config := conf.GetConf()
	store, err := redis.NewStore(10, "tcp",
		config.Redis.Address,
		config.Redis.Password,
		[]byte("your-session-secret-key"))
	if err != nil {
		panic(err)
	}
	h.Use(sessions.New("hertz-session", store))
	log.Println("init session success")
}
