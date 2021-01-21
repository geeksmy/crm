// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Supplier HTTP client types
//
// Command:
// $ goa gen crm/design

package client

import (
	supplier "crm/gen/supplier"
	supplierviews "crm/gen/supplier/views"

	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "Supplier" service "Update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 供应商名
	Name string `form:"name" json:"name" xml:"name"`
	// 级别
	Level int `form:"level" json:"level" xml:"level"`
	// 联系人姓名
	ContactName string `form:"contact_name" json:"contact_name" xml:"contact_name"`
	// 联系人手机
	ContactPhone string `form:"contact_phone" json:"contact_phone" xml:"contact_phone"`
	// 联系人地址
	ContactAddress string `form:"contact_address" json:"contact_address" xml:"contact_address"`
	// 产品备注
	Note string `form:"note" json:"note" xml:"note"`
	// 负责人ID
	HeadID string `form:"head_id" json:"head_id" xml:"head_id"`
}

// CreateRequestBody is the type of the "Supplier" service "Create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// 供应商名
	Name string `form:"name" json:"name" xml:"name"`
	// 级别
	Level int `form:"level" json:"level" xml:"level"`
	// 联系人姓名
	ContactName string `form:"contact_name" json:"contact_name" xml:"contact_name"`
	// 联系人手机
	ContactPhone string `form:"contact_phone" json:"contact_phone" xml:"contact_phone"`
	// 联系人地址
	ContactAddress string `form:"contact_address" json:"contact_address" xml:"contact_address"`
	// 产品备注
	Note string `form:"note" json:"note" xml:"note"`
	// 负责人ID
	HeadID string `form:"head_id" json:"head_id" xml:"head_id"`
	// 创建人ID
	FounderID string `form:"founder_id" json:"founder_id" xml:"founder_id"`
}

// DeleteRequestBody is the type of the "Supplier" service "Delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	Ids []string `form:"ids" json:"ids" xml:"ids"`
}

// GetResponseBody is the type of the "Supplier" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 供应商名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 级别
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 联系人姓名
	ContactName *string `form:"contact_name,omitempty" json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人手机
	ContactPhone *string `form:"contact_phone,omitempty" json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 联系人地址
	ContactAddress *string `form:"contact_address,omitempty" json:"contact_address,omitempty" xml:"contact_address,omitempty"`
	// 产品备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 负责人
	Head *HeadResponseBody `form:"head,omitempty" json:"head,omitempty" xml:"head,omitempty"`
	// 创建人
	Founder *FounderResponseBody `form:"founder,omitempty" json:"founder,omitempty" xml:"founder,omitempty"`
}

// ListResponseBody is the type of the "Supplier" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Items []*SupplierResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// 下一页游标
	NextCursor *int `form:"nextCursor,omitempty" json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 总记录数
	Total *int `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
}

// UpdateResponseBody is the type of the "Supplier" service "Update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 供应商名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 级别
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 联系人姓名
	ContactName *string `form:"contact_name,omitempty" json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人手机
	ContactPhone *string `form:"contact_phone,omitempty" json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 联系人地址
	ContactAddress *string `form:"contact_address,omitempty" json:"contact_address,omitempty" xml:"contact_address,omitempty"`
	// 产品备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 负责人
	Head *HeadResponseBody `form:"head,omitempty" json:"head,omitempty" xml:"head,omitempty"`
	// 创建人
	Founder *FounderResponseBody `form:"founder,omitempty" json:"founder,omitempty" xml:"founder,omitempty"`
}

// CreateResponseBody is the type of the "Supplier" service "Create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 供应商名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 级别
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 联系人姓名
	ContactName *string `form:"contact_name,omitempty" json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人手机
	ContactPhone *string `form:"contact_phone,omitempty" json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 联系人地址
	ContactAddress *string `form:"contact_address,omitempty" json:"contact_address,omitempty" xml:"contact_address,omitempty"`
	// 产品备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 负责人
	Head *HeadResponseBody `form:"head,omitempty" json:"head,omitempty" xml:"head,omitempty"`
	// 创建人
	Founder *FounderResponseBody `form:"founder,omitempty" json:"founder,omitempty" xml:"founder,omitempty"`
}

// DeleteResponseBody is the type of the "Supplier" service "Delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// success
	OK *bool `form:"ok,omitempty" json:"ok,omitempty" xml:"ok,omitempty"`
}

// GetInternalServerErrorResponseBody is the type of the "Supplier" service
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

// GetBadRequestResponseBody is the type of the "Supplier" service "Get"
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

// ListInternalServerErrorResponseBody is the type of the "Supplier" service
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

// ListBadRequestResponseBody is the type of the "Supplier" service "List"
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

// UpdateInternalServerErrorResponseBody is the type of the "Supplier" service
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

// UpdateBadRequestResponseBody is the type of the "Supplier" service "Update"
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

// CreateInternalServerErrorResponseBody is the type of the "Supplier" service
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

// CreateBadRequestResponseBody is the type of the "Supplier" service "Create"
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

// DeleteInternalServerErrorResponseBody is the type of the "Supplier" service
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

// DeleteBadRequestResponseBody is the type of the "Supplier" service "Delete"
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

// HeadResponseBody is used to define fields on response body types.
type HeadResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// FounderResponseBody is used to define fields on response body types.
type FounderResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 姓名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// SupplierResponseBody is used to define fields on response body types.
type SupplierResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 供应商名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 级别
	Level *int `form:"level,omitempty" json:"level,omitempty" xml:"level,omitempty"`
	// 联系人姓名
	ContactName *string `form:"contact_name,omitempty" json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人手机
	ContactPhone *string `form:"contact_phone,omitempty" json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 联系人地址
	ContactAddress *string `form:"contact_address,omitempty" json:"contact_address,omitempty" xml:"contact_address,omitempty"`
	// 产品备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 负责人
	Head *HeadResponseBody `form:"head,omitempty" json:"head,omitempty" xml:"head,omitempty"`
	// 创建人
	Founder *FounderResponseBody `form:"founder,omitempty" json:"founder,omitempty" xml:"founder,omitempty"`
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "Update" endpoint of the "Supplier" service.
func NewUpdateRequestBody(p *supplier.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		ID:             p.ID,
		Name:           p.Name,
		Level:          p.Level,
		ContactName:    p.ContactName,
		ContactPhone:   p.ContactPhone,
		ContactAddress: p.ContactAddress,
		Note:           p.Note,
		HeadID:         p.HeadID,
	}
	return body
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "Create" endpoint of the "Supplier" service.
func NewCreateRequestBody(p *supplier.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:           p.Name,
		Level:          p.Level,
		ContactName:    p.ContactName,
		ContactPhone:   p.ContactPhone,
		ContactAddress: p.ContactAddress,
		Note:           p.Note,
		HeadID:         p.HeadID,
		FounderID:      p.FounderID,
	}
	return body
}

// NewDeleteRequestBody builds the HTTP request body from the payload of the
// "Delete" endpoint of the "Supplier" service.
func NewDeleteRequestBody(p *supplier.DeletePayload) *DeleteRequestBody {
	body := &DeleteRequestBody{}
	if p.Ids != nil {
		body.Ids = make([]string, len(p.Ids))
		for i, val := range p.Ids {
			body.Ids[i] = val
		}
	}
	return body
}

// NewGetSupplierOK builds a "Supplier" service "Get" endpoint result from a
// HTTP "OK" response.
func NewGetSupplierOK(body *GetResponseBody) *supplierviews.SupplierView {
	v := &supplierviews.SupplierView{
		ID:             body.ID,
		Name:           body.Name,
		Level:          body.Level,
		ContactName:    body.ContactName,
		ContactPhone:   body.ContactPhone,
		ContactAddress: body.ContactAddress,
		Note:           body.Note,
	}
	v.Head = unmarshalHeadResponseBodyToSupplierviewsHeadView(body.Head)
	v.Founder = unmarshalFounderResponseBodyToSupplierviewsFounderView(body.Founder)

	return v
}

// NewGetInternalServerError builds a Supplier service Get endpoint
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

// NewGetBadRequest builds a Supplier service Get endpoint bad_request error.
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

// NewListResultOK builds a "Supplier" service "List" endpoint result from a
// HTTP "OK" response.
func NewListResultOK(body *ListResponseBody) *supplier.ListResult {
	v := &supplier.ListResult{
		NextCursor: *body.NextCursor,
		Total:      *body.Total,
	}
	v.Items = make([]*supplier.Supplier, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalSupplierResponseBodyToSupplierSupplier(val)
	}

	return v
}

// NewListInternalServerError builds a Supplier service List endpoint
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

// NewListBadRequest builds a Supplier service List endpoint bad_request error.
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

// NewUpdateSupplierOK builds a "Supplier" service "Update" endpoint result
// from a HTTP "OK" response.
func NewUpdateSupplierOK(body *UpdateResponseBody) *supplierviews.SupplierView {
	v := &supplierviews.SupplierView{
		ID:             body.ID,
		Name:           body.Name,
		Level:          body.Level,
		ContactName:    body.ContactName,
		ContactPhone:   body.ContactPhone,
		ContactAddress: body.ContactAddress,
		Note:           body.Note,
	}
	v.Head = unmarshalHeadResponseBodyToSupplierviewsHeadView(body.Head)
	v.Founder = unmarshalFounderResponseBodyToSupplierviewsFounderView(body.Founder)

	return v
}

// NewUpdateInternalServerError builds a Supplier service Update endpoint
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

// NewUpdateBadRequest builds a Supplier service Update endpoint bad_request
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

// NewCreateSupplierOK builds a "Supplier" service "Create" endpoint result
// from a HTTP "OK" response.
func NewCreateSupplierOK(body *CreateResponseBody) *supplierviews.SupplierView {
	v := &supplierviews.SupplierView{
		ID:             body.ID,
		Name:           body.Name,
		Level:          body.Level,
		ContactName:    body.ContactName,
		ContactPhone:   body.ContactPhone,
		ContactAddress: body.ContactAddress,
		Note:           body.Note,
	}
	v.Head = unmarshalHeadResponseBodyToSupplierviewsHeadView(body.Head)
	v.Founder = unmarshalFounderResponseBodyToSupplierviewsFounderView(body.Founder)

	return v
}

// NewCreateInternalServerError builds a Supplier service Create endpoint
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

// NewCreateBadRequest builds a Supplier service Create endpoint bad_request
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

// NewDeleteSuccessResultOK builds a "Supplier" service "Delete" endpoint
// result from a HTTP "OK" response.
func NewDeleteSuccessResultOK(body *DeleteResponseBody) *supplierviews.SuccessResultView {
	v := &supplierviews.SuccessResultView{
		OK: body.OK,
	}

	return v
}

// NewDeleteInternalServerError builds a Supplier service Delete endpoint
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

// NewDeleteBadRequest builds a Supplier service Delete endpoint bad_request
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
			if err2 := ValidateSupplierResponseBody(e); err2 != nil {
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

// ValidateHeadResponseBody runs the validations defined on HeadResponseBody
func ValidateHeadResponseBody(body *HeadResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateFounderResponseBody runs the validations defined on
// FounderResponseBody
func ValidateFounderResponseBody(body *FounderResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateSupplierResponseBody runs the validations defined on
// SupplierResponseBody
func ValidateSupplierResponseBody(body *SupplierResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Level == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("level", "body"))
	}
	if body.ContactName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contact_name", "body"))
	}
	if body.ContactPhone == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contact_phone", "body"))
	}
	if body.ContactAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contact_address", "body"))
	}
	if body.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "body"))
	}
	if body.Head == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("head", "body"))
	}
	if body.Founder == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("founder", "body"))
	}
	if body.Level != nil {
		if !(*body.Level == 1 || *body.Level == 2 || *body.Level == 3 || *body.Level == 4) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.level", *body.Level, []interface{}{1, 2, 3, 4}))
		}
	}
	if body.Head != nil {
		if err2 := ValidateHeadResponseBody(body.Head); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Founder != nil {
		if err2 := ValidateFounderResponseBody(body.Founder); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
