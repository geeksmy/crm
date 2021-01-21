package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/sales"
)

// Sales service example implementation.
// The example methods log the requests and return zero values.
type salessrvc struct {
	authsrvc
	logger *log.Logger
}

// NewSales returns the Sales service implementation.
func NewSales(logger *log.Logger) sales.Service {
	return &salessrvc{authsrvc{logger: logger}, logger}
}

// 获取单个销售
func (s *salessrvc) Get(ctx context.Context, p *sales.GetPayload) (res *sales.Sales, err error) {
	res = &sales.Sales{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Get")
	return
}

// 获取销售列表
func (s *salessrvc) List(ctx context.Context, p *sales.ListPayload) (res *sales.ListResult, err error) {
	res = &sales.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("sales.List")
	return
}

// 更新销售
func (s *salessrvc) Update(ctx context.Context, p *sales.UpdatePayload) (res *sales.Sales, err error) {
	res = &sales.Sales{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Update")
	return
}

// 创建销售
func (s *salessrvc) Create(ctx context.Context, p *sales.CreatePayload) (res *sales.Sales, err error) {
	res = &sales.Sales{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Create")
	return
}

// 删除销售
func (s *salessrvc) Delete(ctx context.Context, p *sales.DeletePayload) (res *sales.SuccessResult, err error) {
	res = &sales.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Delete")
	return
}
