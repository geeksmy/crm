package handler

import (
	"context"

	"crm/gen/customer"
	"crm/gen/log"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
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

	customerService := service.NewCustomerService(ctx, dao.DB, logger)
	customerModel, errMsg := customerService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个客户失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelCustomer2Customer(customerModel)
	return res, nil
}

// 获取客户列表
func (s *customersrvc) List(ctx context.Context, p *customer.ListPayload) (res *customer.ListResult, err error) {
	res = &customer.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("customer.List")

	customerService := service.NewCustomerService(ctx, dao.DB, logger)
	customersModel, page, errMsg := customerService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取客户列表失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelCustomers2Customers(customersModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新客户
func (s *customersrvc) Update(ctx context.Context, p *customer.UpdatePayload) (res *customer.Customer, err error) {
	// res = &customer.Customer{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Update")

	customerService := service.NewCustomerService(ctx, dao.DB, logger)
	customerModel, errMsg := customerService.Update(p)

	if errMsg != nil {
		logger.Error("更新客户失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelCustomer2Customer(customerModel)
	return res, nil
}

// 创建客户
func (s *customersrvc) Create(ctx context.Context, p *customer.CreatePayload) (res *customer.Customer, err error) {
	// res = &customer.Customer{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Create")

	customerService := service.NewCustomerService(ctx, dao.DB, logger)
	customerModel, errMsg := customerService.Create(p)

	if errMsg != nil {
		logger.Error("创建客户失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelCustomer2Customer(customerModel)
	return res, nil
}

// 删除客户
func (s *customersrvc) Delete(ctx context.Context, p *customer.DeletePayload) (res *customer.SuccessResult, err error) {
	res = &customer.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("customer.Delete")

	customerService := service.NewCustomerService(ctx, dao.DB, logger)
	errMsg := customerService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("删除客户失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
