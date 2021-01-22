package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewSalesDao(ctx context.Context, db *gorm.DB) *salesDao {
	return &salesDao{
		ctx: ctx,
		db:  db,
	}
}

type salesDao struct {
	ctx context.Context
	db  *gorm.DB
}

type SalesScope struct {
	scopes []libsgorm.Scope
}

func NewSalesScope() SalesScope {
	return SalesScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s SalesScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s SalesScope) WithPk(id string) SalesScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增销售
func (d salesDao) Create(sales model.Sales) (model.Sales, error) {
	sales.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&sales).Error

	return sales, err
}

// 销售列表
func (d salesDao) List(limit, cursor int, scope SalesScope) ([]model.Sales, *libsgorm.Page, error) {
	var res []model.Sales

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个销售
func (d salesDao) Get(scope SalesScope) (model.Sales, error) {
	var res model.Sales

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

// 更新销售
func (d salesDao) Update(scope SalesScope, fields map[string]interface{}) (model.Sales, error) {
	var res model.Sales

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Sales{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除销售
func (d salesDao) Delete(pks []string) error {
	var res model.Sales

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
