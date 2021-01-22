package service

import (
	"context"

	"crm/gen/group"
	"crm/internal/dao"
	"crm/internal/model"
	"crm/pkg/helper"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewGroupService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *GroupService {
	return &GroupService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type GroupService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

var (
	ErrInternalError = group.MakeInternalServerError(helper.ErrInternal)
)

func (srv GroupService) Get(pk string) (*model.Group, error) {
	groupDao := dao.NewGroupDao(srv.ctx, srv.db)

	res, err := groupDao.GetByID(pk)

	if err != nil {
		srv.logger.Error("获取组失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv GroupService) List(limit, cursor int) ([]*model.Group, *libsgorm.Page, error) {
	groupDao := dao.NewGroupDao(srv.ctx, srv.db)
	scopes := dao.NewGroupScope()

	res, page, err := groupDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取组列表失败", zap.Error(err))
		return nil, nil, err
	}
	return res, page, nil
}

func (srv GroupService) Create(name string) (*model.Group, error) {
	groupDao := dao.NewGroupDao(srv.ctx, srv.db)

	res, err := groupDao.Create(name)
	if err != nil {
		srv.logger.Error("创建组", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv GroupService) Update(pk string, name *string) (*model.Group, error) {
	groupDao := dao.NewGroupDao(srv.ctx, srv.db)
	scopes := dao.NewGroupScope()
	scopes = scopes.WithPk(pk)

	fields := make(map[string]interface{})

	if name != nil && *name != "" {
		fields["name"] = *name
	}

	res, err := groupDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新组失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv GroupService) Delete(pks []string) error {
	groupDao := dao.NewGroupDao(srv.ctx, srv.db)

	err := groupDao.Delete(pks)

	return err
}
