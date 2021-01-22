package service

import (
	"context"

	"crm/gen/supplier"
	"crm/internal/dao"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewSupplierService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *SupplierService {
	return &SupplierService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type SupplierService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv SupplierService) Get(pk string) (*model.Supplier, error) {
	supplierDao := dao.NewSupplierDao(srv.ctx, srv.db)
	scopes := dao.NewSupplierScope()
	scopes = scopes.WithPk(pk)

	res, err := supplierDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取供应商失败", zap.Error(err))
		return nil, err
	}

	supplierModel, err := srv.getUserByFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取创建人和负责人失败", zap.Error(err))
		return nil, err
	}

	return supplierModel, nil
}

func (srv SupplierService) List(limit, cursor int) ([]*model.Supplier, *libsgorm.Page, error) {
	var suppliers []*model.Supplier
	supplierDao := dao.NewSupplierDao(srv.ctx, srv.db)
	scopes := dao.NewSupplierScope()

	res, page, err := supplierDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取供应商列表失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(res); i++ {
		supplierModel, err := srv.getUserByFounderAndHead(res[i])

		if err != nil {
			srv.logger.Error("获取创建人和负责人失败", zap.Error(err))
			return nil, nil, err
		}

		suppliers = append(suppliers, supplierModel)
	}
	return suppliers, page, nil
}

func (srv SupplierService) Create(p *supplier.CreatePayload) (*model.Supplier, error) {
	supplierModel := model.Supplier{
		Name:           p.Name,
		Level:          p.Level,
		ContactName:    p.ContactName,
		ContactPhone:   p.ContactPhone,
		ContactAddress: p.ContactAddress,
		Note:           p.Note,
		HeadID:         p.HeadID,
		FounderID:      p.FounderID,
	}
	supplierDao := dao.NewSupplierDao(srv.ctx, srv.db)

	res, err := supplierDao.Create(supplierModel)

	if err != nil {
		srv.logger.Error("创建供应商失败", zap.Error(err))
		return nil, err
	}

	resModel, err := srv.getUserByFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取创建人和负责人失败", zap.Error(err))
		return nil, err
	}

	return resModel, nil
}

func (srv SupplierService) Update(p *supplier.UpdatePayload) (*model.Supplier, error) {
	supplierDao := dao.NewSupplierDao(srv.ctx, srv.db)
	scopes := dao.NewSupplierScope()
	scopes = scopes.WithPk(p.ID)
	fields := make(map[string]interface{})

	res, err := supplierDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新供应商失败", zap.Error(err))
		return nil, err
	}

	resModel, err := srv.getUserByFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取创建人和负责人失败", zap.Error(err))
		return nil, err
	}

	return resModel, nil
}

func (srv SupplierService) Delete(pks []string) error {
	supplierDao := dao.NewSupplierDao(srv.ctx, srv.db)

	err := supplierDao.Delete(pks)

	return err
}

func (srv SupplierService) getUserByFounderAndHead(supplier *model.Supplier) (*model.Supplier, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)

	founder, err := userDao.GetByID(supplier.FounderID)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	head, err := userDao.GetByID(supplier.HeadID)

	if err != nil {
		srv.logger.Error("获取负责人失败", zap.Error(err))
		return nil, err
	}

	supplier.Founder.ID = founder.ID
	supplier.Founder.Name = founder.Name
	supplier.Head.ID = head.ID
	supplier.Head.Name = head.Name

	return supplier, nil
}
