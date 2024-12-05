package service

import (
	"context"
	"testing"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

func TestHello_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloService(ctx)
	// init req and assert value

	req := &user.HelloReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
