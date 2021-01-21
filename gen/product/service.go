// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Product service
//
// Command:
// $ goa gen crm/design

package product

import (
	"context"
	productviews "crm/gen/product/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// 产品服务
type Service interface {
	// 获取单个产品
	Get(context.Context, *GetPayload) (res *Product, err error)
	// 获取产品列表
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// 更新产品
	Update(context.Context, *UpdatePayload) (res *Product, err error)
	// 创建产品
	Create(context.Context, *CreatePayload) (res *Product, err error)
	// 删除产品
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
const ServiceName = "Product"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"Get", "List", "Update", "Create", "Delete"}

// GetPayload is the payload type of the Product service Get method.
type GetPayload struct {
	// JWT used for authentication
	Token string
	ID    string
}

// Product is the result type of the Product service Get method.
type Product struct {
	// ID
	ID string
	// 产品名
	Name string
	// 产品单位
	Unit int
	// 成本价
	CostPrice int
	// 市场价
	MarketPrice int
	// 产品备注
	Note string
	// 产品图片
	Image string
	// 产品编码
	Code string
	// 产品规格
	Size string
	// 产品类型
	Type int
	// 是否上架
	IsShelves bool
	// 创建人
	Founder *Founder
}

// ListPayload is the payload type of the Product service List method.
type ListPayload struct {
	// JWT used for authentication
	Token string
}

// ListResult is the result type of the Product service List method.
type ListResult struct {
	Items []*Product
	// 下一页游标
	NextCursor int
	// 总记录数
	Total int
}

// UpdatePayload is the payload type of the Product service Update method.
type UpdatePayload struct {
	// JWT used for authentication
	Token string
	// 用户ID
	ID string
	// 产品名
	Name string
	// 产品单位
	Unit int
	// 成本价
	CostPrice int
	// 市场价
	MarketPrice int
	// 产品备注
	Note string
	// 产品图片
	Image string
	// 产品规格
	Size string
	// 产品类型
	Type int
	// 是否上架
	IsShelves bool
}

// CreatePayload is the payload type of the Product service Create method.
type CreatePayload struct {
	// JWT used for authentication
	Token string
	// 产品名
	Name string
	// 产品单位
	Unit int
	// 成本价
	CostPrice int
	// 市场价
	MarketPrice int
	// 产品备注
	Note string
	// 产品图片
	Image string
	// 产品编码
	Code string
	// 产品规格
	Size string
	// 产品类型
	Type int
	// 是否上架
	IsShelves bool
	// 创建人ID
	FounderID string
}

// DeletePayload is the payload type of the Product service Delete method.
type DeletePayload struct {
	// JWT used for authentication
	Token string
	Ids   []string
}

// SuccessResult is the result type of the Product service Delete method.
type SuccessResult struct {
	// success
	OK bool
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

// NewProduct initializes result type Product from viewed result type Product.
func NewProduct(vres *productviews.Product) *Product {
	return newProduct(vres.Projected)
}

// NewViewedProduct initializes viewed result type Product from result type
// Product using the given view.
func NewViewedProduct(res *Product, view string) *productviews.Product {
	p := newProductView(res)
	return &productviews.Product{Projected: p, View: "default"}
}

// NewSuccessResult initializes result type SuccessResult from viewed result
// type SuccessResult.
func NewSuccessResult(vres *productviews.SuccessResult) *SuccessResult {
	return newSuccessResult(vres.Projected)
}

// NewViewedSuccessResult initializes viewed result type SuccessResult from
// result type SuccessResult using the given view.
func NewViewedSuccessResult(res *SuccessResult, view string) *productviews.SuccessResult {
	p := newSuccessResultView(res)
	return &productviews.SuccessResult{Projected: p, View: "default"}
}

// newProduct converts projected type Product to service type Product.
func newProduct(vres *productviews.ProductView) *Product {
	res := &Product{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Unit != nil {
		res.Unit = *vres.Unit
	}
	if vres.CostPrice != nil {
		res.CostPrice = *vres.CostPrice
	}
	if vres.MarketPrice != nil {
		res.MarketPrice = *vres.MarketPrice
	}
	if vres.Note != nil {
		res.Note = *vres.Note
	}
	if vres.Image != nil {
		res.Image = *vres.Image
	}
	if vres.Code != nil {
		res.Code = *vres.Code
	}
	if vres.Size != nil {
		res.Size = *vres.Size
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.IsShelves != nil {
		res.IsShelves = *vres.IsShelves
	}
	if vres.Founder != nil {
		res.Founder = newFounder(vres.Founder)
	}
	return res
}

// newProductView projects result type Product to projected type ProductView
// using the "default" view.
func newProductView(res *Product) *productviews.ProductView {
	vres := &productviews.ProductView{
		ID:          &res.ID,
		Name:        &res.Name,
		Unit:        &res.Unit,
		CostPrice:   &res.CostPrice,
		MarketPrice: &res.MarketPrice,
		Note:        &res.Note,
		Image:       &res.Image,
		Code:        &res.Code,
		Size:        &res.Size,
		Type:        &res.Type,
		IsShelves:   &res.IsShelves,
	}
	if res.Founder != nil {
		vres.Founder = newFounderView(res.Founder)
	}
	return vres
}

// newFounder converts projected type Founder to service type Founder.
func newFounder(vres *productviews.FounderView) *Founder {
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
func newFounderView(res *Founder) *productviews.FounderView {
	vres := &productviews.FounderView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}

// newSuccessResult converts projected type SuccessResult to service type
// SuccessResult.
func newSuccessResult(vres *productviews.SuccessResultView) *SuccessResult {
	res := &SuccessResult{}
	if vres.OK != nil {
		res.OK = *vres.OK
	}
	return res
}

// newSuccessResultView projects result type SuccessResult to projected type
// SuccessResultView using the "default" view.
func newSuccessResultView(res *SuccessResult) *productviews.SuccessResultView {
	vres := &productviews.SuccessResultView{
		OK: &res.OK,
	}
	return vres
}