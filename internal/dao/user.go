package dao

import (
	"context"

	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
)

func NewUserDao(ctx context.Context, db *gorm.DB) *userDao {
	return &userDao{
		ctx: ctx,
		db:  db,
	}
}

type userDao struct {
	ctx context.Context
	db  *gorm.DB
}

type UserScope struct {
	scopes []libsgorm.Scope
}

func NewUserScope() UserScope {
	return UserScope{
		scopes: []libsgorm.Scope{},
	}
}

// scopes
func (s UserScope) Scopes() []libsgorm.Scope {
	return s.scopes
}

func (s UserScope) WithPk(id string) UserScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return s
}

func (s UserScope) WithUsername(username string) UserScope {
	s.scopes = append(s.scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", username)
	})
	return s
}

// 新增用户
func (d userDao) Create(user *model.User) (*model.User, error) {
	user.BaseUUIDModel = model.NewBaseUUIDModel()
	err := d.db.Create(&user).Error

	return user, err
}

// 用户列表
func (d userDao) List(limit, cursor int, scope UserScope) ([]*model.User, *libsgorm.Page, error) {
	var res []*model.User

	db := d.db.Scopes(scope.Scopes()...)
	query := model.FilterDeleted(db)
	page, err := libsgorm.Pagination(query, limit, cursor, &res)

	return res, page, err
}

// 获取单个用户
func (d userDao) Get(scope UserScope) (*model.User, error) {
	var res *model.User

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Find(&res).Error

	return res, err
}

// 根据ID获取用户
func (d userDao) GetByID(pk string) (*model.User, error) {
	var res *model.User

	err := model.FilterDeleted(d.db).Where("id = ?", pk).Find(&res).Error

	return res, err
}

// 更新用户
func (d userDao) Update(scope UserScope, fields map[string]interface{}) (*model.User, error) {
	var res *model.User

	db := d.db.Scopes(scope.Scopes()...)
	err := model.FilterDeleted(db).Model(&model.User{}).Updates(fields).Scan(&res).Error

	return res, err
}

// 删除用户
func (d userDao) Delete(pks []string) error {
	var res model.User

	db := d.db.Model(&res).Where("id in (?)", pks)
	err := model.SoftDelete(db).Error

	return err
}
