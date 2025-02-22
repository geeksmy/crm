// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Supplier client
//
// Command:
// $ goa gen crm/design

package supplier

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Supplier" service client.
type Client struct {
	GetEndpoint    goa.Endpoint
	ListEndpoint   goa.Endpoint
	UpdateEndpoint goa.Endpoint
	CreateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "Supplier" service client given the endpoints.
func NewClient(get, list, update, create, delete_ goa.Endpoint) *Client {
	return &Client{
		GetEndpoint:    get,
		ListEndpoint:   list,
		UpdateEndpoint: update,
		CreateEndpoint: create,
		DeleteEndpoint: delete_,
	}
}

// Get calls the "Get" endpoint of the "Supplier" service.
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *Supplier, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Supplier), nil
}

// List calls the "List" endpoint of the "Supplier" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res *ListResult, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}

// Update calls the "Update" endpoint of the "Supplier" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *Supplier, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Supplier), nil
}

// Create calls the "Create" endpoint of the "Supplier" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res *Supplier, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Supplier), nil
}

// Delete calls the "Delete" endpoint of the "Supplier" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (res *SuccessResult, err error) {
	var ires interface{}
	ires, err = c.DeleteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SuccessResult), nil
}
