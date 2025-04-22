package service

import (
	"context"
	"math"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run list products
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 从数据库获取商品列表
	products, total, err := mysql.ListProducts(mysql.DB, req)
	if err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := int32(math.Ceil(float64(total) / float64(req.PageSize)))

	// 构建响应
	resp = &product.ListProductsResp{
		Products:    make([]*product.Product, 0, len(products)),
		Total:       int32(total),
		Page:        req.Page,
		PageSize:    req.PageSize,
		TotalPages:  totalPages,
	}

	// 填充商品信息
	for _, p := range products {
		gallery := p.GetGallery()
		attributes := p.GetAttributes()

		resp.Products = append(resp.Products, &product.Product{
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
		})
	}

	return resp, nil
}
