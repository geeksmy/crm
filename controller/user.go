package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/user"
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
	return
}

// 获取用户列表
func (s *usersrvc) List(ctx context.Context, p *user.ListPayload) (res *user.ListResult, err error) {
	res = &user.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.List")
	return
}

// 更新用户
func (s *usersrvc) Update(ctx context.Context, p *user.UpdatePayload) (res *user.User, err error) {
	res = &user.User{}
	logger := L(ctx, s.logger)
	logger.Info("user.Update")
	return
}

// 创建用户
func (s *usersrvc) Create(ctx context.Context, p *user.CreatePayload) (res *user.User, err error) {
	res = &user.User{}
	logger := L(ctx, s.logger)
	logger.Info("user.Create")
	return
}

// 删除用户
func (s *usersrvc) Delete(ctx context.Context, p *user.DeletePayload) (res *user.SuccessResult, err error) {
	res = &user.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.Delete")
	return
}
