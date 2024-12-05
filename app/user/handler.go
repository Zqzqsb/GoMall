package main

import (
	"context"
	user "zqzqsb.com/gomall/app/user/kitex_gen/user"
	"zqzqsb.com/gomall/app/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Hello implements the UserServiceImpl interface.
func (s *UserServiceImpl) Hello(ctx context.Context, req *user.HelloReq) (resp *user.HelloResp, err error) {
	resp, err = service.NewHelloService(ctx).Run(req)

	return resp, err
}
