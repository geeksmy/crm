package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"crm/config"
	"crm/gen/log"

	"github.com/geeksmy/go-libs/jwt"
	"goa.design/goa/v3/security"
)

type (
	// private type used to define context keys
	ctxKey int
)

var (
	ErrorUnauthorized = errors.New("请登录后再试")

	KeyUserID ctxKey = 99
)

type Auther struct {
	logger *log.Logger
	// cache  *cache.CacheService
}

// APIKeyAuth implements the authorization logic for service "rework" for the
// "api_key" security scheme.
func (a *Auther) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	// tifUid, _ := ctx.Value(middleware.RequestTifUidKey).(string)
	// // 保存用户ID
	// ctx = context.WithValue(ctx, KeyUserID, tifUid)
	return ctx, nil
}

// JWTAuth implements the authorization logic for service "secured_service" for
// the "jwt" security scheme.
func (a *Auther) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// parse && verify JWT token,
	validator := jwt.NewValidator()

	// 1. parse JWT token,
	userClaims, err := validator.Verify(token, scheme)
	if err != nil {
		return ctx, err
	}

	// 2. validate provided "scopes" claim
	if userClaims.Scopes == nil || len(userClaims.Scopes) == 0 {
		return ctx, MakeUnauthorizedError(ctx, "invalid scopes in token")
	}

	// 3. 检查权限
	if err := a.validateScopes(scheme.RequiredScopes, userClaims.Scopes); err != nil {
		return ctx, MakeForbiddenError(ctx, err.Error())
	}

	if userClaims.ID == "" {
		return ctx, MakeUnauthorizedError(ctx, "无操作权限")
	}

	// 保存用户ID
	ctx = context.WithValue(ctx, KeyUserID, userClaims.ID)

	return ctx, nil
}

// create JWT token
// nolint: unused
func (a *Auther) createJwtToken(userID string, userType int, scopes []string) (string, error) {
	userClaims := jwt.UserClaims{
		ID:        userID,
		Scopes:    scopes,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(config.C.Jwt.ExpireIn)).Unix(),
	}

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	return jwt.GetToken(userClaims)
}

func (a *Auther) validateScopes(expected, actual []string) error {
	for _, r := range expected {
		found := false

		for _, s := range actual {
			if s == r {
				found = true
				break
			}
		}

		if found {
			return nil
		}
	}

	return fmt.Errorf("您没有权限进行此操作")
}

// 获取当前用户ID
func (a *Auther) GetCurrentUserID(ctx context.Context) (string, error) {
	if ctx != nil {
		userID, _ := ctx.Value(KeyUserID).(string)
		return userID, nil
	}

	return "", errors.New("ctx is nil")
}
