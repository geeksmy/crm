package service

import (
	"context"

	"crm/gen/customer"
	"crm/internal/dao"
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
	customerDao := dao.NewCustomerDao(srv.ctx, srv.db)
	scopes := dao.NewCustomerScope()
	scopes = scopes.WithPk(pk)

	res, err := customerDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取客户失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv CustomerService) List(limit, cursor int) ([]*model.Customer, *libsgorm.Page, error) {
	customerDao := dao.NewCustomerDao(srv.ctx, srv.db)
	scopes := dao.NewCustomerScope()

	res, page, err := customerDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取客户列表失败", zap.Error(err))
		return nil, nil, err
	}

	return res, page, nil
}

func (srv CustomerService) Create(p *customer.CreatePayload) (*model.Customer, error) {
	customerDao := dao.NewCustomerDao(srv.ctx, srv.db)

	customerModel := model.Customer{
		Name:     p.Name,
		Src:      p.Src,
		Mobile:   p.Mobile,
		Url:      p.URL,
		Email:    p.Email,
		Industry: p.Industry,
		Level:    p.Level,
		Note:     p.Note,
		Address:  p.Address,
	}

	res, err := customerDao.Create(customerModel)

	if err != nil {
		srv.logger.Error("创建客户失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv CustomerService) Update(p *customer.UpdatePayload) (*model.Customer, error) {
	customerDao := dao.NewCustomerDao(srv.ctx, srv.db)
	scopes := dao.NewCustomerScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})

	res, err := customerDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新客户失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv CustomerService) Delete(pks []string) error {
	customerDao := dao.NewCustomerDao(srv.ctx, srv.db)

	err := customerDao.Delete(pks)

	return err
}
