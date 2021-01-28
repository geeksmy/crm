package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/warehouse"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
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

	warehouseService := service.NewWarehouseService(ctx, dao.DB, logger)
	warehouseModel, errMsg := warehouseService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个仓库失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelWarehouse2Warehouse(warehouseModel)
	return res, nil
}

// 获取仓库列表
func (s *warehousesrvc) List(ctx context.Context, p *warehouse.ListPayload) (res *warehouse.ListResult, err error) {
	res = &warehouse.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.List")

	warehouseService := service.NewWarehouseService(ctx, dao.DB, logger)
	warehousesModel, page, errMsg := warehouseService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取单个仓库失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelWarehouses2Warehouses(warehousesModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新仓库
func (s *warehousesrvc) Update(ctx context.Context, p *warehouse.UpdatePayload) (res *warehouse.Warehouse, err error) {
	res = &warehouse.Warehouse{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Update")

	warehouseService := service.NewWarehouseService(ctx, dao.DB, logger)
	warehouseModel, errMsg := warehouseService.Update(p)

	if errMsg != nil {
		logger.Error("更新仓库失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelWarehouse2Warehouse(warehouseModel)
	return res, nil
}

// 创建仓库
func (s *warehousesrvc) Create(ctx context.Context, p *warehouse.CreatePayload) (res *warehouse.Warehouse, err error) {
	res = &warehouse.Warehouse{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Create")

	warehouseService := service.NewWarehouseService(ctx, dao.DB, logger)
	warehouseModel, errMsg := warehouseService.Create(p)

	if errMsg != nil {
		logger.Error("创建仓库失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelWarehouse2Warehouse(warehouseModel)
	return res, nil
}

// 删除仓库
func (s *warehousesrvc) Delete(ctx context.Context, p *warehouse.DeletePayload) (res *warehouse.SuccessResult, err error) {
	res = &warehouse.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("warehouse.Delete")

	warehouseService := service.NewWarehouseService(ctx, dao.DB, logger)
	errMsg := warehouseService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("删除仓库失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
