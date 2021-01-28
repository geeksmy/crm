package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/procurement"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
)

// Procurement service example implementation.
// The example methods log the requests and return zero values.
type procurementsrvc struct {
	authsrvc
	logger *log.Logger
}

// NewProcurement returns the Procurement service implementation.
func NewProcurement(logger *log.Logger) procurement.Service {
	return &procurementsrvc{authsrvc{logger: logger}, logger}
}

// 获取单个采购
func (s *procurementsrvc) Get(ctx context.Context, p *procurement.GetPayload) (res *procurement.Procurement, err error) {
	res = &procurement.Procurement{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Get")

	procurementService := service.NewProcurementService(ctx, dao.DB, logger)
	procurementModel, errMsg := procurementService.Get(p.ID)

	if errMsg != nil {
		logger.Error("删除单个采购失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelProcurement2Procurement(procurementModel)
	return res, nil
}

// 获取采购列表
func (s *procurementsrvc) List(ctx context.Context, p *procurement.ListPayload) (res *procurement.ListResult, err error) {
	res = &procurement.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.List")

	procurementService := service.NewProcurementService(ctx, dao.DB, logger)
	procurementsModel, page, errMsg := procurementService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("删除单个采购失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelProcurements2Procurements(procurementsModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新采购
func (s *procurementsrvc) Update(ctx context.Context, p *procurement.UpdatePayload) (res *procurement.Procurement, err error) {
	res = &procurement.Procurement{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Update")

	procurementService := service.NewProcurementService(ctx, dao.DB, logger)
	procurementModel, errMsg := procurementService.Update(p)

	if errMsg != nil {
		logger.Error("更新采购失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelProcurement2Procurement(procurementModel)
	return res, nil
}

// 创建采购
func (s *procurementsrvc) Create(ctx context.Context, p *procurement.CreatePayload) (res *procurement.Procurement, err error) {
	res = &procurement.Procurement{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Create")

	procurementService := service.NewProcurementService(ctx, dao.DB, logger)
	procurementModel, errMsg := procurementService.Create(p)

	if errMsg != nil {
		logger.Error("创建采购失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelProcurement2Procurement(procurementModel)
	return res, nil
}

// 删除采购
func (s *procurementsrvc) Delete(ctx context.Context, p *procurement.DeletePayload) (res *procurement.SuccessResult, err error) {
	res = &procurement.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Delete")

	procurementService := service.NewProcurementService(ctx, dao.DB, logger)
	errMsg := procurementService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("删除单个采购失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
