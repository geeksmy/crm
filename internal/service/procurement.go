package service

import (
	"context"

	"crm/gen/procurement"
	"crm/internal/dao"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewProcurementService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *ProcurementService {
	return &ProcurementService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type ProcurementService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv ProcurementService) Get(pk string) (*model.Procurement, error) {
	procurementDao := dao.NewProcurementDao(srv.ctx, srv.db)
	scopes := dao.NewProcurementScope()
	scopes = scopes.WithPk(pk)

	res, err := procurementDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取单个采购失败", zap.Error(err))
		return nil, err
	}

	procurementModel, err := srv.getByFounderAndHeadAndSupplier(res)

	if err != nil {
		srv.logger.Error("获取创建人和负责人和供应商失败", zap.Error(err))
		return nil, err
	}

	return procurementModel, nil
}

func (srv ProcurementService) List(limit, cursor int) ([]*model.Procurement, *libsgorm.Page, error) {
	var procurementsModel []*model.Procurement
	procurementDao := dao.NewProcurementDao(srv.ctx, srv.db)
	scopes := dao.NewProcurementScope()

	res, page, err := procurementDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取采购列表失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(res); i++ {
		procurementModel, err := srv.getByFounderAndHeadAndSupplier(res[i])

		if err != nil {
			srv.logger.Error("获取创建人和负责人和供应商失败", zap.Error(err))
			return nil, nil, err
		}

		procurementsModel = append(procurementsModel, procurementModel)
	}

	return procurementsModel, page, nil
}

func (srv ProcurementService) Create(p *procurement.CreatePayload) (*model.Procurement, error) {
	procurementDao := dao.NewProcurementDao(srv.ctx, srv.db)

	procurementModel := model.Procurement{
		SupplierID:    p.SupplierID,
		Code:          p.Code,
		Note:          p.Note,
		Money:         p.Money,
		IsSalesReturn: p.IsSalesReturn,
		HeadID:        p.HeadID,
		FounderID:     p.FounderID,
	}

	res, err := procurementDao.Create(procurementModel)

	if err != nil {
		srv.logger.Error("创建采购失败", zap.Error(err))
		return nil, err
	}

	resModel, err := srv.getByFounderAndHeadAndSupplier(res)

	if err != nil {
		srv.logger.Error("获取创建人和负责人和供应商失败", zap.Error(err))
		return nil, err
	}

	return resModel, nil
}

func (srv ProcurementService) Update(p *procurement.UpdatePayload) (*model.Procurement, error) {
	procurementDao := dao.NewProcurementDao(srv.ctx, srv.db)
	scopes := dao.NewProcurementScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})

	res, err := procurementDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新采购失败", zap.Error(err))
		return nil, err
	}

	procurementModel, err := srv.getByFounderAndHeadAndSupplier(res)

	if err != nil {
		srv.logger.Error("获取创建人和负责人和供应商失败", zap.Error(err))
		return nil, err
	}

	return procurementModel, nil
}

func (srv ProcurementService) Delete(pks []string) error {
	procurementDao := dao.NewProcurementDao(srv.ctx, srv.db)

	err := procurementDao.Delete(pks)

	return err
}

func (srv ProcurementService) getByFounderAndHeadAndSupplier(procurementModel *model.Procurement) (*model.Procurement, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	supplierDao := dao.NewSupplierDao(srv.ctx, srv.db)

	founder, err := userDao.GetByID(procurementModel.FounderID)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	head, err := userDao.GetByID(procurementModel.HeadID)

	if err != nil {
		srv.logger.Error("获取负责人失败", zap.Error(err))
		return nil, err
	}

	supplier, err := supplierDao.GetByID(procurementModel.SupplierID)

	if err != nil {
		srv.logger.Error("获取供应商失败", zap.Error(err))
		return nil, err
	}

	procurementModel.Founder.ID = founder.ID
	procurementModel.Founder.Name = founder.Name
	procurementModel.Head.ID = head.ID
	procurementModel.Head.Name = head.Name
	procurementModel.Supplier.ID = supplier.ID
	procurementModel.Supplier.Name = supplier.Name

	return procurementModel, nil
}
