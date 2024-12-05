package service

import (
	"context"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
)

type HelloService struct {
	ctx context.Context
} // NewHelloService new HelloService
func NewHelloService(ctx context.Context) *HelloService {
	return &HelloService{ctx: ctx}
}

// Run create note info
func (s *HelloService) Run(req *user.HelloReq) (resp *user.HelloResp, err error) {
	// Finish your business logic.

	return
}
