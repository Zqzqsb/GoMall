package service

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"zqzqsb.com/gomall/app/user/biz/dal/mysql"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test@test.com",
		Password:        "testpasswd@20241204",
		PasswordConfirm: "testpasswd@20241204",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
