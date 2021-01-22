package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewProcurementDao(ctx context.Context, db *gorm.DB) *procurementDao {
	return &procurementDao{
		ctx: ctx,
		db:  db,
	}
}

type procurementDao struct {
	ctx context.Context
	db  *gorm.DB
}

type ProcurementScope struct {
	scopes []libsgorm.Scope
}

func NewProcurementScope() ProcurementScope {
	return ProcurementScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s ProcurementScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s ProcurementScope) WithPk(id string) ProcurementScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增采购
func (d procurementDao) Create(procurement model.Procurement) (*model.Procurement, error) {
	procurement.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&procurement).Error

	return &procurement, err
}

// 采购列表
func (d procurementDao) List(limit, cursor int, scope ProcurementScope) ([]*model.Procurement, *libsgorm.Page, error) {
	var res []*model.Procurement

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个采购
func (d procurementDao) Get(scope ProcurementScope) (*model.Procurement, error) {
	var res *model.Procurement

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

// 更新采购
func (d procurementDao) Update(scope ProcurementScope, fields map[string]interface{}) (*model.Procurement, error) {
	var res *model.Procurement

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Procurement{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除采购
func (d procurementDao) Delete(pks []string) error {
	var res model.Procurement

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
