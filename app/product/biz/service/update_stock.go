package service

import (
	"context"
	"errors"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type UpdateStockService struct {
	ctx context.Context
} // NewUpdateStockService new UpdateStockService
func NewUpdateStockService(ctx context.Context) *UpdateStockService {
	return &UpdateStockService{ctx: ctx}
}

// Run update product stock
func (s *UpdateStockService) Run(req *product.UpdateStockReq) (resp *product.UpdateStockResp, err error) {
	// 参数验证
	if req.ProductId <= 0 {
		return nil, errors.New("invalid product id")
	}

	// 更新商品库存
	currentStock, err := mysql.UpdateStock(mysql.DB, req.ProductId, req.Quantity)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp = &product.UpdateStockResp{
		Success:      true,
		CurrentStock: currentStock,
	}

	return resp, nil
}
