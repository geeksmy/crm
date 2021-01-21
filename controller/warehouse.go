package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/warehouse"
)

// Warehouse service example implementation.
// The example methods log the requests and return zero values.
type warehousesrvc struct {
	authsrvc
	logger *log.Logger
}

// NewWarehouse returns the Warehouse service implementation.
func NewWarehouse(logger *log.Logger) warehouse.Service {
	return &warehousesrvc{authsrvc{logger: logger}, logger}
}

// 获取单个仓库
func (s *warehousesrvc) Get(ctx context.Context, p *warehouse.GetPayload) (res *warehouse.Warehouse, err error) {
	res = &warehouse.Warehouse{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Get")
	return
}

// 获取仓库列表
func (s *warehousesrvc) List(ctx context.Context, p *warehouse.ListPayload) (res *warehouse.ListResult, err error) {
	res = &warehouse.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.List")
	return
}

// 更新仓库
func (s *warehousesrvc) Update(ctx context.Context, p *warehouse.UpdatePayload) (res *warehouse.Warehouse, err error) {
	res = &warehouse.Warehouse{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Update")
	return
}

// 创建仓库
func (s *warehousesrvc) Create(ctx context.Context, p *warehouse.CreatePayload) (res *warehouse.Warehouse, err error) {
	res = &warehouse.Warehouse{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Create")
	return
}

// 删除仓库
func (s *warehousesrvc) Delete(ctx context.Context, p *warehouse.DeletePayload) (res *warehouse.SuccessResult, err error) {
	res = &warehouse.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Delete")
	return
}
