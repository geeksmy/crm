package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/supplier"
)

// Supplier service example implementation.
// The example methods log the requests and return zero values.
type suppliersrvc struct {
	authsrvc
	logger *log.Logger
}

// NewSupplier returns the Supplier service implementation.
func NewSupplier(logger *log.Logger) supplier.Service {
	return &suppliersrvc{authsrvc{logger: logger}, logger}
}

// 获取单个供应商
func (s *suppliersrvc) Get(ctx context.Context, p *supplier.GetPayload) (res *supplier.Supplier, err error) {
	res = &supplier.Supplier{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Get")
	return
}

// 获取供应商列表
func (s *suppliersrvc) List(ctx context.Context, p *supplier.ListPayload) (res *supplier.ListResult, err error) {
	res = &supplier.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.List")
	return
}

// 更新供应商
func (s *suppliersrvc) Update(ctx context.Context, p *supplier.UpdatePayload) (res *supplier.Supplier, err error) {
	res = &supplier.Supplier{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Update")
	return
}

// 创建供应商
func (s *suppliersrvc) Create(ctx context.Context, p *supplier.CreatePayload) (res *supplier.Supplier, err error) {
	res = &supplier.Supplier{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Create")
	return
}

// 删除供应商
func (s *suppliersrvc) Delete(ctx context.Context, p *supplier.DeletePayload) (res *supplier.SuccessResult, err error) {
	res = &supplier.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Delete")
	return
}
