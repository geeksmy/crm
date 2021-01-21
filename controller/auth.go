package handler

import (
	"context"

	"crm/gen/auth"
	"crm/gen/log"
)

// Auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	Auther
	logger *log.Logger
}

// NewAuth returns the Auth service implementation.
func NewAuth(logger *log.Logger) auth.Service {
	return &authsrvc{Auther{logger: logger}, logger}
}

// 使用账号密码登录
func (s *authsrvc) Login(ctx context.Context, p *auth.LoginPayload) (res *auth.Session, err error) {
	res = &auth.Session{}
	logger := L(ctx, s.logger)
	logger.Info("auth.Login")
	return
}

// 修改登录密码
func (s *authsrvc) UpdatePassword(ctx context.Context, p *auth.UpdatePasswordPayload) (res *auth.SuccessResult, err error) {
	res = &auth.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("auth.UpdatePassword")
	return
}

// 获取图形验证码
func (s *authsrvc) CaptchaImage(ctx context.Context) (res *auth.Captcha, err error) {
	res = &auth.Captcha{}
	logger := L(ctx, s.logger)
	logger.Info("auth.CaptchaImage")
	return
}
