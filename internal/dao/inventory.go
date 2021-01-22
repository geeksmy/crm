package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewInventoryDao(ctx context.Context, db *gorm.DB) *inventoryDao {
	return &inventoryDao{
		ctx: ctx,
		db:  db,
	}
}

type inventoryDao struct {
	ctx context.Context
	db  *gorm.DB
}

type InventoryScope struct {
	scopes []libsgorm.Scope
}

func NewInventoryScope() InventoryScope {
	return InventoryScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s InventoryScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s InventoryScope) WithPk(id string) InventoryScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增库存
func (d inventoryDao) Create(inventory model.Inventory) (*model.Inventory, error) {
	inventory.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&inventory).Error

	return &inventory, err
}

// 库存列表
func (d inventoryDao) List(limit, cursor int, scope InventoryScope) ([]*model.Inventory, *libsgorm.Page, error) {
	var res []*model.Inventory

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个库存
func (d inventoryDao) Get(scope InventoryScope) (*model.Inventory, error) {
	var res *model.Inventory

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

// 更新库存
func (d inventoryDao) Update(scope InventoryScope, fields map[string]interface{}) (*model.Inventory, error) {
	var res *model.Inventory

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Inventory{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除库存
func (d inventoryDao) Delete(pks []string) error {
	var res model.Inventory

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
