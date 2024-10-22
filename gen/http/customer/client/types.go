// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Customer HTTP client types
//
// Command:
// $ goa gen crm/design

package client

import (
	customer "crm/gen/customer"
	customerviews "crm/gen/customer/views"

	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "Customer" service "Update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 客户姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 客户来源
	Src *int `form:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	// 客户手机
	Mobile *string `form:"mobile,omitempty" json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 客户网址
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// 客户邮箱
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 客户行业
	Industry *int `form:"industry,omitempty" json:"industry,omitempty" xml:"industry,omitempty"`
	// 客户等级
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 客户地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// CreateRequestBody is the type of the "Customer" service "Create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// 客户姓名
	Name string `form:"name" json:"name" xml:"name"`
	// 客户来源
	Src int `form:"src" json:"src" xml:"src"`
	// 客户手机
	Mobile string `form:"mobile" json:"mobile" xml:"mobile"`
	// 客户网址
	URL string `form:"url" json:"url" xml:"url"`
	// 客户邮箱
	Email string `form:"email" json:"email" xml:"email"`
	// 客户行业
	Industry int `form:"industry" json:"industry" xml:"industry"`
	// 客户等级
	Level int `form:"level" json:"level" xml:"level"`
	// 备注
	Note string `form:"note" json:"note" xml:"note"`
	// 客户地址
	Address string `form:"address" json:"address" xml:"address"`
}

// DeleteRequestBody is the type of the "Customer" service "Delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	Ids []string `form:"ids" json:"ids" xml:"ids"`
}

// GetResponseBody is the type of the "Customer" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 客户姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 客户来源
	Src *int `form:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	// 客户手机
	Mobile *string `form:"mobile,omitempty" json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 客户网址
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// 客户邮箱
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 客户行业
	Industry *int `form:"industry,omitempty" json:"industry,omitempty" xml:"industry,omitempty"`
	// 客户等级
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 客户地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// ListResponseBody is the type of the "Customer" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Items []*CustomerResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// 下一页游标
	NextCursor *int `form:"nextCursor,omitempty" json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 总记录数
	Total *int `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
}

// UpdateResponseBody is the type of the "Customer" service "Update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 客户姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 客户来源
	Src *int `form:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	// 客户手机
	Mobile *string `form:"mobile,omitempty" json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 客户网址
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// 客户邮箱
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 客户行业
	Industry *int `form:"industry,omitempty" json:"industry,omitempty" xml:"industry,omitempty"`
	// 客户等级
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 客户地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// CreateResponseBody is the type of the "Customer" service "Create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 客户姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 客户来源
	Src *int `form:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	// 客户手机
	Mobile *string `form:"mobile,omitempty" json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 客户网址
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// 客户邮箱
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 客户行业
	Industry *int `form:"industry,omitempty" json:"industry,omitempty" xml:"industry,omitempty"`
	// 客户等级
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 客户地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// DeleteResponseBody is the type of the "Customer" service "Delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// success
	OK *bool `form:"ok,omitempty" json:"ok,omitempty" xml:"ok,omitempty"`
}

// GetInternalServerErrorResponseBody is the type of the "Customer" service
// "Get" endpoint HTTP response body for the "internal_server_error" error.
type GetInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// GetBadRequestResponseBody is the type of the "Customer" service "Get"
// endpoint HTTP response body for the "bad_request" error.
type GetBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListInternalServerErrorResponseBody is the type of the "Customer" service
// "List" endpoint HTTP response body for the "internal_server_error" error.
type ListInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListBadRequestResponseBody is the type of the "Customer" service "List"
// endpoint HTTP response body for the "bad_request" error.
type ListBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateInternalServerErrorResponseBody is the type of the "Customer" service
// "Update" endpoint HTTP response body for the "internal_server_error" error.
type UpdateInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateBadRequestResponseBody is the type of the "Customer" service "Update"
// endpoint HTTP response body for the "bad_request" error.
type UpdateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreateInternalServerErrorResponseBody is the type of the "Customer" service
// "Create" endpoint HTTP response body for the "internal_server_error" error.
type CreateInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "Customer" service "Create"
// endpoint HTTP response body for the "bad_request" error.
type CreateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteInternalServerErrorResponseBody is the type of the "Customer" service
// "Delete" endpoint HTTP response body for the "internal_server_error" error.
type DeleteInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteBadRequestResponseBody is the type of the "Customer" service "Delete"
// endpoint HTTP response body for the "bad_request" error.
type DeleteBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CustomerResponseBody is used to define fields on response body types.
type CustomerResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 客户姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 客户来源
	Src *int `form:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	// 客户手机
	Mobile *string `form:"mobile,omitempty" json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 客户网址
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// 客户邮箱
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 客户行业
	Industry *int `form:"industry,omitempty" json:"industry,omitempty" xml:"industry,omitempty"`
	// 客户等级
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 客户地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "Update" endpoint of the "Customer" service.
func NewUpdateRequestBody(p *customer.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		ID:       p.ID,
		Name:     p.Name,
		Src:      p.Src,
		Mobile:   p.Mobile,
		URL:      p.URL,
		Email:    p.Email,
		Industry: p.Industry,
		Level:    p.Level,
		Note:     p.Note,
		Address:  p.Address,
	}
	return body
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "Create" endpoint of the "Customer" service.
func NewCreateRequestBody(p *customer.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:     p.Name,
		Src:      p.Src,
		Mobile:   p.Mobile,
		URL:      p.URL,
		Email:    p.Email,
		Industry: p.Industry,
		Level:    p.Level,
		Note:     p.Note,
		Address:  p.Address,
	}
	return body
}

// NewDeleteRequestBody builds the HTTP request body from the payload of the
// "Delete" endpoint of the "Customer" service.
func NewDeleteRequestBody(p *customer.DeletePayload) *DeleteRequestBody {
	body := &DeleteRequestBody{}
	if p.Ids != nil {
		body.Ids = make([]string, len(p.Ids))
		for i, val := range p.Ids {
			body.Ids[i] = val
		}
	}
	return body
}

// NewGetCustomerOK builds a "Customer" service "Get" endpoint result from a
// HTTP "OK" response.
func NewGetCustomerOK(body *GetResponseBody) *customerviews.CustomerView {
	v := &customerviews.CustomerView{
		ID:       body.ID,
		Name:     body.Name,
		Src:      body.Src,
		Mobile:   body.Mobile,
		URL:      body.URL,
		Email:    body.Email,
		Industry: body.Industry,
		Level:    body.Level,
		Note:     body.Note,
		Address:  body.Address,
	}

	return v
}

// NewGetInternalServerError builds a Customer service Get endpoint
// internal_server_error error.
func NewGetInternalServerError(body *GetInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetBadRequest builds a Customer service Get endpoint bad_request error.
func NewGetBadRequest(body *GetBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewListResultOK builds a "Customer" service "List" endpoint result from a
// HTTP "OK" response.
func NewListResultOK(body *ListResponseBody) *customer.ListResult {
	v := &customer.ListResult{
		NextCursor: *body.NextCursor,
		Total:      *body.Total,
	}
	v.Items = make([]*customer.Customer, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalCustomerResponseBodyToCustomerCustomer(val)
	}

	return v
}

// NewListInternalServerError builds a Customer service List endpoint
// internal_server_error error.
func NewListInternalServerError(body *ListInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewListBadRequest builds a Customer service List endpoint bad_request error.
func NewListBadRequest(body *ListBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateCustomerOK builds a "Customer" service "Update" endpoint result
// from a HTTP "OK" response.
func NewUpdateCustomerOK(body *UpdateResponseBody) *customerviews.CustomerView {
	v := &customerviews.CustomerView{
		ID:       body.ID,
		Name:     body.Name,
		Src:      body.Src,
		Mobile:   body.Mobile,
		URL:      body.URL,
		Email:    body.Email,
		Industry: body.Industry,
		Level:    body.Level,
		Note:     body.Note,
		Address:  body.Address,
	}

	return v
}

// NewUpdateInternalServerError builds a Customer service Update endpoint
// internal_server_error error.
func NewUpdateInternalServerError(body *UpdateInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateBadRequest builds a Customer service Update endpoint bad_request
// error.
func NewUpdateBadRequest(body *UpdateBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateCustomerOK builds a "Customer" service "Create" endpoint result
// from a HTTP "OK" response.
func NewCreateCustomerOK(body *CreateResponseBody) *customerviews.CustomerView {
	v := &customerviews.CustomerView{
		ID:       body.ID,
		Name:     body.Name,
		Src:      body.Src,
		Mobile:   body.Mobile,
		URL:      body.URL,
		Email:    body.Email,
		Industry: body.Industry,
		Level:    body.Level,
		Note:     body.Note,
		Address:  body.Address,
	}

	return v
}

// NewCreateInternalServerError builds a Customer service Create endpoint
// internal_server_error error.
func NewCreateInternalServerError(body *CreateInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateBadRequest builds a Customer service Create endpoint bad_request
// error.
func NewCreateBadRequest(body *CreateBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteSuccessResultOK builds a "Customer" service "Delete" endpoint
// result from a HTTP "OK" response.
func NewDeleteSuccessResultOK(body *DeleteResponseBody) *customerviews.SuccessResultView {
	v := &customerviews.SuccessResultView{
		OK: body.OK,
	}

	return v
}

// NewDeleteInternalServerError builds a Customer service Delete endpoint
// internal_server_error error.
func NewDeleteInternalServerError(body *DeleteInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteBadRequest builds a Customer service Delete endpoint bad_request
// error.
func NewDeleteBadRequest(body *DeleteBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.NextCursor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("nextCursor", "body"))
	}
	if body.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "body"))
	}
	for _, e := range body.Items {
		if e != nil {
			if err2 := ValidateCustomerResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateGetInternalServerErrorResponseBody runs the validations defined on
// Get_internal_server_error_Response_Body
func ValidateGetInternalServerErrorResponseBody(body *GetInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateGetBadRequestResponseBody runs the validations defined on
// Get_bad_request_Response_Body
func ValidateGetBadRequestResponseBody(body *GetBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListInternalServerErrorResponseBody runs the validations defined on
// List_internal_server_error_Response_Body
func ValidateListInternalServerErrorResponseBody(body *ListInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListBadRequestResponseBody runs the validations defined on
// List_bad_request_Response_Body
func ValidateListBadRequestResponseBody(body *ListBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateInternalServerErrorResponseBody runs the validations defined
// on Update_internal_server_error_Response_Body
func ValidateUpdateInternalServerErrorResponseBody(body *UpdateInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateBadRequestResponseBody runs the validations defined on
// Update_bad_request_Response_Body
func ValidateUpdateBadRequestResponseBody(body *UpdateBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreateInternalServerErrorResponseBody runs the validations defined
// on Create_internal_server_error_Response_Body
func ValidateCreateInternalServerErrorResponseBody(body *CreateInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreateBadRequestResponseBody runs the validations defined on
// Create_bad_request_Response_Body
func ValidateCreateBadRequestResponseBody(body *CreateBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteInternalServerErrorResponseBody runs the validations defined
// on Delete_internal_server_error_Response_Body
func ValidateDeleteInternalServerErrorResponseBody(body *DeleteInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteBadRequestResponseBody runs the validations defined on
// Delete_bad_request_Response_Body
func ValidateDeleteBadRequestResponseBody(body *DeleteBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCustomerResponseBody runs the validations defined on
// CustomerResponseBody
func ValidateCustomerResponseBody(body *CustomerResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Src == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("src", "body"))
	}
	if body.Mobile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mobile", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Industry == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("industry", "body"))
	}
	if body.Level == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("level", "body"))
	}
	if body.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	return
}
