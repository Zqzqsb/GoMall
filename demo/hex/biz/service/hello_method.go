package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	example "zqzqsb.com/gomall/demo/hex/kitex_gen/hello/example"
)

type HelloMethodService struct {
	ctx context.Context
} // NewHelloMethodService new HelloMethodService
func NewHelloMethodService(ctx context.Context) *HelloMethodService {
	return &HelloMethodService{ctx: ctx}
}

// Run create note info
func (s *HelloMethodService) Run(request *example.HelloReq) (resp *example.HelloResp, err error) {
	// Finish your business logic.
	resp = new(example.HelloResp)
	resp.RespBody = fmt.Sprintf("[KITEX] hello, %s", request.Name)
	klog.Infof("[KITEX] hello, %s", request.Name)
	return
}
