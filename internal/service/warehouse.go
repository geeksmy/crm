package service

import (
	"context"

	"crm/gen/warehouse"
	"crm/internal/dao"
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
	warehouseDao := dao.NewWarehouseDao(srv.ctx, srv.db)
	scopes := dao.NewWarehouseScope()
	scopes = scopes.WithPk(pk)

	res, err := warehouseDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取仓库失败", zap.Error(err))
		return nil, err
	}

	warehouseModel, err := srv.getByFounder(res)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}
	return warehouseModel, nil
}

func (srv WarehouseService) List(limit, cursor int) ([]*model.Warehouse, *libsgorm.Page, error) {
	var warehousesModel []*model.Warehouse
	warehouseDao := dao.NewWarehouseDao(srv.ctx, srv.db)
	scopes := dao.NewWarehouseScope()

	res, page, err := warehouseDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取仓库列表失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(res); i++ {
		warehouseModel, err := srv.getByFounder(res[i])

		if err != nil {
			srv.logger.Error("获取创建人失败", zap.Error(err))
			return nil, nil, err
		}
		warehousesModel = append(warehousesModel, warehouseModel)
	}
	return warehousesModel, page, nil
}

func (srv WarehouseService) Create(p *warehouse.CreatePayload) (*model.Warehouse, error) {
	warehouseDao := dao.NewWarehouseDao(srv.ctx, srv.db)

	res := &model.Warehouse{
		Name:      p.Name,
		Code:      p.Code,
		Address:   p.Address,
		Type:      p.Type,
		FounderID: p.FounderID,
	}

	res, err := warehouseDao.Create(*res)

	if err != nil {
		srv.logger.Error("创建仓库失败", zap.Error(err))
		return nil, err
	}

	warehouseModel, err := srv.getByFounder(res)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	return warehouseModel, nil
}

func (srv WarehouseService) Update(p *warehouse.UpdatePayload) (*model.Warehouse, error) {
	warehouseDao := dao.NewWarehouseDao(srv.ctx, srv.db)
	scopes := dao.NewWarehouseScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})

	res, err := warehouseDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新仓库失败", zap.Error(err))
		return nil, err
	}

	warehouseModel, err := srv.getByFounder(res)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	return warehouseModel, nil
}

func (srv WarehouseService) Delete(pks []string) error {
	warehouseDao := dao.NewWarehouseDao(srv.ctx, srv.db)

	err := warehouseDao.Delete(pks)

	return err
}

func (srv WarehouseService) getByFounder(warehouseModel *model.Warehouse) (*model.Warehouse, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)

	founder, err := userDao.GetByID(warehouseModel.FounderID)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	warehouseModel.Founder.ID = founder.ID
	warehouseModel.Founder.Name = founder.Name

	return warehouseModel, nil
}
