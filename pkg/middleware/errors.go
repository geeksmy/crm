package middleware

import (
	"net/http"

	"crm/config"

	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

type ErrorResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`

	// Name is a name for that class of errors.
	Name string `json:"-"`
	// ID is a unique value for each occurrence of the error.
	ID string `json:"-"`
	// Message contains the specific error details.
	Message string `json:"-"`
	// Is the error a timeout?
	Timeout bool `json:"-"`
	// Is the error temporary?
	Temporary bool `json:"-"`
	// Is the error a server-side fault?
	Fault bool `json:"-"`
}

func GoaErrorFormatterFunc(err error) goahttp.Statuser {
	return NewErrorResponse(err)
}

// NewErrorResponse creates a HTTP response from the given error.
func NewErrorResponse(err error) goahttp.Statuser {
	// nolint: errorlint
	// fixme: need remove nolint
	if goaerr, ok := err.(*goa.ServiceError); ok {
		resp := &ErrorResponse{
			Errcode: 0,
			Errmsg:  goaerr.Message,

			Name:      goaerr.Name,
			ID:        goaerr.ID,
			Message:   goaerr.Message,
			Timeout:   goaerr.Timeout,
			Temporary: goaerr.Temporary,
			Fault:     goaerr.Fault,
		}

		if resp.Fault && config.C.Logger.Level != "debug" {
			resp.Errmsg = "服务器开小差了，稍后再试吧"
		}

		resp.Errcode = 600000 + resp.StatusCode()

		zap.L().Error("handle error:", zap.Error(err),
			zap.String("Message", resp.Message),
			zap.Int("Errcode", resp.Errcode),
			zap.Int("statusCode", resp.StatusCode()),
			zap.String("errorID", goaerr.ID),
		)

		return resp
	}

	return NewErrorResponse(goa.Fault(err.Error()))
}

// StatusCode implements a heuristic that computes a HTTP response status code
// appropriate for the timeout, temporary and fault characteristics of the
// error. This method is used by the generated server code when the error is not
// described explicitly in the design.
func (resp *ErrorResponse) StatusCode() int {
	if resp.Name == "unauthorized" {
		return http.StatusUnauthorized
	}

	if resp.Name == "forbidden" {
		return http.StatusForbidden
	}

	if resp.Fault {
		return http.StatusInternalServerError
	}

	if resp.Timeout {
		if resp.Temporary {
			return http.StatusGatewayTimeout
		}

		return http.StatusRequestTimeout
	}

	if resp.Temporary {
		return http.StatusServiceUnavailable
	}

	return http.StatusBadRequest
}
