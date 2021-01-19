package handler

import (
	"context"

	"crm/config"
	"crm/gen/log"

	"go.uber.org/zap"
	"goa.design/goa/v3/middleware"
	goa "goa.design/goa/v3/pkg"
)

func getErrorID(ctx context.Context) string {
	reqID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		return goa.NewErrorID()
	}

	return reqID
}

func L(ctx context.Context, logger *log.Logger) *zap.Logger {
	reqID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		return logger.Desugar().With(zap.String("reqID", reqID))
	}

	return logger.Desugar()
}

// 创建内部错误
func MakeInternalServerError(ctx context.Context, errmsg string) *goa.ServiceError {
	if errmsg == "" || !config.C.Debug {
		errmsg = "服务器开小差了，稍后再试吧"
	}

	return &goa.ServiceError{
		Name:    "internal_server_error",
		ID:      getErrorID(ctx),
		Message: errmsg,
		Fault:   true,
	}
}

// 创建参数错误
func MakeBadRequestError(ctx context.Context, errmsg string) *goa.ServiceError {
	if errmsg == "" {
		errmsg = "参数错误"
	}

	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      getErrorID(ctx),
		Message: errmsg,
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorizedError(ctx context.Context, errmsg string) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      getErrorID(ctx),
		Message: errmsg,
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbiddenError(ctx context.Context, errmsg string) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      getErrorID(ctx),
		Message: errmsg,
	}
}
