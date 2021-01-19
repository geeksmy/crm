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
	return &usersrvc{
		logger: logger,
		Auther: Auther{
			logger: logger,
		},
	}
}

// 使用账号密码登录
func (s *usersrvc) LoginByUsername(ctx context.Context, p *user.LoginByUsernamePayload) (res *user.LoginByUsernameResult, err error) {
	res = &user.LoginByUsernameResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.LoginByUsername")
	return
}

// 修改登录密码
func (s *usersrvc) UpdatePassword(ctx context.Context, p *user.UpdatePasswordPayload) (res *user.UpdatePasswordResult, err error) {
	res = &user.UpdatePasswordResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.UpdatePassword")
	return
}

// 获取图形验证码
func (s *usersrvc) GetCaptchaImage(ctx context.Context) (res *user.GetCaptchaImageResult, err error) {
	res = &user.GetCaptchaImageResult{}
	logger := L(ctx, s.logger)
	logger.Info("user.GetCaptchaImage")
	return
}
