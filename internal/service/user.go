package service

import (
	"context"
	"errors"
	"time"

	"crm/gen/user"
	"crm/internal/dao"
	"crm/internal/model"
	"crm/pkg/helper"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewUserService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *UseService {
	return &UseService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type UseService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

var (
	ErrAuthInternalError     = user.MakeInternalServerError(helper.ErrInternal)
	ErrAuthUserNotFoundError = user.MakeBadRequest(helper.ErrUserNotFound)
	ErrAuthUserPasswordError = user.MakeBadRequest(errors.New("密码错误"))
)

func (srv UseService) GetUserByUsername(username string) (*model.User, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	scopes := dao.NewUserScope()
	scopes = scopes.WithUsername(username)

	userModel, err := userDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取用户失败", zap.Error(err))
		return nil, err
	}

	res, err := srv.getUserBySuperiorIDAndGroupID(userModel)
	if err != nil {
		srv.logger.Error("获取上级或组失败", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (srv UseService) Create(account *model.User) (*model.User, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)

	userModel, err := userDao.Create(account)

	if err != nil {
		srv.logger.Error("创建用户失败", zap.Error(err))
		return nil, err
	}

	res, err := srv.getUserBySuperiorIDAndGroupID(userModel)
	if err != nil {
		srv.logger.Error("获取上级或组失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv UseService) Get(pk string) (*model.User, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	scopes := dao.NewUserScope()
	scopes = scopes.WithPk(pk)

	userModel, err := userDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取用户失败", zap.Error(err))
		return nil, err
	}

	res, err := srv.getUserBySuperiorIDAndGroupID(userModel)
	if err != nil {
		srv.logger.Error("获取上级或组失败", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (srv UseService) List(limit, cursor int) ([]*model.User, *libsgorm.Page, error) {
	var res []*model.User
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	scopes := dao.NewUserScope()

	usersModel, page, err := userDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取用户列表失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(usersModel); i++ {

		userModel, err := srv.getUserBySuperiorIDAndGroupID(usersModel[i])
		if err != nil {
			srv.logger.Error("获取上级或组失败", zap.Error(err))
			return nil, nil, err
		}

		res = append(res, userModel)
	}

	return res, page, nil
}

func (srv UseService) Update(p *user.UpdatePayload) (*model.User, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	scopes := dao.NewUserScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})
	if p.Name != nil && *p.Name != "" {
		fields["name"] = *p.Name
	}
	if p.Mobile != nil && *p.Mobile != "" {
		fields["mobile"] = *p.Mobile
	}
	if p.Email != nil && *p.Email != "" {
		fields["email"] = *p.Email
	}
	if p.SuperiorID != nil && *p.SuperiorID != "" {
		fields["superior_id"] = *p.SuperiorID
	}
	if p.GroupID != nil && *p.GroupID != "" {
		fields["group_id"] = *p.GroupID
	}
	if p.Jobs != nil && *p.Jobs != 0 {
		fields["jobs"] = *p.Jobs
	}

	userModel, err := userDao.Update(scopes, fields)
	if err != nil {
		srv.logger.Error("更新用户失败", zap.Error(err))
		return nil, err
	}

	res, err := srv.getUserBySuperiorIDAndGroupID(userModel)
	if err != nil {
		srv.logger.Error("获取上级或组失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv UseService) UpdateByLoginAt(pk, password string) error {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	scopes := dao.NewUserScope()
	scopes = scopes.WithPk(pk)

	fields := make(map[string]interface{})
	fields["login_at"] = time.Now()

	if password != "" {
		fields["password"] = password
	}

	_, err := userDao.Update(scopes, fields)
	if err != nil {
		srv.logger.Error("更新用户失败", zap.Error(err))
		return err
	}

	return nil
}

func (srv UseService) Delete(pks []string) error {
	userDao := dao.NewUserDao(srv.ctx, srv.db)

	err := userDao.Delete(pks)

	return err
}

func (srv UseService) getUserBySuperiorIDAndGroupID(user *model.User) (*model.User, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	groupDao := dao.NewGroupDao(srv.ctx, srv.db)

	if user.SuperiorID != "" {
		superior, err := userDao.GetByID(user.SuperiorID)

		if err != nil {
			srv.logger.Error("获取上级失败", zap.Error(err))
			return nil, err
		}

		user.Superior.ID = superior.ID
		user.Superior.Name = superior.Name
	}

	if user.GroupID != "" {
		group, err := groupDao.GetByID(user.GroupID)

		if err != nil {
			srv.logger.Error("获取用户组失败", zap.Error(err))
			return nil, err
		}

		user.Group.ID = group.ID
		user.Group.Name = group.Name
	}

	return user, nil
}
