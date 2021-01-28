package handler

import (
	"context"

	"crm/gen/inventory"
	"crm/gen/log"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
)

// Inventory service example implementation.
// The example methods log the requests and return zero values.
type inventorysrvc struct {
	authsrvc
	logger *log.Logger
}

// NewInventory returns the Inventory service implementation.
func NewInventory(logger *log.Logger) inventory.Service {
	return &inventorysrvc{authsrvc{logger: logger}, logger}
}

// 获取单个库存
func (s *inventorysrvc) Get(ctx context.Context, p *inventory.GetPayload) (res *inventory.Inventory, err error) {
	res = &inventory.Inventory{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Get")

	inventoryService := service.NewInventoryService(ctx, dao.DB, logger)
	inventoryModel, errMsg := inventoryService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个库存失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelInventory2Inventory(inventoryModel)
	return res, nil
}

// 获取库存列表
func (s *inventorysrvc) List(ctx context.Context, p *inventory.ListPayload) (res *inventory.ListResult, err error) {
	res = &inventory.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.List")

	inventoryService := service.NewInventoryService(ctx, dao.DB, logger)
	inventorysModel, page, errMsg := inventoryService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取库存列表失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelInventorys2Inventorys(inventorysModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新库存
func (s *inventorysrvc) Update(ctx context.Context, p *inventory.UpdatePayload) (res *inventory.Inventory, err error) {
	res = &inventory.Inventory{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Update")

	inventoryService := service.NewInventoryService(ctx, dao.DB, logger)
	inventoryModel, errMsg := inventoryService.Update(p)

	if errMsg != nil {
		logger.Error("更新库存失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelInventory2Inventory(inventoryModel)
	return res, nil
}

// 创建库存
func (s *inventorysrvc) Create(ctx context.Context, p *inventory.CreatePayload) (res *inventory.Inventory, err error) {
	res = &inventory.Inventory{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Create")

	inventoryService := service.NewInventoryService(ctx, dao.DB, logger)
	inventoryModel, errMsg := inventoryService.Create(p)

	if errMsg != nil {
		logger.Error("创建库存失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelInventory2Inventory(inventoryModel)
	return res, nil
}

// 删除库存
func (s *inventorysrvc) Delete(ctx context.Context, p *inventory.DeletePayload) (res *inventory.SuccessResult, err error) {
	res = &inventory.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Delete")

	inventoryService := service.NewInventoryService(ctx, dao.DB, logger)
	errMsg := inventoryService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("删除库存失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
