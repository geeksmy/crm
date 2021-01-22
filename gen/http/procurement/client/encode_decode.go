// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Procurement HTTP client encoders and decoders
//
// Command:
// $ goa gen crm/design

package client

import (
	"bytes"
	"context"
	procurement "crm/gen/procurement"
	procurementviews "crm/gen/procurement/views"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "Procurement" service "Get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*procurement.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Procurement", "Get", "*procurement.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetProcurementPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Procurement", "Get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the Procurement Get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*procurement.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("Procurement", "Get", "*procurement.GetPayload", v)
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

// DecodeGetResponse returns a decoder for responses returned by the
// Procurement Get endpoint. restoreBody controls whether the response body
// should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("Procurement", "Get", err)
			}
			p := NewGetProcurementOK(&body)
			view := "default"
			vres := &procurementviews.Procurement{Projected: p, View: view}
			if err = procurementviews.ValidateProcurement(vres); err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Get", err)
			}
			res := procurement.NewProcurement(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Get", err)
			}
			err = ValidateGetInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Get", err)
			}
			return nil, NewGetInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Get", err)
			}
			return nil, NewGetBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Procurement", "Get", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "Procurement" service "List" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListProcurementPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Procurement", "List", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the Procurement
// List server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*procurement.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("Procurement", "List", "*procurement.ListPayload", v)
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

// DecodeListResponse returns a decoder for responses returned by the
// Procurement List endpoint. restoreBody controls whether the response body
// should be restored after having been read.
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
				return nil, goahttp.ErrDecodingError("Procurement", "List", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "List", err)
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
				return nil, goahttp.ErrDecodingError("Procurement", "List", err)
			}
			err = ValidateListInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "List", err)
			}
			return nil, NewListInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "List", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "List", err)
			}
			return nil, NewListBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Procurement", "List", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "Procurement" service "Update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateProcurementPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Procurement", "Update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the Procurement
// Update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*procurement.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Procurement", "Update", "*procurement.UpdatePayload", v)
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
			return goahttp.ErrEncodingError("Procurement", "Update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// Procurement Update endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Procurement", "Update", err)
			}
			p := NewUpdateProcurementOK(&body)
			view := "default"
			vres := &procurementviews.Procurement{Projected: p, View: view}
			if err = procurementviews.ValidateProcurement(vres); err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Update", err)
			}
			res := procurement.NewProcurement(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Update", err)
			}
			err = ValidateUpdateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Update", err)
			}
			return nil, NewUpdateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Procurement", "Update", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "Procurement" service "Create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateProcurementPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Procurement", "Create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the Procurement
// Create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*procurement.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("Procurement", "Create", "*procurement.CreatePayload", v)
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
			return goahttp.ErrEncodingError("Procurement", "Create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// Procurement Create endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Procurement", "Create", err)
			}
			p := NewCreateProcurementOK(&body)
			view := "default"
			vres := &procurementviews.Procurement{Projected: p, View: view}
			if err = procurementviews.ValidateProcurement(vres); err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Create", err)
			}
			res := procurement.NewProcurement(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Create", err)
			}
			err = ValidateCreateInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Create", err)
			}
			return nil, NewCreateInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Create", err)
			}
			return nil, NewCreateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Procurement", "Create", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "Procurement" service "Delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteProcurementPath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Procurement", "Delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the Procurement
// Delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*procurement.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("Procurement", "Delete", "*procurement.DeletePayload", v)
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
			return goahttp.ErrEncodingError("Procurement", "Delete", err)
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// Procurement Delete endpoint. restoreBody controls whether the response body
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
				return nil, goahttp.ErrDecodingError("Procurement", "Delete", err)
			}
			p := NewDeleteSuccessResultOK(&body)
			view := "default"
			vres := &procurementviews.SuccessResult{Projected: p, View: view}
			if err = procurementviews.ValidateSuccessResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Delete", err)
			}
			res := procurement.NewSuccessResult(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Delete", err)
			}
			err = ValidateDeleteInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Delete", err)
			}
			return nil, NewDeleteInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Procurement", "Delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Procurement", "Delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Procurement", "Delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSupplierResponseBodyToProcurementviewsSupplierView builds a value
// of type *procurementviews.SupplierView from a value of type
// *SupplierResponseBody.
func unmarshalSupplierResponseBodyToProcurementviewsSupplierView(v *SupplierResponseBody) *procurementviews.SupplierView {
	res := &procurementviews.SupplierView{
		ID:             v.ID,
		Name:           v.Name,
		Level:          v.Level,
		ContactName:    v.ContactName,
		ContactPhone:   v.ContactPhone,
		ContactAddress: v.ContactAddress,
		Note:           v.Note,
	}
	res.Head = unmarshalHeadResponseBodyToProcurementviewsHeadView(v.Head)
	res.Founder = unmarshalFounderResponseBodyToProcurementviewsFounderView(v.Founder)

	return res
}

// unmarshalHeadResponseBodyToProcurementviewsHeadView builds a value of type
// *procurementviews.HeadView from a value of type *HeadResponseBody.
func unmarshalHeadResponseBodyToProcurementviewsHeadView(v *HeadResponseBody) *procurementviews.HeadView {
	res := &procurementviews.HeadView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalFounderResponseBodyToProcurementviewsFounderView builds a value of
// type *procurementviews.FounderView from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToProcurementviewsFounderView(v *FounderResponseBody) *procurementviews.FounderView {
	res := &procurementviews.FounderView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalProcurementResponseBodyToProcurementProcurement builds a value of
// type *procurement.Procurement from a value of type *ProcurementResponseBody.
func unmarshalProcurementResponseBodyToProcurementProcurement(v *ProcurementResponseBody) *procurement.Procurement {
	res := &procurement.Procurement{
		ID:            *v.ID,
		Code:          *v.Code,
		Money:         *v.Money,
		IsSalesReturn: *v.IsSalesReturn,
		Note:          *v.Note,
	}
	res.Supplier = unmarshalSupplierResponseBodyToProcurementSupplier(v.Supplier)
	res.Head = unmarshalHeadResponseBodyToProcurementHead(v.Head)
	res.Founder = unmarshalFounderResponseBodyToProcurementFounder(v.Founder)

	return res
}

// unmarshalSupplierResponseBodyToProcurementSupplier builds a value of type
// *procurement.Supplier from a value of type *SupplierResponseBody.
func unmarshalSupplierResponseBodyToProcurementSupplier(v *SupplierResponseBody) *procurement.Supplier {
	res := &procurement.Supplier{
		ID:             *v.ID,
		Name:           *v.Name,
		Level:          *v.Level,
		ContactName:    *v.ContactName,
		ContactPhone:   *v.ContactPhone,
		ContactAddress: *v.ContactAddress,
		Note:           *v.Note,
	}
	res.Head = unmarshalHeadResponseBodyToProcurementHead(v.Head)
	res.Founder = unmarshalFounderResponseBodyToProcurementFounder(v.Founder)

	return res
}

// unmarshalHeadResponseBodyToProcurementHead builds a value of type
// *procurement.Head from a value of type *HeadResponseBody.
func unmarshalHeadResponseBodyToProcurementHead(v *HeadResponseBody) *procurement.Head {
	res := &procurement.Head{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalFounderResponseBodyToProcurementFounder builds a value of type
// *procurement.Founder from a value of type *FounderResponseBody.
func unmarshalFounderResponseBodyToProcurementFounder(v *FounderResponseBody) *procurement.Founder {
	res := &procurement.Founder{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}
