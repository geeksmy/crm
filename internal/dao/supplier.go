package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewSupplierDao(ctx context.Context, db *gorm.DB) *supplierDao {
	return &supplierDao{
		ctx: ctx,
		db:  db,
	}
}

type supplierDao struct {
	ctx context.Context
	db  *gorm.DB
}

type SupplierScope struct {
	scopes []libsgorm.Scope
}

func NewSupplierScope() SupplierScope {
	return SupplierScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s SupplierScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s SupplierScope) WithPk(id string) SupplierScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增供应商
func (d supplierDao) Create(supplier model.Supplier) (*model.Supplier, error) {
	supplier.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&supplier).Error

	return &supplier, err
}

// 供应商列表
func (d supplierDao) List(limit, cursor int, scope SupplierScope) ([]*model.Supplier, *libsgorm.Page, error) {
	var res []*model.Supplier

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个供应商
func (d supplierDao) Get(scope SupplierScope) (*model.Supplier, error) {
	var res *model.Supplier

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

func (d supplierDao) GetByID(pk string) (*model.Supplier, error) {
	var res *model.Supplier

	err := model.FilterDeleted(d.db).Where("id = ?", pk).Find(&res).Error

	return res, err
}

// 更新供应商
func (d supplierDao) Update(scope SupplierScope, fields map[string]interface{}) (*model.Supplier, error) {
	var res *model.Supplier

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Supplier{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除供应商
func (d supplierDao) Delete(pks []string) error {
	var res model.Supplier

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
