package handler

import (
	"context"

	"crm/gen/group"
	"crm/gen/log"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
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

	groupService := service.NewGroupService(ctx, dao.DB, logger)
	groupModel, errMsg := groupService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个组失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModeGroup2Group(groupModel)

	return res, nil
}

// 获取组列表
func (s *groupsrvc) List(ctx context.Context, p *group.ListPayload) (res *group.ListResult, err error) {
	res = &group.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("group.List")

	groupService := service.NewGroupService(ctx, dao.DB, logger)
	groupsModel, page, errMsg := groupService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取组列表失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModeGroups2Groups(groupsModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新组
func (s *groupsrvc) Update(ctx context.Context, p *group.UpdatePayload) (res *group.Group, err error) {
	res = &group.Group{}
	logger := L(ctx, s.logger)
	logger.Info("group.Update")

	groupService := service.NewGroupService(ctx, dao.DB, logger)

	groupModel, errMsg := groupService.Update(p.ID, p.Name)

	if errMsg != nil {
		logger.Error("更新组失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModeGroup2Group(groupModel)
	return res, nil
}

// 创建组
func (s *groupsrvc) Create(ctx context.Context, p *group.CreatePayload) (res *group.Group, err error) {
	res = &group.Group{}
	logger := L(ctx, s.logger)
	logger.Info("group.Create")

	groupService := service.NewGroupService(ctx, dao.DB, logger)

	groupModel, errMsg := groupService.Create(p.Name)

	if errMsg != nil {
		logger.Error("创建组失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModeGroup2Group(groupModel)
	return res, nil
}

// 删除组
func (s *groupsrvc) Delete(ctx context.Context, p *group.DeletePayload) (res *group.SuccessResult, err error) {
	res = &group.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("group.Delete")

	groupService := service.NewGroupService(ctx, dao.DB, logger)

	errMsg := groupService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("删除组失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
