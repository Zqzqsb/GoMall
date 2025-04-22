package service

import (
	"context"
	
	"zqzqsb/gomall/app/product/biz/dal/mysql"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

type GetCategoriesService struct {
	ctx context.Context
} // NewGetCategoriesService new GetCategoriesService
func NewGetCategoriesService(ctx context.Context) *GetCategoriesService {
	return &GetCategoriesService{ctx: ctx}
}

// Run get categories
func (s *GetCategoriesService) Run(req *product.GetCategoriesReq) (resp *product.GetCategoriesResp, err error) {
	// 从数据库获取所有商品分类
	categories, err := mysql.GetCategories(mysql.DB)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp = &product.GetCategoriesResp{
		Categories: categories,
	}

	return resp, nil
}
