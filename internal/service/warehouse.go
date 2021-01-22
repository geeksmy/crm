package service

import (
	"context"

	"crm/gen/warehouse"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewWarehouseService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *WarehouseService {
	return &WarehouseService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type WarehouseService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv WarehouseService) Get(pk string) (*model.Warehouse, error) {
	return nil, nil
}

func (srv WarehouseService) List(limit, cursor int) (*model.Warehouse, *libsgorm.Page, error) {
	return nil, nil, nil
}

func (srv WarehouseService) Create(p *warehouse.CreatePayload) (*model.Warehouse, error) {
	return nil, nil
}

func (srv WarehouseService) Update(p *warehouse.UpdatePayload) (*model.Warehouse, error) {
	return nil, nil
}

func (srv WarehouseService) Delete(pks []string) error {
	return nil
}
