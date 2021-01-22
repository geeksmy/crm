package service

import (
	"context"
	"time"

	"crm/gen/inventory"
	"crm/internal/dao"
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
	inventoryDao := dao.NewInventoryDao(srv.ctx, srv.db)
	scopes := dao.NewInventoryScope()
	scopes = scopes.WithPk(pk)

	res, err := inventoryDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取库存失败", zap.Error(err))
		return nil, err
	}

	inventoryModel, err := srv.getByProductAndWarehouseAndFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取附加返回值失败", zap.Error(err))
		return nil, err
	}

	return inventoryModel, nil
}

func (srv InventoryService) List(limit, cursor int) ([]*model.Inventory, *libsgorm.Page, error) {
	var inventorysModel []*model.Inventory
	inventoryDao := dao.NewInventoryDao(srv.ctx, srv.db)
	scopes := dao.NewInventoryScope()

	res, page, err := inventoryDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取库存列表失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(res); i++ {
		inventoryModel, err := srv.getByProductAndWarehouseAndFounderAndHead(res[i])

		if err != nil {
			srv.logger.Error("获取附加返回值失败", zap.Error(err))
			return nil, nil, err
		}

		inventorysModel = append(inventorysModel, inventoryModel)
	}

	return inventorysModel, page, nil
}

func (srv InventoryService) Create(p *inventory.CreatePayload) (*model.Inventory, error) {
	inventoryDao := dao.NewInventoryDao(srv.ctx, srv.db)

	inventoryDate, err := time.Parse("2006-01-02 15:04:05", p.InventoryDate)

	if err != nil {
		srv.logger.Error("string转time失败", zap.Error(err))
		return nil, err
	}

	inventoryModel := &model.Inventory{
		ProductID:     p.ProductID,
		Number:        p.Number,
		Code:          p.Code,
		WarehouseID:   p.WarehouseID,
		Type:          p.Type,
		InventoryDate: inventoryDate,
		Note:          p.Note,
		HeadID:        p.HeadID,
		FounderID:     p.FounderID,
		InAndOut:      p.InAndOut,
	}

	res, err := inventoryDao.Create(*inventoryModel)

	if err != nil {
		srv.logger.Error("创建库存失败", zap.Error(err))
		return nil, err
	}

	inventoryModel, err = srv.getByProductAndWarehouseAndFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取附加返回值失败", zap.Error(err))
		return nil, err
	}

	return inventoryModel, nil
}

func (srv InventoryService) Update(p *inventory.UpdatePayload) (*model.Inventory, error) {
	inventoryDao := dao.NewInventoryDao(srv.ctx, srv.db)
	scopes := dao.NewInventoryScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})

	res, err := inventoryDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新库存失败", zap.Error(err))
		return nil, err
	}

	inventoryModel, err := srv.getByProductAndWarehouseAndFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取附加返回值失败", zap.Error(err))
		return nil, err
	}

	return inventoryModel, nil
}

func (srv InventoryService) Delete(pks []string) error {
	inventoryDao := dao.NewInventoryDao(srv.ctx, srv.db)

	err := inventoryDao.Delete(pks)

	return err
}

func (srv InventoryService) getByProductAndWarehouseAndFounderAndHead(inventoryModel *model.Inventory) (*model.Inventory, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	productDao := dao.NewProductDao(srv.ctx, srv.db)
	warehouseDao := dao.NewWarehouseDao(srv.ctx, srv.db)

	head, err := userDao.GetByID(inventoryModel.HeadID)

	if err != nil {
		srv.logger.Error("获取负责人失败", zap.Error(err))
		return nil, err
	}

	founder, err := userDao.GetByID(inventoryModel.FounderID)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	product, err := productDao.GetByID(inventoryModel.ProductID)

	if err != nil {
		srv.logger.Error("获取产品失败", zap.Error(err))
		return nil, err
	}

	warehouse, err := warehouseDao.GetByID(inventoryModel.WarehouseID)

	if err != nil {
		srv.logger.Error("获取仓库失败", zap.Error(err))
		return nil, err
	}

	inventoryModel.Head.ID = head.ID
	inventoryModel.Head.Name = head.Name
	inventoryModel.Founder.ID = founder.ID
	inventoryModel.Founder.Name = founder.Name
	inventoryModel.Product = *product
	inventoryModel.Warehouse = *warehouse

	return inventoryModel, nil
}
