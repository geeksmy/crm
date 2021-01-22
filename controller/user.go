package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/user"
	"crm/internal/dao"
	"crm/internal/model"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
)

// User service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	Auther
	logger *log.Logger
}

// NewUser returns the User service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{Auther{logger: logger}, logger}
}

// 获取单个用户
func (s *usersrvc) Get(ctx context.Context, p *user.GetPayload) (res *user.User, err error) {
	res = &user.User{}
	logger := L(ctx, s.logger)
	logger.Info("user.Get")

	userService := service.NewUserService(ctx, dao.DB, logger)
	resUser, errMsg := userService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个用户失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	res = serializer.ModelUser2User(resUser)

	return res, nil
}

// 获取用户列表
func (s *usersrvc) List(ctx context.Context, p *user.ListPayload) (res *user.ListResult, err error) {
	res = &user.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.List")

	userService := service.NewUserService(ctx, dao.DB, logger)
	resUsers, page, errMsg := userService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取用户列表失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	res.Items = serializer.ModelUsers2Users(resUsers)

	res.NextCursor = page.NextCursor
	res.Total = page.TotalRecord

	return res, nil
}

// 更新用户
func (s *usersrvc) Update(ctx context.Context, p *user.UpdatePayload) (res *user.User, err error) {
	res = &user.User{}
	logger := L(ctx, s.logger)
	logger.Info("user.Update")

	userService := service.NewUserService(ctx, dao.DB, logger)
	resUser, errMsg := userService.Update(p)

	if errMsg != nil {
		logger.Error("更新用户失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	res = serializer.ModelUser2User(resUser)

	return res, nil
}

// 创建用户
func (s *usersrvc) Create(ctx context.Context, p *user.CreatePayload) (res *user.User, err error) {
	res = &user.User{}
	logger := L(ctx, s.logger)
	logger.Info("user.Create")

	account := model.User{
		Username:   p.Username,
		Name:       p.Name,
		Mobile:     p.Mobile,
		Email:      p.Email,
		Jobs:       p.Jobs,
		IsAdmin:    p.IsAdmin,
		SuperiorID: p.SuperiorID,
		GroupID:    p.GroupID,
	}

	password, errMsg := account.CreatePassword(p.Password)

	if errMsg != nil {
		logger.Error("加密密码失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	account.Password = password

	userService := service.NewUserService(ctx, dao.DB, logger)
	userModel, errMsg := userService.Create(&account)

	if errMsg != nil {
		logger.Error("创建用户失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	res = serializer.ModelUser2User(userModel)

	return res, nil
}

// 删除用户
func (s *usersrvc) Delete(ctx context.Context, p *user.DeletePayload) (res *user.SuccessResult, err error) {
	res = &user.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.Delete")

	userService := service.NewUserService(ctx, dao.DB, logger)
	err = userService.Delete(p.Ids)

	if err != nil {
		logger.Error("删除失败", zap.Error(err))
		return nil, service.ErrAuthInternalError
	}
	res.OK = true
	return
}
