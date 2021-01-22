package service

import (
	"context"

	"crm/gen/customer"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewCustomerService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *CustomerService {
	return &CustomerService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type CustomerService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv CustomerService) Get(pk string) (*model.Customer, error) {
	return nil, nil
}

func (srv CustomerService) List(limit, cursor int) (*model.Customer, *libsgorm.Page, error) {
	return nil, nil, nil
}

func (srv CustomerService) Create(p *customer.CreatePayload) (*model.Customer, error) {
	return nil, nil
}

func (srv CustomerService) Update(p *customer.UpdatePayload) (*model.Customer, error) {
	return nil, nil
}

func (srv CustomerService) Delete(pks []string) error {
	return nil
}
