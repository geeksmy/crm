// Code generated by goa v3.2.6, DO NOT EDIT.
//
// User client
//
// Command:
// $ goa gen crm/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "User" service client.
type Client struct {
	LoginByUsernameEndpoint goa.Endpoint
	UpdatePasswordEndpoint  goa.Endpoint
	GetCaptchaImageEndpoint goa.Endpoint
}

// NewClient initializes a "User" service client given the endpoints.
func NewClient(loginByUsername, updatePassword, getCaptchaImage goa.Endpoint) *Client {
	return &Client{
		LoginByUsernameEndpoint: loginByUsername,
		UpdatePasswordEndpoint:  updatePassword,
		GetCaptchaImageEndpoint: getCaptchaImage,
	}
}

// LoginByUsername calls the "LoginByUsername" endpoint of the "User" service.
func (c *Client) LoginByUsername(ctx context.Context, p *LoginByUsernamePayload) (res *LoginByUsernameResult, err error) {
	var ires interface{}
	ires, err = c.LoginByUsernameEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginByUsernameResult), nil
}

// UpdatePassword calls the "UpdatePassword" endpoint of the "User" service.
func (c *Client) UpdatePassword(ctx context.Context, p *UpdatePasswordPayload) (res *UpdatePasswordResult, err error) {
	var ires interface{}
	ires, err = c.UpdatePasswordEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdatePasswordResult), nil
}

// GetCaptchaImage calls the "GetCaptchaImage" endpoint of the "User" service.
func (c *Client) GetCaptchaImage(ctx context.Context) (res *GetCaptchaImageResult, err error) {
	var ires interface{}
	ires, err = c.GetCaptchaImageEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*GetCaptchaImageResult), nil
}