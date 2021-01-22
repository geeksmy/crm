package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewWarehouseDao(ctx context.Context, db *gorm.DB) *warehouseDao {
	return &warehouseDao{
		ctx: ctx,
		db:  db,
	}
}

type warehouseDao struct {
	ctx context.Context
	db  *gorm.DB
}

type WarehouseScope struct {
	scopes []libsgorm.Scope
}

func NewWarehouseScope() WarehouseScope {
	return WarehouseScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s WarehouseScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s WarehouseScope) WithPk(id string) WarehouseScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增仓库
func (d warehouseDao) Create(warehouse model.Warehouse) (*model.Warehouse, error) {
	warehouse.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&warehouse).Error

	return &warehouse, err
}

// 仓库列表
func (d warehouseDao) List(limit, cursor int, scope WarehouseScope) ([]*model.Warehouse, *libsgorm.Page, error) {
	var res []*model.Warehouse

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个仓库
func (d warehouseDao) Get(scope WarehouseScope) (*model.Warehouse, error) {
	var res *model.Warehouse

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

func (d warehouseDao) GetByID(pk string) (*model.Warehouse, error) {
	var res *model.Warehouse

	err := model.FilterDeleted(d.db).Where("id = ?", pk).Find(&res).Error

	return res, err
}

// 更新仓库
func (d warehouseDao) Update(scope WarehouseScope, fields map[string]interface{}) (*model.Warehouse, error) {
	var res *model.Warehouse

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Warehouse{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除仓库
func (d warehouseDao) Delete(pks []string) error {
	var res model.Warehouse

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
