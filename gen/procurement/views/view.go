// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Procurement views
//
// Command:
// $ goa gen crm/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Procurement is the viewed result type that is projected based on a view.
type Procurement struct {
	// Type to project
	Projected *ProcurementView
	// View to render
	View string
}

// SuccessResult is the viewed result type that is projected based on a view.
type SuccessResult struct {
	// Type to project
	Projected *SuccessResultView
	// View to render
	View string
}

// ProcurementView is a type that runs validations on a projected type.
type ProcurementView struct {
	// ID
	ID *string
	// 供应商
	Supplier *SupplierView
	// 采购编码
	Code *string
	// 采购金额
	Money *int
	// 采购还是退货
	IsSalesReturn *bool
	// 产品备注
	Note *string
	// 负责人
	Head *HeadView
	// 创建人
	Founder *FounderView
}

// SupplierView is a type that runs validations on a projected type.
type SupplierView struct {
	// ID
	ID *string
	// 供应商名
	Name *string
	// 级别
	Level *int
	// 联系人姓名
	ContactName *string
	// 联系人手机
	ContactPhone *string
	// 联系人地址
	ContactAddress *string
	// 产品备注
	Note *string
	// 负责人
	Head *HeadView
	// 创建人
	Founder *FounderView
}

// HeadView is a type that runs validations on a projected type.
type HeadView struct {
	// ID
	ID *string
	// 姓名
	Name *string
}

// FounderView is a type that runs validations on a projected type.
type FounderView struct {
	// ID
	ID *string
	// 姓名
	Name *string
}

// SuccessResultView is a type that runs validations on a projected type.
type SuccessResultView struct {
	// success
	OK *bool
}

var (
	// ProcurementMap is a map of attribute names in result type Procurement
	// indexed by view name.
	ProcurementMap = map[string][]string{
		"default": []string{
			"id",
			"supplier",
			"code",
			"money",
			"is_sales_return",
			"note",
			"head",
			"founder",
		},
	}
	// SuccessResultMap is a map of attribute names in result type SuccessResult
	// indexed by view name.
	SuccessResultMap = map[string][]string{
		"default": []string{
			"ok",
		},
	}
	// SupplierMap is a map of attribute names in result type Supplier indexed by
	// view name.
	SupplierMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"level",
			"contact_name",
			"contact_phone",
			"contact_address",
			"note",
			"head",
			"founder",
		},
	}
	// HeadMap is a map of attribute names in result type Head indexed by view name.
	HeadMap = map[string][]string{
		"default": []string{
			"id",
			"name",
		},
	}
	// FounderMap is a map of attribute names in result type Founder indexed by
	// view name.
	FounderMap = map[string][]string{
		"default": []string{
			"id",
			"name",
		},
	}
)

// ValidateProcurement runs the validations defined on the viewed result type
// Procurement.
func ValidateProcurement(result *Procurement) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProcurementView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSuccessResult runs the validations defined on the viewed result type
// SuccessResult.
func ValidateSuccessResult(result *SuccessResult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSuccessResultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateProcurementView runs the validations defined on ProcurementView
// using the "default" view.
func ValidateProcurementView(result *ProcurementView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "result"))
	}
	if result.Money == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("money", "result"))
	}
	if result.IsSalesReturn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("is_sales_return", "result"))
	}
	if result.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "result"))
	}
	if result.Supplier != nil {
		if err2 := ValidateSupplierView(result.Supplier); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Head != nil {
		if err2 := ValidateHeadView(result.Head); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Founder != nil {
		if err2 := ValidateFounderView(result.Founder); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSupplierView runs the validations defined on SupplierView using the
// "default" view.
func ValidateSupplierView(result *SupplierView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Level == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("level", "result"))
	}
	if result.ContactName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contact_name", "result"))
	}
	if result.ContactPhone == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contact_phone", "result"))
	}
	if result.ContactAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("contact_address", "result"))
	}
	if result.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "result"))
	}
	if result.Level != nil {
		if !(*result.Level == 1 || *result.Level == 2 || *result.Level == 3 || *result.Level == 4) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.level", *result.Level, []interface{}{1, 2, 3, 4}))
		}
	}
	if result.Head != nil {
		if err2 := ValidateHeadView(result.Head); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Founder != nil {
		if err2 := ValidateFounderView(result.Founder); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateHeadView runs the validations defined on HeadView using the
// "default" view.
func ValidateHeadView(result *HeadView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateFounderView runs the validations defined on FounderView using the
// "default" view.
func ValidateFounderView(result *FounderView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateSuccessResultView runs the validations defined on SuccessResultView
// using the "default" view.
func ValidateSuccessResultView(result *SuccessResultView) (err error) {
	if result.OK == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ok", "result"))
	}
	return
}
