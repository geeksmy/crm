// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Customer HTTP server types
//
// Command:
// $ goa gen crm/design

package server

import (
	customer "crm/gen/customer"
	customerviews "crm/gen/customer/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "Customer" service "Update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
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

// CreateRequestBody is the type of the "Customer" service "Create" endpoint
// HTTP request body.
type CreateRequestBody struct {
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

// DeleteRequestBody is the type of the "Customer" service "Delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	Ids []string `form:"ids,omitempty" json:"ids,omitempty" xml:"ids,omitempty"`
}

// GetResponseBody is the type of the "Customer" service "Get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
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

// ListResponseBody is the type of the "Customer" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Items []*CustomerResponseBody `form:"items" json:"items" xml:"items"`
	// 下一页游标
	NextCursor int `form:"nextCursor" json:"nextCursor" xml:"nextCursor"`
	// 总记录数
	Total int `form:"total" json:"total" xml:"total"`
}

// UpdateResponseBody is the type of the "Customer" service "Update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
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

// CreateResponseBody is the type of the "Customer" service "Create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
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

// DeleteResponseBody is the type of the "Customer" service "Delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// success
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// GetInternalServerErrorResponseBody is the type of the "Customer" service
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

// GetBadRequestResponseBody is the type of the "Customer" service "Get"
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

// ListInternalServerErrorResponseBody is the type of the "Customer" service
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

// ListBadRequestResponseBody is the type of the "Customer" service "List"
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

// UpdateInternalServerErrorResponseBody is the type of the "Customer" service
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

// UpdateBadRequestResponseBody is the type of the "Customer" service "Update"
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

// CreateInternalServerErrorResponseBody is the type of the "Customer" service
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

// CreateBadRequestResponseBody is the type of the "Customer" service "Create"
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

// DeleteInternalServerErrorResponseBody is the type of the "Customer" service
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

// DeleteBadRequestResponseBody is the type of the "Customer" service "Delete"
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

// CustomerResponseBody is used to define fields on response body types.
type CustomerResponseBody struct {
	// ID
	ID string `form:"id" json:"id" xml:"id"`
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

// NewGetResponseBody builds the HTTP response body from the result of the
// "Get" endpoint of the "Customer" service.
func NewGetResponseBody(res *customerviews.CustomerView) *GetResponseBody {
	body := &GetResponseBody{
		ID:       *res.ID,
		Name:     *res.Name,
		Src:      *res.Src,
		Mobile:   *res.Mobile,
		URL:      *res.URL,
		Email:    *res.Email,
		Industry: *res.Industry,
		Level:    *res.Level,
		Note:     *res.Note,
		Address:  *res.Address,
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "Customer" service.
func NewListResponseBody(res *customer.ListResult) *ListResponseBody {
	body := &ListResponseBody{
		NextCursor: res.NextCursor,
		Total:      res.Total,
	}
	if res.Items != nil {
		body.Items = make([]*CustomerResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalCustomerCustomerToCustomerResponseBody(val)
		}
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "Update" endpoint of the "Customer" service.
func NewUpdateResponseBody(res *customerviews.CustomerView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:       *res.ID,
		Name:     *res.Name,
		Src:      *res.Src,
		Mobile:   *res.Mobile,
		URL:      *res.URL,
		Email:    *res.Email,
		Industry: *res.Industry,
		Level:    *res.Level,
		Note:     *res.Note,
		Address:  *res.Address,
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "Create" endpoint of the "Customer" service.
func NewCreateResponseBody(res *customerviews.CustomerView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:       *res.ID,
		Name:     *res.Name,
		Src:      *res.Src,
		Mobile:   *res.Mobile,
		URL:      *res.URL,
		Email:    *res.Email,
		Industry: *res.Industry,
		Level:    *res.Level,
		Note:     *res.Note,
		Address:  *res.Address,
	}
	return body
}

// NewDeleteResponseBody builds the HTTP response body from the result of the
// "Delete" endpoint of the "Customer" service.
func NewDeleteResponseBody(res *customerviews.SuccessResultView) *DeleteResponseBody {
	body := &DeleteResponseBody{
		OK: *res.OK,
	}
	return body
}

// NewGetInternalServerErrorResponseBody builds the HTTP response body from the
// result of the "Get" endpoint of the "Customer" service.
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
// of the "Get" endpoint of the "Customer" service.
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
// the result of the "List" endpoint of the "Customer" service.
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
// of the "List" endpoint of the "Customer" service.
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
// the result of the "Update" endpoint of the "Customer" service.
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
// result of the "Update" endpoint of the "Customer" service.
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
// the result of the "Create" endpoint of the "Customer" service.
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
// result of the "Create" endpoint of the "Customer" service.
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
// the result of the "Delete" endpoint of the "Customer" service.
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
// result of the "Delete" endpoint of the "Customer" service.
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

// NewGetPayload builds a Customer service Get endpoint payload.
func NewGetPayload(id string, token string) *customer.GetPayload {
	v := &customer.GetPayload{}
	v.ID = id
	v.Token = token

	return v
}

// NewListPayload builds a Customer service List endpoint payload.
func NewListPayload(cursor *int, limit *int, token string) *customer.ListPayload {
	v := &customer.ListPayload{}
	v.Cursor = cursor
	v.Limit = limit
	v.Token = token

	return v
}

// NewUpdatePayload builds a Customer service Update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, token string) *customer.UpdatePayload {
	v := &customer.UpdatePayload{
		ID:       *body.ID,
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
	v.Token = token

	return v
}

// NewCreatePayload builds a Customer service Create endpoint payload.
func NewCreatePayload(body *CreateRequestBody, token string) *customer.CreatePayload {
	v := &customer.CreatePayload{
		Name:     *body.Name,
		Src:      *body.Src,
		Mobile:   *body.Mobile,
		URL:      *body.URL,
		Email:    *body.Email,
		Industry: *body.Industry,
		Level:    *body.Level,
		Note:     *body.Note,
		Address:  *body.Address,
	}
	v.Token = token

	return v
}

// NewDeletePayload builds a Customer service Delete endpoint payload.
func NewDeletePayload(body *DeleteRequestBody, token string) *customer.DeletePayload {
	v := &customer.DeletePayload{}
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
	if body.Mobile != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.mobile", *body.Mobile, goa.FormatRegexp))
	}
	if body.Mobile != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.mobile", *body.Mobile, "^1(?:3\\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\\d|9\\d)\\d{8}$"))
	}
	if body.Mobile != nil {
		if utf8.RuneCountInString(*body.Mobile) < 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", *body.Mobile, utf8.RuneCountInString(*body.Mobile), 11, true))
		}
	}
	if body.Mobile != nil {
		if utf8.RuneCountInString(*body.Mobile) > 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", *body.Mobile, utf8.RuneCountInString(*body.Mobile), 11, false))
		}
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.url", *body.URL, goa.FormatURI))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
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
	if body.Mobile != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.mobile", *body.Mobile, goa.FormatRegexp))
	}
	if body.Mobile != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.mobile", *body.Mobile, "^1(?:3\\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\\d|9\\d)\\d{8}$"))
	}
	if body.Mobile != nil {
		if utf8.RuneCountInString(*body.Mobile) < 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", *body.Mobile, utf8.RuneCountInString(*body.Mobile), 11, true))
		}
	}
	if body.Mobile != nil {
		if utf8.RuneCountInString(*body.Mobile) > 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", *body.Mobile, utf8.RuneCountInString(*body.Mobile), 11, false))
		}
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.url", *body.URL, goa.FormatURI))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
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
