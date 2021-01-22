package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewCustomerDao(ctx context.Context, db *gorm.DB) *customerDao {
	return &customerDao{
		ctx: ctx,
		db:  db,
	}
}

type customerDao struct {
	ctx context.Context
	db  *gorm.DB
}

type CustomerScope struct {
	scopes []libsgorm.Scope
}

func NewCustomerScope() CustomerScope {
	return CustomerScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s CustomerScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s CustomerScope) WithPk(id string) CustomerScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增客户
func (d customerDao) Create(customer model.Customer) (*model.Customer, error) {
	customer.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&customer).Error

	return &customer, err
}

// 客户列表
func (d customerDao) List(limit, cursor int, scope CustomerScope) ([]*model.Customer, *libsgorm.Page, error) {
	var res []*model.Customer

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个客户
func (d customerDao) Get(scope CustomerScope) (*model.Customer, error) {
	var res *model.Customer

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

func (d customerDao) GetByID(pk string) (*model.Customer, error) {
	var res *model.Customer

	err := model.FilterDeleted(d.db).Where("id = ?", pk).Find(&res).Error

	return res, err
}

// 更新客户
func (d customerDao) Update(scope CustomerScope, fields map[string]interface{}) (*model.Customer, error) {
	var res *model.Customer

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Customer{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除客户
func (d customerDao) Delete(pks []string) error {
	var res model.Customer

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
