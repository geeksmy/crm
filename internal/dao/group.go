package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewGroupDao(ctx context.Context, db *gorm.DB) *groupDao {
	return &groupDao{
		ctx: ctx,
		db:  db,
	}
}

type groupDao struct {
	ctx context.Context
	db  *gorm.DB
}

type GroupScope struct {
	scopes []libsgorm.Scope
}

func NewGroupScope() GroupScope {
	return GroupScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s GroupScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s GroupScope) WithPk(id string) GroupScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

// 新增组
func (d groupDao) Create(name string) (*model.Group, error) {
	res := model.Group{
		BaseUUIDModel: model.NewBaseUUIDModel(),
		Name:          name,
	}
	err := d.db.Create(&res).Error

	return &res, err
}

// 组列表
func (d groupDao) List(limit, cursor int, scope GroupScope) ([]*model.Group, *libsgorm.Page, error) {
	var res []*model.Group

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个组
func (d groupDao) Get(scope GroupScope) (*model.Group, error) {
	var res *model.Group

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

// 根据ID获取组
func (d groupDao) GetByID(pk string) (*model.Group, error) {
	var res *model.Group

	err := model.FilterDeleted(d.db).Where("id = ?", pk).Find(&res).Error

	return res, err
}

// 更新组
func (d groupDao) Update(scope GroupScope, fields map[string]interface{}) (*model.Group, error) {
	var res *model.Group

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.Group{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除组
func (d groupDao) Delete(pks []string) error {
	var res model.Group

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
