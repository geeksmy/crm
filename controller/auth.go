package handler

import (
	"context"

	"crm/config"
	"crm/gen/auth"
	"crm/gen/log"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
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

	userService := service.NewUserService(ctx, dao.DB, logger)
	resUser, errMsg := userService.GetUserByUsername(p.Username)

	// 根据用户名获取用户信息
	if errMsg != nil {
		logger.Error("获取用户失败", zap.Error(errMsg))
		return nil, service.ErrAuthUserNotFoundError
	}

	// 校验密码
	if !resUser.CheckPassword(p.Password, resUser.Password) {
		logger.Error("密码错误", zap.Error(err))
		return nil, service.ErrAuthUserPasswordError
	}

	// 更新登录时间
	errMsg = userService.UpdateByLoginAt(resUser.ID, "")

	if errMsg != nil {
		logger.Error("更新登录时间失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	// 生成token
	token, errMsg := s.createJwtToken(resUser.ID, 1, []string{"role:user", "role:admin"})

	if errMsg != nil {
		logger.Error("签发token失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	credentials := auth.Credentials{
		Token:     token,
		ExpiresIn: config.C.Jwt.ExpireIn,
	}
	res.Credentials = &credentials
	res.User = serializer.ModelUser2AuthUser(resUser)

	return res, nil
}

// 修改登录密码
func (s *authsrvc) UpdatePassword(ctx context.Context, p *auth.UpdatePasswordPayload) (res *auth.SuccessResult, err error) {
	res = &auth.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("auth.UpdatePassword")

	userService := service.NewUserService(ctx, dao.DB, logger)
	resModel, errMsg := userService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取用户失败", zap.Error(errMsg))
		return nil, service.ErrAuthUserNotFoundError
	}

	// 校验密码
	if !resModel.CheckPassword(p.OldPassword, resModel.Password) {
		logger.Error("密码错误", zap.Error(err))
		return nil, service.ErrAuthUserPasswordError
	}

	// 加密密码
	password, errMsg := resModel.CreatePassword(p.NewPassword)

	if errMsg != nil {
		logger.Error("加密密码失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	// 更新密码
	errMsg = userService.UpdateByLoginAt(resModel.ID, password)

	if errMsg != nil {
		logger.Error("更新密码失败", zap.Error(errMsg))
		return nil, service.ErrAuthInternalError
	}

	res.OK = true

	return res, nil
}

// 获取图形验证码
func (s *authsrvc) CaptchaImage(ctx context.Context) (res *auth.Captcha, err error) {
	res = &auth.Captcha{}
	logger := L(ctx, s.logger)
	logger.Info("auth.CaptchaImage")
	return
}
