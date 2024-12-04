package service

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"zqzqsb.com/gomall/app/user/biz/dal/mysql"
	user "zqzqsb.com/gomall/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "test@test.com",
		Password: "testpasswd@20241204",
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
