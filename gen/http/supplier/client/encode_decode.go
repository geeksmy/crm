// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Supplier HTTP client encoders and decoders
//
// Command:
// $ goa gen crm/design

package client

import (
	"bytes"
	"context"
	supplier "crm/gen/supplier"
	supplierviews "crm/gen/supplier/views"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "Supplier" service "Get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*supplier.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Supplier", "Get", "*supplier.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetSupplierPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Supplier", "Get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the Supplier Get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*supplier.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("Supplier", "Get", "*supplier.GetPayload", v)
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

// DecodeGetResponse returns a decoder for responses returned by the Supplier
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
				return nil, goahttp.ErrDecodingError("Supplier", "Get", err)
			}
			p := NewGetSupplierOK(&body)
			view := "default"
			vres := &supplierviews.Supplier{Projected: p, View: view}
			if err = supplierviews.ValidateSupplier(vres); err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Get", err)
			}
			res := supplier.NewSupplier(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Get", err)
			}
			err = ValidateGetInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Get", err)
			}
			return nil, NewGetInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Get", err)
			}
			return nil, NewGetBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Supplier", "Get", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "Supplier" service "List" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListSupplierPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Supplier", "List", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the Supplier List
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*supplier.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("Supplier", "List", "*supplier.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the Supplier
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
				return nil, goahttp.ErrDecodingError("Supplier", "List", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "List", err)
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
				return nil, goahttp.ErrDecodingError("Supplier", "List", err)
			}
			err = ValidateListInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "List", err)
			}
			return nil, NewListInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "List", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "List", err)
			}
			return nil, NewListBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Supplier", "List", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "Supplier" service "Update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateSupplierPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Supplier", "Update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the Supplier
// Update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*supplier.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Supplier", "Update", "*supplier.UpdatePayload", v)
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
			return goahttp.ErrEncodingError("Supplier", "Update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// Supplier Update endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Supplier", "Update", err)
			}
			p := NewUpdateSupplierOK(&body)
			view := "default"
			vres := &supplierviews.Supplier{Projected: p, View: view}
			if err = supplierviews.ValidateSupplier(vres); err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Update", err)
			}
			res := supplier.NewSupplier(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Update", err)
			}
			err = ValidateUpdateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Update", err)
			}
			return nil, NewUpdateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Supplier", "Update", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "Supplier" service "Create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateSupplierPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Supplier", "Create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the Supplier
// Create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*supplier.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Supplier", "Create", "*supplier.CreatePayload", v)
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
			return goahttp.ErrEncodingError("Supplier", "Create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// Supplier Create endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Supplier", "Create", err)
			}
			p := NewCreateSupplierOK(&body)
			view := "default"
			vres := &supplierviews.Supplier{Projected: p, View: view}
			if err = supplierviews.ValidateSupplier(vres); err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Create", err)
			}
			res := supplier.NewSupplier(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Create", err)
			}
			err = ValidateCreateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Create", err)
			}
			return nil, NewCreateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Create", err)
			}
			return nil, NewCreateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Supplier", "Create", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "Supplier" service "Delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteSupplierPath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Supplier", "Delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the Supplier
// Delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*supplier.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("Supplier", "Delete", "*supplier.DeletePayload", v)
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
			return goahttp.ErrEncodingError("Supplier", "Delete", err)
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// Supplier Delete endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Supplier", "Delete", err)
			}
			p := NewDeleteSuccessResultOK(&body)
			view := "default"
			vres := &supplierviews.SuccessResult{Projected: p, View: view}
			if err = supplierviews.ValidateSuccessResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Delete", err)
			}
			res := supplier.NewSuccessResult(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Delete", err)
			}
			err = ValidateDeleteInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Delete", err)
			}
			return nil, NewDeleteInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Supplier", "Delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Supplier", "Delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Supplier", "Delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalHeadResponseBodyToSupplierviewsHeadView builds a value of type
// *supplierviews.HeadView from a value of type *HeadResponseBody.
func unmarshalHeadResponseBodyToSupplierviewsHeadView(v *HeadResponseBody) *supplierviews.HeadView {
	res := &supplierviews.HeadView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalFounderResponseBodyToSupplierviewsFounderView builds a value of
// type *supplierviews.FounderView from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToSupplierviewsFounderView(v *FounderResponseBody) *supplierviews.FounderView {
	res := &supplierviews.FounderView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalSupplierResponseBodyToSupplierSupplier builds a value of type
// *supplier.Supplier from a value of type *SupplierResponseBody.
func unmarshalSupplierResponseBodyToSupplierSupplier(v *SupplierResponseBody) *supplier.Supplier {
	res := &supplier.Supplier{
		ID:             *v.ID,
		Name:           *v.Name,
		Level:          *v.Level,
		ContactName:    *v.ContactName,
		ContactPhone:   *v.ContactPhone,
		ContactAddress: *v.ContactAddress,
		Note:           *v.Note,
	}
	res.Head = unmarshalHeadResponseBodyToSupplierHead(v.Head)
	res.Founder = unmarshalFounderResponseBodyToSupplierFounder(v.Founder)

	return res
}

// unmarshalHeadResponseBodyToSupplierHead builds a value of type
// *supplier.Head from a value of type *HeadResponseBody.
func unmarshalHeadResponseBodyToSupplierHead(v *HeadResponseBody) *supplier.Head {
	res := &supplier.Head{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalFounderResponseBodyToSupplierFounder builds a value of type
// *supplier.Founder from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToSupplierFounder(v *FounderResponseBody) *supplier.Founder {
	res := &supplier.Founder{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}