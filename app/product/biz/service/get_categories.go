package service

import (
	"context"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type GetCategoriesService struct {
	ctx context.Context
} // NewGetCategoriesService new GetCategoriesService
func NewGetCategoriesService(ctx context.Context) *GetCategoriesService {
	return &GetCategoriesService{ctx: ctx}
}

// Run create note info
func (s *GetCategoriesService) Run(req *product.GetCategoriesReq) (resp *product.GetCategoriesResp, err error) {
	// Finish your business logic.

	return
}
