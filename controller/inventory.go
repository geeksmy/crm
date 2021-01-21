package handler

import (
	"context"

	"crm/gen/inventory"
	"crm/gen/log"
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
	return
}

// 获取库存列表
func (s *inventorysrvc) List(ctx context.Context, p *inventory.ListPayload) (res *inventory.ListResult, err error) {
	res = &inventory.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.List")
	return
}

// 更新库存
func (s *inventorysrvc) Update(ctx context.Context, p *inventory.UpdatePayload) (res *inventory.Inventory, err error) {
	res = &inventory.Inventory{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Update")
	return
}

// 创建库存
func (s *inventorysrvc) Create(ctx context.Context, p *inventory.CreatePayload) (res *inventory.Inventory, err error) {
	res = &inventory.Inventory{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Create")
	return
}

// 删除库存
func (s *inventorysrvc) Delete(ctx context.Context, p *inventory.DeletePayload) (res *inventory.SuccessResult, err error) {
	res = &inventory.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("inventory.Delete")
	return
}
