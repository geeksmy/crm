package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewProductDao(ctx context.Context, db *gorm.DB) *productDao {
	return &productDao{
		ctx: ctx,
		db:  db,
	}
}

type productDao struct {
	ctx context.Context
	db  *gorm.DB
}

type ProductScope struct {
	scopes []libsgorm.Scope
}

func NewProductScope() ProductScope {
	return ProductScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s ProductScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s ProductScope) WithPk(id string) ProductScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增产品
func (d productDao) Create(product *model.Product) (*model.Product, error) {
	product.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&product).Error

	return product, err
}

// 产品列表
func (d productDao) List(limit, cursor int, scope ProductScope) ([]*model.Product, *libsgorm.Page, error) {
	var res []*model.Product

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个产品
func (d productDao) Get(scope ProductScope) (*model.Product, error) {
	var res *model.Product

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

func (d productDao) GetByID(pk string) (*model.Product, error) {
	var res *model.Product

	err := model.FilterDeleted(d.db).Where("id = ?", pk).Find(&res).Error

	return res, err
}

// 更新产品
func (d productDao) Update(scope ProductScope, fields map[string]interface{}) (*model.Product, error) {
	var res *model.Product

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Product{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除产品
func (d productDao) Delete(pks []string) error {
	var res model.Product

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
