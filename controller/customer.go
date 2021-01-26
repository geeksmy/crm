package handler

import (
	"context"

	"crm/gen/customer"
	"crm/gen/log"
)

// Customer service example implementation.
// The example methods log the requests and return zero values.
type customersrvc struct {
	authsrvc
	logger *log.Logger
}

// NewCustomer returns the Customer service implementation.
func NewCustomer(logger *log.Logger) customer.Service {
	return &customersrvc{authsrvc{logger: logger}, logger}
}

// 获取单个客户
func (s *customersrvc) Get(ctx context.Context, p *customer.GetPayload) (res *customer.Customer, err error) {
	// res = &customer.Customer{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Get")
	return
}

// 获取客户列表
func (s *customersrvc) List(ctx context.Context, p *customer.ListPayload) (res *customer.ListResult, err error) {
	res = &customer.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("customer.List")
	return
}

// 更新客户
func (s *customersrvc) Update(ctx context.Context, p *customer.UpdatePayload) (res *customer.Customer, err error) {
	// res = &customer.Customer{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Update")
	return
}

// 创建客户
func (s *customersrvc) Create(ctx context.Context, p *customer.CreatePayload) (res *customer.Customer, err error) {
	// res = &customer.Customer{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Create")
	return
}

// 删除客户
func (s *customersrvc) Delete(ctx context.Context, p *customer.DeletePayload) (res *customer.SuccessResult, err error) {
	res = &customer.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Delete")
	return
}
