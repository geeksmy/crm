package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/procurement"
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
	return
}

// 获取采购列表
func (s *procurementsrvc) List(ctx context.Context, p *procurement.ListPayload) (res *procurement.ListResult, err error) {
	res = &procurement.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.List")
	return
}

// 更新采购
func (s *procurementsrvc) Update(ctx context.Context, p *procurement.UpdatePayload) (res *procurement.Procurement, err error) {
	res = &procurement.Procurement{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Update")
	return
}

// 创建采购
func (s *procurementsrvc) Create(ctx context.Context, p *procurement.CreatePayload) (res *procurement.Procurement, err error) {
	res = &procurement.Procurement{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Create")
	return
}

// 删除采购
func (s *procurementsrvc) Delete(ctx context.Context, p *procurement.DeletePayload) (res *procurement.SuccessResult, err error) {
	res = &procurement.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("procurement.Delete")
	return
}
