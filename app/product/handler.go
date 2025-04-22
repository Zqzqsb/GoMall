package main

import (
	"context"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
	"zqzqsb/gomall/app/product/biz/service"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp, err = service.NewCreateProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// ListProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetCategories implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetCategories(ctx context.Context, req *product.GetCategoriesReq) (resp *product.GetCategoriesResp, err error) {
	resp, err = service.NewGetCategoriesService(ctx).Run(req)

	return resp, err
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateStock(ctx context.Context, req *product.UpdateStockReq) (resp *product.UpdateStockResp, err error) {
	resp, err = service.NewUpdateStockService(ctx).Run(req)

	return resp, err
}
