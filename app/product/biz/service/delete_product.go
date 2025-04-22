package service

import (
	"context"
	"errors"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run delete product
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// 参数验证
	if req.Id <= 0 {
		return nil, errors.New("invalid product id")
	}

	// 从数据库删除商品
	if err = mysql.DeleteProduct(mysql.DB, req.Id); err != nil {
		return nil, err
	}

	// 返回响应
	resp = &product.DeleteProductResp{
		Success: true,
	}

	return resp, nil
}
