// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Inventory HTTP client encoders and decoders
//
// Command:
// $ goa gen crm/design

package client

import (
	"bytes"
	"context"
	inventory "crm/gen/inventory"
	inventoryviews "crm/gen/inventory/views"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "Inventory" service "Get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*inventory.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Inventory", "Get", "*inventory.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetInventoryPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Inventory", "Get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the Inventory Get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("Inventory", "Get", "*inventory.GetPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the Inventory
// Get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Get", err)
			}
			p := NewGetInventoryOK(&body)
			view := "default"
			vres := &inventoryviews.Inventory{Projected: p, View: view}
			if err = inventoryviews.ValidateInventory(vres); err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Get", err)
			}
			res := inventory.NewInventory(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Get", err)
			}
			err = ValidateGetInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Get", err)
			}
			return nil, NewGetInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Get", err)
			}
			return nil, NewGetBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Inventory", "Get", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "Inventory" service "List" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListInventoryPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Inventory", "List", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the Inventory List
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("Inventory", "List", "*inventory.ListPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		if p.Cursor != nil {
			values.Add("cursor", fmt.Sprintf("%v", *p.Cursor))
		}
		if p.Limit != nil {
			values.Add("limit", fmt.Sprintf("%v", *p.Limit))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the Inventory
// List endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "List", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "List", err)
			}
			res := NewListResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body ListInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "List", err)
			}
			err = ValidateListInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "List", err)
			}
			return nil, NewListInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "List", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "List", err)
			}
			return nil, NewListBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Inventory", "List", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "Inventory" service "Update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateInventoryPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Inventory", "Update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the Inventory
// Update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Inventory", "Update", "*inventory.UpdatePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Inventory", "Update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// Inventory Update endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Update", err)
			}
			p := NewUpdateInventoryOK(&body)
			view := "default"
			vres := &inventoryviews.Inventory{Projected: p, View: view}
			if err = inventoryviews.ValidateInventory(vres); err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Update", err)
			}
			res := inventory.NewInventory(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Update", err)
			}
			err = ValidateUpdateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Update", err)
			}
			return nil, NewUpdateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Inventory", "Update", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "Inventory" service "Create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateInventoryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Inventory", "Create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the Inventory
// Create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Inventory", "Create", "*inventory.CreatePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Inventory", "Create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// Inventory Create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Create", err)
			}
			p := NewCreateInventoryOK(&body)
			view := "default"
			vres := &inventoryviews.Inventory{Projected: p, View: view}
			if err = inventoryviews.ValidateInventory(vres); err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Create", err)
			}
			res := inventory.NewInventory(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Create", err)
			}
			err = ValidateCreateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Create", err)
			}
			return nil, NewCreateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Create", err)
			}
			return nil, NewCreateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Inventory", "Create", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "Inventory" service "Delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteInventoryPath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Inventory", "Delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the Inventory
// Delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*inventory.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("Inventory", "Delete", "*inventory.DeletePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewDeleteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Inventory", "Delete", err)
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// Inventory Delete endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Delete", err)
			}
			p := NewDeleteSuccessResultOK(&body)
			view := "default"
			vres := &inventoryviews.SuccessResult{Projected: p, View: view}
			if err = inventoryviews.ValidateSuccessResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Delete", err)
			}
			res := inventory.NewSuccessResult(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Delete", err)
			}
			err = ValidateDeleteInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Delete", err)
			}
			return nil, NewDeleteInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Inventory", "Delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Inventory", "Delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Inventory", "Delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalProductResponseBodyToInventoryviewsProductView builds a value of
// type *inventoryviews.ProductView from a value of type *ProductResponseBody.
func unmarshalProductResponseBodyToInventoryviewsProductView(v *ProductResponseBody) *inventoryviews.ProductView {
	res := &inventoryviews.ProductView{
		ID:          v.ID,
		Name:        v.Name,
		Unit:        v.Unit,
		CostPrice:   v.CostPrice,
		MarketPrice: v.MarketPrice,
		Note:        v.Note,
		Image:       v.Image,
		Code:        v.Code,
		Size:        v.Size,
		Type:        v.Type,
		IsShelves:   v.IsShelves,
	}
	res.Founder = unmarshalFounderResponseBodyToInventoryviewsFounderView(v.Founder)

	return res
}

// unmarshalFounderResponseBodyToInventoryviewsFounderView builds a value of
// type *inventoryviews.FounderView from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToInventoryviewsFounderView(v *FounderResponseBody) *inventoryviews.FounderView {
	res := &inventoryviews.FounderView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalWarehouseResponseBodyToInventoryviewsWarehouseView builds a value
// of type *inventoryviews.WarehouseView from a value of type
// *WarehouseResponseBody.
func unmarshalWarehouseResponseBodyToInventoryviewsWarehouseView(v *WarehouseResponseBody) *inventoryviews.WarehouseView {
	res := &inventoryviews.WarehouseView{
		ID:      v.ID,
		Name:    v.Name,
		Code:    v.Code,
		Address: v.Address,
		Type:    v.Type,
	}
	res.Founder = unmarshalFounderResponseBodyToInventoryviewsFounderView(v.Founder)

	return res
}

// unmarshalHeadResponseBodyToInventoryviewsHeadView builds a value of type
// *inventoryviews.HeadView from a value of type *HeadResponseBody.
func unmarshalHeadResponseBodyToInventoryviewsHeadView(v *HeadResponseBody) *inventoryviews.HeadView {
	res := &inventoryviews.HeadView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalInventoryResponseBodyToInventoryInventory builds a value of type
// *inventory.Inventory from a value of type *InventoryResponseBody.
func unmarshalInventoryResponseBodyToInventoryInventory(v *InventoryResponseBody) *inventory.Inventory {
	res := &inventory.Inventory{
		ID:            *v.ID,
		Number:        *v.Number,
		Code:          *v.Code,
		Type:          *v.Type,
		InventoryDate: *v.InventoryDate,
		InAndOut:      *v.InAndOut,
		Note:          *v.Note,
	}
	res.Product = unmarshalProductResponseBodyToInventoryProduct(v.Product)
	res.Warehouse = unmarshalWarehouseResponseBodyToInventoryWarehouse(v.Warehouse)
	res.Head = unmarshalHeadResponseBodyToInventoryHead(v.Head)
	res.Founder = unmarshalFounderResponseBodyToInventoryFounder(v.Founder)

	return res
}

// unmarshalProductResponseBodyToInventoryProduct builds a value of type
// *inventory.Product from a value of type *ProductResponseBody.
func unmarshalProductResponseBodyToInventoryProduct(v *ProductResponseBody) *inventory.Product {
	res := &inventory.Product{
		ID:          *v.ID,
		Name:        *v.Name,
		Unit:        *v.Unit,
		CostPrice:   *v.CostPrice,
		MarketPrice: *v.MarketPrice,
		Note:        *v.Note,
		Image:       *v.Image,
		Code:        *v.Code,
		Size:        *v.Size,
		Type:        *v.Type,
		IsShelves:   *v.IsShelves,
	}
	res.Founder = unmarshalFounderResponseBodyToInventoryFounder(v.Founder)

	return res
}

// unmarshalFounderResponseBodyToInventoryFounder builds a value of type
// *inventory.Founder from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToInventoryFounder(v *FounderResponseBody) *inventory.Founder {
	res := &inventory.Founder{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalWarehouseResponseBodyToInventoryWarehouse builds a value of type
// *inventory.Warehouse from a value of type *WarehouseResponseBody.
func unmarshalWarehouseResponseBodyToInventoryWarehouse(v *WarehouseResponseBody) *inventory.Warehouse {
	res := &inventory.Warehouse{
		ID:      *v.ID,
		Name:    *v.Name,
		Code:    *v.Code,
		Address: *v.Address,
		Type:    *v.Type,
	}
	res.Founder = unmarshalFounderResponseBodyToInventoryFounder(v.Founder)

	return res
}

// unmarshalHeadResponseBodyToInventoryHead builds a value of type
// *inventory.Head from a value of type *HeadResponseBody.
func unmarshalHeadResponseBodyToInventoryHead(v *HeadResponseBody) *inventory.Head {
	res := &inventory.Head{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}
