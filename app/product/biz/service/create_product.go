package service

import (
	"context"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	"zqzqsb/gomall/app/product/biz/model"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create product info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// 创建商品模型
	p := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageUrl,
		Category:    req.Category,
		IsOnSale:    req.IsOnSale,
	}

	// 设置图片集
	if len(req.Gallery) > 0 {
		if err = p.SetGallery(req.Gallery); err != nil {
			return nil, err
		}
	}

	// 设置属性
	if len(req.Attributes) > 0 {
		if err = p.SetAttributes(req.Attributes); err != nil {
			return nil, err
		}
	}

	// 保存到数据库
	productID, err := mysql.CreateProduct(mysql.DB, p)
	if err != nil {
		return nil, err
	}

	// 返回响应
	resp = &product.CreateProductResp{
		ProductId: productID,
	}

	return resp, nil
}
