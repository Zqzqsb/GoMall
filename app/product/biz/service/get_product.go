package service

import (
	"context"
	"errors"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run get product info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// 参数验证
	if req.Id <= 0 {
		return nil, errors.New("invalid product id")
	}

	// 从数据库获取商品
	p, err := mysql.GetProductByID(mysql.DB, req.Id)
	if err != nil {
		return nil, err
	}

	// 构建响应
	gallery := p.GetGallery()
	attributes := p.GetAttributes()

	resp = &product.GetProductResp{
		Product: &product.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			ImageUrl:    p.ImageURL,
			Gallery:     gallery,
			Category:    p.Category,
			IsOnSale:    p.IsOnSale,
			Attributes:  attributes,
			Rating:      p.Rating,
			SalesCount:  p.SalesCount,
			CreateTime:  p.CreatedAt.Unix(),
			UpdateTime:  p.UpdatedAt.Unix(),
		},
	}

	return resp, nil
}
