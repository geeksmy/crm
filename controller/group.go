package handler

import (
	"context"

	"crm/gen/group"
	"crm/gen/log"
)

// Group service example implementation.
// The example methods log the requests and return zero values.
type groupsrvc struct {
	authsrvc
	logger *log.Logger
}

// NewGroup returns the Group service implementation.
func NewGroup(logger *log.Logger) group.Service {
	return &groupsrvc{authsrvc{logger: logger}, logger}
}

// 获取单个组
func (s *groupsrvc) Get(ctx context.Context, p *group.GetPayload) (res *group.Group, err error) {
	res = &group.Group{}
	logger := L(ctx, s.logger)
	logger.Info("group.Get")
	return
}

// 获取组列表
func (s *groupsrvc) List(ctx context.Context, p *group.ListPayload) (res *group.ListResult, err error) {
	res = &group.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("group.List")
	return
}

// 更新组
func (s *groupsrvc) Update(ctx context.Context, p *group.UpdatePayload) (res *group.Group, err error) {
	res = &group.Group{}
	logger := L(ctx, s.logger)
	logger.Info("group.Update")
	return
}

// 创建组
func (s *groupsrvc) Create(ctx context.Context, p *group.CreatePayload) (res *group.Group, err error) {
	res = &group.Group{}
	logger := L(ctx, s.logger)
	logger.Info("group.Create")
	return
}

// 删除组
func (s *groupsrvc) Delete(ctx context.Context, p *group.DeletePayload) (res *group.SuccessResult, err error) {
	res = &group.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("group.Delete")
	return
}
