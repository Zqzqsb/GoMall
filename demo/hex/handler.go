package main

import (
	"context"
	"zqzqsb.com/gomall/demo/hex/biz/service"
	example "zqzqsb.com/gomall/demo/hex/kitex_gen/hello/example"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	resp, err = service.NewHelloMethodService(ctx).Run(request)

	return resp, err
}
