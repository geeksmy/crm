// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Auth endpoints
//
// Command:
// $ goa gen crm/design

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "Auth" service endpoints.
type Endpoints struct {
	Login          goa.Endpoint
	UpdatePassword goa.Endpoint
	CaptchaImage   goa.Endpoint
}

// NewEndpoints wraps the methods of the "Auth" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Login:          NewLoginEndpoint(s),
		UpdatePassword: NewUpdatePasswordEndpoint(s, a.JWTAuth),
		CaptchaImage:   NewCaptchaImageEndpoint(s),
	}
}

// Use applies the given middleware to all the "Auth" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Login = m(e.Login)
	e.UpdatePassword = m(e.UpdatePassword)
	e.CaptchaImage = m(e.CaptchaImage)
}

// NewLoginEndpoint returns an endpoint function that calls the method "Login"
// of service "Auth".
func NewLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		res, err := s.Login(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSession(res, "default")
		return vres, nil
	}
}

// NewUpdatePasswordEndpoint returns an endpoint function that calls the method
// "UpdatePassword" of service "Auth".
func NewUpdatePasswordEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePasswordPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"role:user", "role:admin"},
			RequiredScopes: []string{"role:user"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.UpdatePassword(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessResult(res, "default")
		return vres, nil
	}
}

// NewCaptchaImageEndpoint returns an endpoint function that calls the method
// "CaptchaImage" of service "Auth".
func NewCaptchaImageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.CaptchaImage(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCaptcha(res, "default")
		return vres, nil
	}
}
