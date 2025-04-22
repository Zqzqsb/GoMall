package service

import (
	"context"
	"testing"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

func TestGetCategories_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetCategoriesService(ctx)
	// init req and assert value

	req := &product.GetCategoriesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
