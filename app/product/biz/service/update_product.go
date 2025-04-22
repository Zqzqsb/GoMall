package service

import (
	"context"
	"errors"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run update product info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// 参数验证
	if req.Id <= 0 {
		return nil, errors.New("invalid product id")
	}

	// 从数据库获取商品
	p, err := mysql.GetProductByID(mysql.DB, req.Id)
	if err != nil {
		return nil, err
	}

	// 更新商品信息
	p.Name = req.Name
	p.Description = req.Description
	p.Price = req.Price
	p.Stock = req.Stock
	p.ImageURL = req.ImageUrl
	p.Category = req.Category
	p.IsOnSale = req.IsOnSale

	// 更新图片集
	if len(req.Gallery) > 0 {
		if err = p.SetGallery(req.Gallery); err != nil {
			return nil, err
		}
	}

	// 更新属性
	if len(req.Attributes) > 0 {
		if err = p.SetAttributes(req.Attributes); err != nil {
			return nil, err
		}
	}

	// 保存到数据库
	if err = mysql.UpdateProduct(mysql.DB, p); err != nil {
		return nil, err
	}

	// 返回响应
	resp = &product.UpdateProductResp{
		Success: true,
	}

	return resp, nil
}
