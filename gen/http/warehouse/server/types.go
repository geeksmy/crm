// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Warehouse HTTP server types
//
// Command:
// $ goa gen crm/design

package server

import (
	warehouse "crm/gen/warehouse"
	warehouseviews "crm/gen/warehouse/views"

	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "Warehouse" service "Update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 仓库名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 仓库编码
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// 仓库地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 仓库状态
	Type *int `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// CreateRequestBody is the type of the "Warehouse" service "Create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// 仓库名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 仓库编码
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// 仓库地址
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 仓库状态
	Type *int `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// 创建人ID
	FounderID *string `form:"founder_id,omitempty" json:"founder_id,omitempty" xml:"founder_id,omitempty"`
}

// DeleteRequestBody is the type of the "Warehouse" service "Delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	Ids []string `form:"ids,omitempty" json:"ids,omitempty" xml:"ids,omitempty"`
}

// GetResponseBody is the type of the "Warehouse" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 仓库名
	Name string `form:"name" json:"name" xml:"name"`
	// 仓库编码
	Code string `form:"code" json:"code" xml:"code"`
	// 仓库地址
	Address string `form:"address" json:"address" xml:"address"`
	// 仓库状态
	Type int `form:"type" json:"type" xml:"type"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// ListResponseBody is the type of the "Warehouse" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Items []*WarehouseResponseBody `form:"items" json:"items" xml:"items"`
	// 下一页游标
	NextCursor int `form:"nextCursor" json:"nextCursor" xml:"nextCursor"`
	// 总记录数
	Total int `form:"total" json:"total" xml:"total"`
}

// UpdateResponseBody is the type of the "Warehouse" service "Update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 仓库名
	Name string `form:"name" json:"name" xml:"name"`
	// 仓库编码
	Code string `form:"code" json:"code" xml:"code"`
	// 仓库地址
	Address string `form:"address" json:"address" xml:"address"`
	// 仓库状态
	Type int `form:"type" json:"type" xml:"type"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// CreateResponseBody is the type of the "Warehouse" service "Create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 仓库名
	Name string `form:"name" json:"name" xml:"name"`
	// 仓库编码
	Code string `form:"code" json:"code" xml:"code"`
	// 仓库地址
	Address string `form:"address" json:"address" xml:"address"`
	// 仓库状态
	Type int `form:"type" json:"type" xml:"type"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// DeleteResponseBody is the type of the "Warehouse" service "Delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// success
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// GetInternalServerErrorResponseBody is the type of the "Warehouse" service
// "Get" endpoint HTTP response body for the "internal_server_error" error.
type GetInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetBadRequestResponseBody is the type of the "Warehouse" service "Get"
// endpoint HTTP response body for the "bad_request" error.
type GetBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListInternalServerErrorResponseBody is the type of the "Warehouse" service
// "List" endpoint HTTP response body for the "internal_server_error" error.
type ListInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListBadRequestResponseBody is the type of the "Warehouse" service "List"
// endpoint HTTP response body for the "bad_request" error.
type ListBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInternalServerErrorResponseBody is the type of the "Warehouse" service
// "Update" endpoint HTTP response body for the "internal_server_error" error.
type UpdateInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateBadRequestResponseBody is the type of the "Warehouse" service "Update"
// endpoint HTTP response body for the "bad_request" error.
type UpdateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateInternalServerErrorResponseBody is the type of the "Warehouse" service
// "Create" endpoint HTTP response body for the "internal_server_error" error.
type CreateInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateBadRequestResponseBody is the type of the "Warehouse" service "Create"
// endpoint HTTP response body for the "bad_request" error.
type CreateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteInternalServerErrorResponseBody is the type of the "Warehouse" service
// "Delete" endpoint HTTP response body for the "internal_server_error" error.
type DeleteInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteBadRequestResponseBody is the type of the "Warehouse" service "Delete"
// endpoint HTTP response body for the "bad_request" error.
type DeleteBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// FounderResponseBody is used to define fields on response body types.
type FounderResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 姓名
	Name string `form:"name" json:"name" xml:"name"`
}

// WarehouseResponseBody is used to define fields on response body types.
type WarehouseResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 仓库名
	Name string `form:"name" json:"name" xml:"name"`
	// 仓库编码
	Code string `form:"code" json:"code" xml:"code"`
	// 仓库地址
	Address string `form:"address" json:"address" xml:"address"`
	// 仓库状态
	Type int `form:"type" json:"type" xml:"type"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "Get" endpoint of the "Warehouse" service.
func NewGetResponseBody(res *warehouseviews.WarehouseView) *GetResponseBody {
	body := &GetResponseBody{
		ID:      *res.ID,
		Name:    *res.Name,
		Code:    *res.Code,
		Address: *res.Address,
		Type:    *res.Type,
	}
	if res.Founder != nil {
		body.Founder = marshalWarehouseviewsFounderViewToFounderResponseBody(res.Founder)
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "Warehouse" service.
func NewListResponseBody(res *warehouse.ListResult) *ListResponseBody {
	body := &ListResponseBody{
		NextCursor: res.NextCursor,
		Total:      res.Total,
	}
	if res.Items != nil {
		body.Items = make([]*WarehouseResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalWarehouseWarehouseToWarehouseResponseBody(val)
		}
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "Update" endpoint of the "Warehouse" service.
func NewUpdateResponseBody(res *warehouseviews.WarehouseView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:      *res.ID,
		Name:    *res.Name,
		Code:    *res.Code,
		Address: *res.Address,
		Type:    *res.Type,
	}
	if res.Founder != nil {
		body.Founder = marshalWarehouseviewsFounderViewToFounderResponseBody(res.Founder)
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "Create" endpoint of the "Warehouse" service.
func NewCreateResponseBody(res *warehouseviews.WarehouseView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:      *res.ID,
		Name:    *res.Name,
		Code:    *res.Code,
		Address: *res.Address,
		Type:    *res.Type,
	}
	if res.Founder != nil {
		body.Founder = marshalWarehouseviewsFounderViewToFounderResponseBody(res.Founder)
	}
	return body
}

// NewDeleteResponseBody builds the HTTP response body from the result of the
// "Delete" endpoint of the "Warehouse" service.
func NewDeleteResponseBody(res *warehouseviews.SuccessResultView) *DeleteResponseBody {
	body := &DeleteResponseBody{
		OK: *res.OK,
	}
	return body
}

// NewGetInternalServerErrorResponseBody builds the HTTP response body from the
// result of the "Get" endpoint of the "Warehouse" service.
func NewGetInternalServerErrorResponseBody(res *goa.ServiceError) *GetInternalServerErrorResponseBody {
	body := &GetInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetBadRequestResponseBody builds the HTTP response body from the result
// of the "Get" endpoint of the "Warehouse" service.
func NewGetBadRequestResponseBody(res *goa.ServiceError) *GetBadRequestResponseBody {
	body := &GetBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "List" endpoint of the "Warehouse" service.
func NewListInternalServerErrorResponseBody(res *goa.ServiceError) *ListInternalServerErrorResponseBody {
	body := &ListInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListBadRequestResponseBody builds the HTTP response body from the result
// of the "List" endpoint of the "Warehouse" service.
func NewListBadRequestResponseBody(res *goa.ServiceError) *ListBadRequestResponseBody {
	body := &ListBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "Update" endpoint of the "Warehouse" service.
func NewUpdateInternalServerErrorResponseBody(res *goa.ServiceError) *UpdateInternalServerErrorResponseBody {
	body := &UpdateInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "Update" endpoint of the "Warehouse" service.
func NewUpdateBadRequestResponseBody(res *goa.ServiceError) *UpdateBadRequestResponseBody {
	body := &UpdateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "Create" endpoint of the "Warehouse" service.
func NewCreateInternalServerErrorResponseBody(res *goa.ServiceError) *CreateInternalServerErrorResponseBody {
	body := &CreateInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "Create" endpoint of the "Warehouse" service.
func NewCreateBadRequestResponseBody(res *goa.ServiceError) *CreateBadRequestResponseBody {
	body := &CreateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteInternalServerErrorResponseBody builds the HTTP response body from
// the result of the "Delete" endpoint of the "Warehouse" service.
func NewDeleteInternalServerErrorResponseBody(res *goa.ServiceError) *DeleteInternalServerErrorResponseBody {
	body := &DeleteInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteBadRequestResponseBody builds the HTTP response body from the
// result of the "Delete" endpoint of the "Warehouse" service.
func NewDeleteBadRequestResponseBody(res *goa.ServiceError) *DeleteBadRequestResponseBody {
	body := &DeleteBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetPayload builds a Warehouse service Get endpoint payload.
func NewGetPayload(id string, token string) *warehouse.GetPayload {
	v := &warehouse.GetPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewListPayload builds a Warehouse service List endpoint payload.
func NewListPayload(token string) *warehouse.ListPayload {
	v := &warehouse.ListPayload{}
	v.Token = token

	return v
}

// NewUpdatePayload builds a Warehouse service Update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, token string) *warehouse.UpdatePayload {
	v := &warehouse.UpdatePayload{
		ID:      *body.ID,
		Name:    *body.Name,
		Code:    *body.Code,
		Address: *body.Address,
		Type:    *body.Type,
	}
	v.Token = token

	return v
}

// NewCreatePayload builds a Warehouse service Create endpoint payload.
func NewCreatePayload(body *CreateRequestBody, token string) *warehouse.CreatePayload {
	v := &warehouse.CreatePayload{
		Name:      *body.Name,
		Code:      *body.Code,
		Address:   *body.Address,
		Type:      *body.Type,
		FounderID: *body.FounderID,
	}
	v.Token = token

	return v
}

// NewDeletePayload builds a Warehouse service Delete endpoint payload.
func NewDeletePayload(body *DeleteRequestBody, token string) *warehouse.DeletePayload {
	v := &warehouse.DeletePayload{}
	v.Ids = make([]string, len(body.Ids))
	for i, val := range body.Ids {
		v.Ids[i] = val
	}
	v.Token = token

	return v
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Type != nil {
		if !(*body.Type == 1 || *body.Type == 2 || *body.Type == 3) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []interface{}{1, 2, 3}))
		}
	}
	return
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.FounderID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("founder_id", "body"))
	}
	if body.Type != nil {
		if !(*body.Type == 1 || *body.Type == 2 || *body.Type == 3) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []interface{}{1, 2, 3}))
		}
	}
	return
}

// ValidateDeleteRequestBody runs the validations defined on DeleteRequestBody
func ValidateDeleteRequestBody(body *DeleteRequestBody) (err error) {
	if body.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ids", "body"))
	}
	if len(body.Ids) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.ids", body.Ids, len(body.Ids), 100, false))
	}
	return
}
