package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/product"
)

// Product service example implementation.
// The example methods log the requests and return zero values.
type productsrvc struct {
	authsrvc
	logger *log.Logger
}

// NewProduct returns the Product service implementation.
func NewProduct(logger *log.Logger) product.Service {
	return &productsrvc{authsrvc{logger: logger}, logger}
}

// 获取单个产品
func (s *productsrvc) Get(ctx context.Context, p *product.GetPayload) (res *product.Product, err error) {
	res = &product.Product{}
	logger := L(ctx, s.logger)
	logger.Info("product.Get")
	return
}

// 获取产品列表
func (s *productsrvc) List(ctx context.Context, p *product.ListPayload) (res *product.ListResult, err error) {
	res = &product.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("product.List")
	return
}

// 更新产品
func (s *productsrvc) Update(ctx context.Context, p *product.UpdatePayload) (res *product.Product, err error) {
	res = &product.Product{}
	logger := L(ctx, s.logger)
	logger.Info("product.Update")
	return
}

// 创建产品
func (s *productsrvc) Create(ctx context.Context, p *product.CreatePayload) (res *product.Product, err error) {
	res = &product.Product{}
	logger := L(ctx, s.logger)
	logger.Info("product.Create")
	return
}

// 删除产品
func (s *productsrvc) Delete(ctx context.Context, p *product.DeletePayload) (res *product.SuccessResult, err error) {
	res = &product.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("product.Delete")
	return
}
