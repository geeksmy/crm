package service

import (
	"context"

	"crm/gen/sales"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewSalesService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *SalesService {
	return &SalesService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type SalesService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv SalesService) Get(pk string) (*model.Sales, error) {
	return nil, nil
}

func (srv SalesService) List(limit, cursor int) (*model.Sales, *libsgorm.Page, error) {
	return nil, nil, nil
}

func (srv SalesService) Create(p *sales.CreatePayload) (*model.Sales, error) {
	return nil, nil
}

func (srv SalesService) Update(p *sales.UpdatePayload) (*model.Sales, error) {
	return nil, nil
}

func (srv SalesService) Delete(pks []string) error {
	return nil
}
