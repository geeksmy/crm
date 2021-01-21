// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Procurement HTTP server types
//
// Command:
// $ goa gen crm/design

package server

import (
	procurement "crm/gen/procurement"
	procurementviews "crm/gen/procurement/views"

	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "Procurement" service "Update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 采购编码
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// 采购金额
	Money *int `form:"money,omitempty" json:"money,omitempty" xml:"money,omitempty"`
	// 采购还是退货
	IsSalesReturn *bool `form:"is_sales_return,omitempty" json:"is_sales_return,omitempty" xml:"is_sales_return,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 负责人ID
	HeadID *string `form:"head_id,omitempty" json:"head_id,omitempty" xml:"head_id,omitempty"`
}

// CreateRequestBody is the type of the "Procurement" service "Create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// 供应商ID
	SupplierID *string `form:"supplier_id,omitempty" json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 采购编码
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// 采购金额
	Money *int `form:"money,omitempty" json:"money,omitempty" xml:"money,omitempty"`
	// 采购还是退货
	IsSalesReturn *bool `form:"is_sales_return,omitempty" json:"is_sales_return,omitempty" xml:"is_sales_return,omitempty"`
	// 备注
	Note *string `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// 负责人ID
	HeadID *string `form:"head_id,omitempty" json:"head_id,omitempty" xml:"head_id,omitempty"`
	// 创建人ID
	FounderID *string `form:"founder_id,omitempty" json:"founder_id,omitempty" xml:"founder_id,omitempty"`
}

// DeleteRequestBody is the type of the "Procurement" service "Delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	Ids []string `form:"ids,omitempty" json:"ids,omitempty" xml:"ids,omitempty"`
}

// GetResponseBody is the type of the "Procurement" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 供应商
	Supplier *SupplierResponseBody `form:"supplier" json:"supplier" xml:"supplier"`
	// 采购编码
	Code string `form:"code" json:"code" xml:"code"`
	// 采购金额
	Money int `form:"money" json:"money" xml:"money"`
	// 采购还是退货
	IsSalesReturn bool `form:"is_sales_return" json:"is_sales_return" xml:"is_sales_return"`
	// 产品备注
	Note string `form:"note" json:"note" xml:"note"`
	// 负责人
	Head *HeadResponseBody `form:"head" json:"head" xml:"head"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// ListResponseBody is the type of the "Procurement" service "List" endpoint
// HTTP response body.
type ListResponseBody struct {
	Items []*ProcurementResponseBody `form:"items" json:"items" xml:"items"`
	// 下一页游标
	NextCursor int `form:"nextCursor" json:"nextCursor" xml:"nextCursor"`
	// 总记录数
	Total int `form:"total" json:"total" xml:"total"`
}

// UpdateResponseBody is the type of the "Procurement" service "Update"
// endpoint HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 供应商
	Supplier *SupplierResponseBody `form:"supplier" json:"supplier" xml:"supplier"`
	// 采购编码
	Code string `form:"code" json:"code" xml:"code"`
	// 采购金额
	Money int `form:"money" json:"money" xml:"money"`
	// 采购还是退货
	IsSalesReturn bool `form:"is_sales_return" json:"is_sales_return" xml:"is_sales_return"`
	// 产品备注
	Note string `form:"note" json:"note" xml:"note"`
	// 负责人
	Head *HeadResponseBody `form:"head" json:"head" xml:"head"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// CreateResponseBody is the type of the "Procurement" service "Create"
// endpoint HTTP response body.
type CreateResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 供应商
	Supplier *SupplierResponseBody `form:"supplier" json:"supplier" xml:"supplier"`
	// 采购编码
	Code string `form:"code" json:"code" xml:"code"`
	// 采购金额
	Money int `form:"money" json:"money" xml:"money"`
	// 采购还是退货
	IsSalesReturn bool `form:"is_sales_return" json:"is_sales_return" xml:"is_sales_return"`
	// 产品备注
	Note string `form:"note" json:"note" xml:"note"`
	// 负责人
	Head *HeadResponseBody `form:"head" json:"head" xml:"head"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// DeleteResponseBody is the type of the "Procurement" service "Delete"
// endpoint HTTP response body.
type DeleteResponseBody struct {
	// success
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// GetInternalServerErrorResponseBody is the type of the "Procurement" service
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

// GetBadRequestResponseBody is the type of the "Procurement" service "Get"
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

// ListInternalServerErrorResponseBody is the type of the "Procurement" service
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

// ListBadRequestResponseBody is the type of the "Procurement" service "List"
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

// UpdateInternalServerErrorResponseBody is the type of the "Procurement"
// service "Update" endpoint HTTP response body for the "internal_server_error"
// error.
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

// UpdateBadRequestResponseBody is the type of the "Procurement" service
// "Update" endpoint HTTP response body for the "bad_request" error.
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

// CreateInternalServerErrorResponseBody is the type of the "Procurement"
// service "Create" endpoint HTTP response body for the "internal_server_error"
// error.
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

// CreateBadRequestResponseBody is the type of the "Procurement" service
// "Create" endpoint HTTP response body for the "bad_request" error.
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

// DeleteInternalServerErrorResponseBody is the type of the "Procurement"
// service "Delete" endpoint HTTP response body for the "internal_server_error"
// error.
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

// DeleteBadRequestResponseBody is the type of the "Procurement" service
// "Delete" endpoint HTTP response body for the "bad_request" error.
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

// SupplierResponseBody is used to define fields on response body types.
type SupplierResponseBody struct {
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
	// 负责人
	Head *HeadResponseBody `form:"head" json:"head" xml:"head"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// HeadResponseBody is used to define fields on response body types.
type HeadResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 姓名
	Name string `form:"name" json:"name" xml:"name"`
}

// FounderResponseBody is used to define fields on response body types.
type FounderResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 姓名
	Name string `form:"name" json:"name" xml:"name"`
}

// ProcurementResponseBody is used to define fields on response body types.
type ProcurementResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
	// 供应商
	Supplier *SupplierResponseBody `form:"supplier" json:"supplier" xml:"supplier"`
	// 采购编码
	Code string `form:"code" json:"code" xml:"code"`
	// 采购金额
	Money int `form:"money" json:"money" xml:"money"`
	// 采购还是退货
	IsSalesReturn bool `form:"is_sales_return" json:"is_sales_return" xml:"is_sales_return"`
	// 产品备注
	Note string `form:"note" json:"note" xml:"note"`
	// 负责人
	Head *HeadResponseBody `form:"head" json:"head" xml:"head"`
	// 创建人
	Founder *FounderResponseBody `form:"founder" json:"founder" xml:"founder"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "Get" endpoint of the "Procurement" service.
func NewGetResponseBody(res *procurementviews.ProcurementView) *GetResponseBody {
	body := &GetResponseBody{
		ID:            *res.ID,
		Code:          *res.Code,
		Money:         *res.Money,
		IsSalesReturn: *res.IsSalesReturn,
		Note:          *res.Note,
	}
	if res.Supplier != nil {
		body.Supplier = marshalProcurementviewsSupplierViewToSupplierResponseBody(res.Supplier)
	}
	if res.Head != nil {
		body.Head = marshalProcurementviewsHeadViewToHeadResponseBody(res.Head)
	}
	if res.Founder != nil {
		body.Founder = marshalProcurementviewsFounderViewToFounderResponseBody(res.Founder)
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "Procurement" service.
func NewListResponseBody(res *procurement.ListResult) *ListResponseBody {
	body := &ListResponseBody{
		NextCursor: res.NextCursor,
		Total:      res.Total,
	}
	if res.Items != nil {
		body.Items = make([]*ProcurementResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalProcurementProcurementToProcurementResponseBody(val)
		}
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "Update" endpoint of the "Procurement" service.
func NewUpdateResponseBody(res *procurementviews.ProcurementView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:            *res.ID,
		Code:          *res.Code,
		Money:         *res.Money,
		IsSalesReturn: *res.IsSalesReturn,
		Note:          *res.Note,
	}
	if res.Supplier != nil {
		body.Supplier = marshalProcurementviewsSupplierViewToSupplierResponseBody(res.Supplier)
	}
	if res.Head != nil {
		body.Head = marshalProcurementviewsHeadViewToHeadResponseBody(res.Head)
	}
	if res.Founder != nil {
		body.Founder = marshalProcurementviewsFounderViewToFounderResponseBody(res.Founder)
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "Create" endpoint of the "Procurement" service.
func NewCreateResponseBody(res *procurementviews.ProcurementView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:            *res.ID,
		Code:          *res.Code,
		Money:         *res.Money,
		IsSalesReturn: *res.IsSalesReturn,
		Note:          *res.Note,
	}
	if res.Supplier != nil {
		body.Supplier = marshalProcurementviewsSupplierViewToSupplierResponseBody(res.Supplier)
	}
	if res.Head != nil {
		body.Head = marshalProcurementviewsHeadViewToHeadResponseBody(res.Head)
	}
	if res.Founder != nil {
		body.Founder = marshalProcurementviewsFounderViewToFounderResponseBody(res.Founder)
	}
	return body
}

// NewDeleteResponseBody builds the HTTP response body from the result of the
// "Delete" endpoint of the "Procurement" service.
func NewDeleteResponseBody(res *procurementviews.SuccessResultView) *DeleteResponseBody {
	body := &DeleteResponseBody{
		OK: *res.OK,
	}
	return body
}

// NewGetInternalServerErrorResponseBody builds the HTTP response body from the
// result of the "Get" endpoint of the "Procurement" service.
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
// of the "Get" endpoint of the "Procurement" service.
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
// the result of the "List" endpoint of the "Procurement" service.
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
// of the "List" endpoint of the "Procurement" service.
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
// the result of the "Update" endpoint of the "Procurement" service.
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
// result of the "Update" endpoint of the "Procurement" service.
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
// the result of the "Create" endpoint of the "Procurement" service.
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
// result of the "Create" endpoint of the "Procurement" service.
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
// the result of the "Delete" endpoint of the "Procurement" service.
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
// result of the "Delete" endpoint of the "Procurement" service.
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

// NewGetPayload builds a Procurement service Get endpoint payload.
func NewGetPayload(id string, token string) *procurement.GetPayload {
	v := &procurement.GetPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewListPayload builds a Procurement service List endpoint payload.
func NewListPayload(token string) *procurement.ListPayload {
	v := &procurement.ListPayload{}
	v.Token = token

	return v
}

// NewUpdatePayload builds a Procurement service Update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, token string) *procurement.UpdatePayload {
	v := &procurement.UpdatePayload{
		ID:            *body.ID,
		Code:          *body.Code,
		Money:         *body.Money,
		IsSalesReturn: *body.IsSalesReturn,
		Note:          *body.Note,
		HeadID:        *body.HeadID,
	}
	v.Token = token

	return v
}

// NewCreatePayload builds a Procurement service Create endpoint payload.
func NewCreatePayload(body *CreateRequestBody, token string) *procurement.CreatePayload {
	v := &procurement.CreatePayload{
		SupplierID:    *body.SupplierID,
		Code:          *body.Code,
		Money:         *body.Money,
		IsSalesReturn: *body.IsSalesReturn,
		Note:          *body.Note,
		HeadID:        *body.HeadID,
		FounderID:     *body.FounderID,
	}
	v.Token = token

	return v
}

// NewDeletePayload builds a Procurement service Delete endpoint payload.
func NewDeletePayload(body *DeleteRequestBody, token string) *procurement.DeletePayload {
	v := &procurement.DeletePayload{}
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
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Money == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("money", "body"))
	}
	if body.IsSalesReturn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("is_sales_return", "body"))
	}
	if body.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "body"))
	}
	if body.HeadID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("head_id", "body"))
	}
	if body.IsSalesReturn != nil {
		if !(*body.IsSalesReturn == false || *body.IsSalesReturn == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.is_sales_return", *body.IsSalesReturn, []interface{}{false, true}))
		}
	}
	return
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.SupplierID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("supplier_id", "body"))
	}
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Money == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("money", "body"))
	}
	if body.IsSalesReturn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("is_sales_return", "body"))
	}
	if body.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "body"))
	}
	if body.HeadID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("head_id", "body"))
	}
	if body.FounderID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("founder_id", "body"))
	}
	if body.IsSalesReturn != nil {
		if !(*body.IsSalesReturn == false || *body.IsSalesReturn == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.is_sales_return", *body.IsSalesReturn, []interface{}{false, true}))
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