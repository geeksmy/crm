package service

import (
	"context"

	"crm/gen/inventory"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewInventoryService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *InventoryService {
	return &InventoryService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type InventoryService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv InventoryService) Get(pk string) (*model.Inventory, error) {
	return nil, nil
}

func (srv InventoryService) List(limit, cursor int) (*model.Inventory, *libsgorm.Page, error) {
	return nil, nil, nil
}

func (srv InventoryService) Create(p *inventory.CreatePayload) (*model.Inventory, error) {
	return nil, nil
}

func (srv InventoryService) Update(p *inventory.UpdatePayload) (*model.Inventory, error) {
	return nil, nil
}

func (srv InventoryService) Delete(pks []string) error {
	return nil
}
