package service

import (
	"context"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type UpdateStockService struct {
	ctx context.Context
} // NewUpdateStockService new UpdateStockService
func NewUpdateStockService(ctx context.Context) *UpdateStockService {
	return &UpdateStockService{ctx: ctx}
}

// Run create note info
func (s *UpdateStockService) Run(req *product.UpdateStockReq) (resp *product.UpdateStockResp, err error) {
	// Finish your business logic.

	return
}
