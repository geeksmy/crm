package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/sales"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
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

	salesService := service.NewSalesService(ctx, dao.DB, logger)
	salesModel, errMsg := salesService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个销售失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelSales2Sales(salesModel)
	return res, nil
}

// 获取销售列表
func (s *salessrvc) List(ctx context.Context, p *sales.ListPayload) (res *sales.ListResult, err error) {
	res = &sales.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("sales.List")

	salesService := service.NewSalesService(ctx, dao.DB, logger)
	salessModel, page, errMsg := salesService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取单个销售失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelSaless2Saless(salessModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新销售
func (s *salessrvc) Update(ctx context.Context, p *sales.UpdatePayload) (res *sales.Sales, err error) {
	res = &sales.Sales{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Update")

	salesService := service.NewSalesService(ctx, dao.DB, logger)
	salesModel, errMsg := salesService.Update(p)

	if errMsg != nil {
		logger.Error("更新销售失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelSales2Sales(salesModel)
	return res, nil
}

// 创建销售
func (s *salessrvc) Create(ctx context.Context, p *sales.CreatePayload) (res *sales.Sales, err error) {
	res = &sales.Sales{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Create")

	salesService := service.NewSalesService(ctx, dao.DB, logger)
	salesModel, errMsg := salesService.Create(p)

	if errMsg != nil {
		logger.Error("创建销售失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelSales2Sales(salesModel)
	return res, nil
}

// 删除销售
func (s *salessrvc) Delete(ctx context.Context, p *sales.DeletePayload) (res *sales.SuccessResult, err error) {
	res = &sales.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("sales.Delete")

	salesService := service.NewSalesService(ctx, dao.DB, logger)
	errMsg := salesService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("获取单个销售失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
