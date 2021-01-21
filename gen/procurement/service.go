// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Procurement service
//
// Command:
// $ goa gen crm/design

package procurement

import (
	"context"
	procurementviews "crm/gen/procurement/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// 采购服务
type Service interface {
	// 获取单个采购
	Get(context.Context, *GetPayload) (res *Procurement, err error)
	// 获取采购列表
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// 更新采购
	Update(context.Context, *UpdatePayload) (res *Procurement, err error)
	// 创建采购
	Create(context.Context, *CreatePayload) (res *Procurement, err error)
	// 删除采购
	Delete(context.Context, *DeletePayload) (res *SuccessResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Procurement"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"Get", "List", "Update", "Create", "Delete"}

// GetPayload is the payload type of the Procurement service Get method.
type GetPayload struct {
	// JWT used for authentication
	Token string
	ID    string
}

// Procurement is the result type of the Procurement service Get method.
type Procurement struct {
	// ID
	ID string
	// 供应商
	Supplier *Supplier
	// 采购编码
	Code string
	// 采购金额
	Money int
	// 采购还是退货
	IsSalesReturn bool
	// 产品备注
	Note string
	// 负责人
	Head *Head
	// 创建人
	Founder *Founder
}

// ListPayload is the payload type of the Procurement service List method.
type ListPayload struct {
	// JWT used for authentication
	Token string
}

// ListResult is the result type of the Procurement service List method.
type ListResult struct {
	Items []*Procurement
	// 下一页游标
	NextCursor int
	// 总记录数
	Total int
}

// UpdatePayload is the payload type of the Procurement service Update method.
type UpdatePayload struct {
	// JWT used for authentication
	Token string
	// ID
	ID string
	// 采购编码
	Code string
	// 采购金额
	Money int
	// 采购还是退货
	IsSalesReturn bool
	// 备注
	Note string
	// 负责人ID
	HeadID string
}

// CreatePayload is the payload type of the Procurement service Create method.
type CreatePayload struct {
	// JWT used for authentication
	Token string
	// 供应商ID
	SupplierID string
	// 采购编码
	Code string
	// 采购金额
	Money int
	// 采购还是退货
	IsSalesReturn bool
	// 备注
	Note string
	// 负责人ID
	HeadID string
	// 创建人ID
	FounderID string
}

// DeletePayload is the payload type of the Procurement service Delete method.
type DeletePayload struct {
	// JWT used for authentication
	Token string
	Ids   []string
}

// SuccessResult is the result type of the Procurement service Delete method.
type SuccessResult struct {
	// success
	OK bool
}

type Supplier struct {
	// ID
	ID string
	// 供应商名
	Name string
	// 级别
	Level int
	// 联系人姓名
	ContactName string
	// 联系人手机
	ContactPhone string
	// 联系人地址
	ContactAddress string
	// 产品备注
	Note string
	// 负责人
	Head *Head
	// 创建人
	Founder *Founder
}

type Head struct {
	// ID
	ID string
	// 姓名
	Name string
}

type Founder struct {
	// ID
	ID string
	// 姓名
	Name string
}

// MakeInternalServerError builds a goa.ServiceError from an error.
func MakeInternalServerError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal_server_error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewProcurement initializes result type Procurement from viewed result type
// Procurement.
func NewProcurement(vres *procurementviews.Procurement) *Procurement {
	return newProcurement(vres.Projected)
}

// NewViewedProcurement initializes viewed result type Procurement from result
// type Procurement using the given view.
func NewViewedProcurement(res *Procurement, view string) *procurementviews.Procurement {
	p := newProcurementView(res)
	return &procurementviews.Procurement{Projected: p, View: "default"}
}

// NewSuccessResult initializes result type SuccessResult from viewed result
// type SuccessResult.
func NewSuccessResult(vres *procurementviews.SuccessResult) *SuccessResult {
	return newSuccessResult(vres.Projected)
}

// NewViewedSuccessResult initializes viewed result type SuccessResult from
// result type SuccessResult using the given view.
func NewViewedSuccessResult(res *SuccessResult, view string) *procurementviews.SuccessResult {
	p := newSuccessResultView(res)
	return &procurementviews.SuccessResult{Projected: p, View: "default"}
}

// newProcurement converts projected type Procurement to service type
// Procurement.
func newProcurement(vres *procurementviews.ProcurementView) *Procurement {
	res := &Procurement{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Code != nil {
		res.Code = *vres.Code
	}
	if vres.Money != nil {
		res.Money = *vres.Money
	}
	if vres.IsSalesReturn != nil {
		res.IsSalesReturn = *vres.IsSalesReturn
	}
	if vres.Note != nil {
		res.Note = *vres.Note
	}
	if vres.Supplier != nil {
		res.Supplier = newSupplier(vres.Supplier)
	}
	if vres.Head != nil {
		res.Head = newHead(vres.Head)
	}
	if vres.Founder != nil {
		res.Founder = newFounder(vres.Founder)
	}
	return res
}

// newProcurementView projects result type Procurement to projected type
// ProcurementView using the "default" view.
func newProcurementView(res *Procurement) *procurementviews.ProcurementView {
	vres := &procurementviews.ProcurementView{
		ID:            &res.ID,
		Code:          &res.Code,
		Money:         &res.Money,
		IsSalesReturn: &res.IsSalesReturn,
		Note:          &res.Note,
	}
	if res.Supplier != nil {
		vres.Supplier = newSupplierView(res.Supplier)
	}
	if res.Head != nil {
		vres.Head = newHeadView(res.Head)
	}
	if res.Founder != nil {
		vres.Founder = newFounderView(res.Founder)
	}
	return vres
}

// newSupplier converts projected type Supplier to service type Supplier.
func newSupplier(vres *procurementviews.SupplierView) *Supplier {
	res := &Supplier{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Level != nil {
		res.Level = *vres.Level
	}
	if vres.ContactName != nil {
		res.ContactName = *vres.ContactName
	}
	if vres.ContactPhone != nil {
		res.ContactPhone = *vres.ContactPhone
	}
	if vres.ContactAddress != nil {
		res.ContactAddress = *vres.ContactAddress
	}
	if vres.Note != nil {
		res.Note = *vres.Note
	}
	if vres.Head != nil {
		res.Head = newHead(vres.Head)
	}
	if vres.Founder != nil {
		res.Founder = newFounder(vres.Founder)
	}
	return res
}

// newSupplierView projects result type Supplier to projected type SupplierView
// using the "default" view.
func newSupplierView(res *Supplier) *procurementviews.SupplierView {
	vres := &procurementviews.SupplierView{
		ID:             &res.ID,
		Name:           &res.Name,
		Level:          &res.Level,
		ContactName:    &res.ContactName,
		ContactPhone:   &res.ContactPhone,
		ContactAddress: &res.ContactAddress,
		Note:           &res.Note,
	}
	if res.Head != nil {
		vres.Head = newHeadView(res.Head)
	}
	if res.Founder != nil {
		vres.Founder = newFounderView(res.Founder)
	}
	return vres
}

// newHead converts projected type Head to service type Head.
func newHead(vres *procurementviews.HeadView) *Head {
	res := &Head{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newHeadView projects result type Head to projected type HeadView using the
// "default" view.
func newHeadView(res *Head) *procurementviews.HeadView {
	vres := &procurementviews.HeadView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}

// newFounder converts projected type Founder to service type Founder.
func newFounder(vres *procurementviews.FounderView) *Founder {
	res := &Founder{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newFounderView projects result type Founder to projected type FounderView
// using the "default" view.
func newFounderView(res *Founder) *procurementviews.FounderView {
	vres := &procurementviews.FounderView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}

// newSuccessResult converts projected type SuccessResult to service type
// SuccessResult.
func newSuccessResult(vres *procurementviews.SuccessResultView) *SuccessResult {
	res := &SuccessResult{}
	if vres.OK != nil {
		res.OK = *vres.OK
	}
	return res
}

// newSuccessResultView projects result type SuccessResult to projected type
// SuccessResultView using the "default" view.
func newSuccessResultView(res *SuccessResult) *procurementviews.SuccessResultView {
	vres := &procurementviews.SuccessResultView{
		OK: &res.OK,
	}
	return vres
}