// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Warehouse HTTP client encoders and decoders
//
// Command:
// $ goa gen crm/design

package client

import (
	"bytes"
	"context"
	warehouse "crm/gen/warehouse"
	warehouseviews "crm/gen/warehouse/views"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "Warehouse" service "Get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*warehouse.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Warehouse", "Get", "*warehouse.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetWarehousePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Warehouse", "Get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the Warehouse Get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*warehouse.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("Warehouse", "Get", "*warehouse.GetPayload", v)
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

// DecodeGetResponse returns a decoder for responses returned by the Warehouse
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
				return nil, goahttp.ErrDecodingError("Warehouse", "Get", err)
			}
			p := NewGetWarehouseOK(&body)
			view := "default"
			vres := &warehouseviews.Warehouse{Projected: p, View: view}
			if err = warehouseviews.ValidateWarehouse(vres); err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Get", err)
			}
			res := warehouse.NewWarehouse(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Get", err)
			}
			err = ValidateGetInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Get", err)
			}
			return nil, NewGetInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Get", err)
			}
			return nil, NewGetBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Warehouse", "Get", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "Warehouse" service "List" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListWarehousePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Warehouse", "List", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the Warehouse List
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*warehouse.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("Warehouse", "List", "*warehouse.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the Warehouse
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
				return nil, goahttp.ErrDecodingError("Warehouse", "List", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "List", err)
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
				return nil, goahttp.ErrDecodingError("Warehouse", "List", err)
			}
			err = ValidateListInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "List", err)
			}
			return nil, NewListInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "List", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "List", err)
			}
			return nil, NewListBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Warehouse", "List", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "Warehouse" service "Update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateWarehousePath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Warehouse", "Update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the Warehouse
// Update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*warehouse.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Warehouse", "Update", "*warehouse.UpdatePayload", v)
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
			return goahttp.ErrEncodingError("Warehouse", "Update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// Warehouse Update endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Warehouse", "Update", err)
			}
			p := NewUpdateWarehouseOK(&body)
			view := "default"
			vres := &warehouseviews.Warehouse{Projected: p, View: view}
			if err = warehouseviews.ValidateWarehouse(vres); err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Update", err)
			}
			res := warehouse.NewWarehouse(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Update", err)
			}
			err = ValidateUpdateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Update", err)
			}
			return nil, NewUpdateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Warehouse", "Update", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "Warehouse" service "Create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateWarehousePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Warehouse", "Create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the Warehouse
// Create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*warehouse.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Warehouse", "Create", "*warehouse.CreatePayload", v)
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
			return goahttp.ErrEncodingError("Warehouse", "Create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// Warehouse Create endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Warehouse", "Create", err)
			}
			p := NewCreateWarehouseOK(&body)
			view := "default"
			vres := &warehouseviews.Warehouse{Projected: p, View: view}
			if err = warehouseviews.ValidateWarehouse(vres); err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Create", err)
			}
			res := warehouse.NewWarehouse(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Create", err)
			}
			err = ValidateCreateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Create", err)
			}
			return nil, NewCreateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Create", err)
			}
			return nil, NewCreateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Warehouse", "Create", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "Warehouse" service "Delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteWarehousePath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Warehouse", "Delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the Warehouse
// Delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*warehouse.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("Warehouse", "Delete", "*warehouse.DeletePayload", v)
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
			return goahttp.ErrEncodingError("Warehouse", "Delete", err)
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// Warehouse Delete endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Warehouse", "Delete", err)
			}
			p := NewDeleteSuccessResultOK(&body)
			view := "default"
			vres := &warehouseviews.SuccessResult{Projected: p, View: view}
			if err = warehouseviews.ValidateSuccessResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Delete", err)
			}
			res := warehouse.NewSuccessResult(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Delete", err)
			}
			err = ValidateDeleteInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Delete", err)
			}
			return nil, NewDeleteInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Warehouse", "Delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Warehouse", "Delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Warehouse", "Delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalFounderResponseBodyToWarehouseviewsFounderView builds a value of
// type *warehouseviews.FounderView from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToWarehouseviewsFounderView(v *FounderResponseBody) *warehouseviews.FounderView {
	res := &warehouseviews.FounderView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalWarehouseResponseBodyToWarehouseWarehouse builds a value of type
// *warehouse.Warehouse from a value of type *WarehouseResponseBody.
func unmarshalWarehouseResponseBodyToWarehouseWarehouse(v *WarehouseResponseBody) *warehouse.Warehouse {
	res := &warehouse.Warehouse{
		ID:      *v.ID,
		Name:    *v.Name,
		Code:    *v.Code,
		Address: *v.Address,
		Type:    *v.Type,
	}
	res.Founder = unmarshalFounderResponseBodyToWarehouseFounder(v.Founder)

	return res
}

// unmarshalFounderResponseBodyToWarehouseFounder builds a value of type
// *warehouse.Founder from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToWarehouseFounder(v *FounderResponseBody) *warehouse.Founder {
	res := &warehouse.Founder{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}
