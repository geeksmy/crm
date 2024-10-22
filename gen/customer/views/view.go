// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Customer views
//
// Command:
// $ goa gen crm/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Customer is the viewed result type that is projected based on a view.
type Customer struct {
	// Type to project
	Projected *CustomerView
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

// CustomerView is a type that runs validations on a projected type.
type CustomerView struct {
	// ID
	ID *string
	// 客户姓名
	Name *string
	// 客户来源
	Src *int
	// 客户手机
	Mobile *string
	// 客户网址
	URL *string
	// 客户邮箱
	Email *string
	// 客户行业
	Industry *int
	// 客户等级
	Level *int
	// 备注
	Note *string
	// 客户地址
	Address *string
}

// SuccessResultView is a type that runs validations on a projected type.
type SuccessResultView struct {
	// success
	OK *bool
}

var (
	// CustomerMap is a map of attribute names in result type Customer indexed by
	// view name.
	CustomerMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"src",
			"mobile",
			"url",
			"email",
			"industry",
			"level",
			"note",
			"address",
		},
	}
	// SuccessResultMap is a map of attribute names in result type SuccessResult
	// indexed by view name.
	SuccessResultMap = map[string][]string{
		"default": []string{
			"ok",
		},
	}
)

// ValidateCustomer runs the validations defined on the viewed result type
// Customer.
func ValidateCustomer(result *Customer) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateCustomerView(result.Projected)
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

// ValidateCustomerView runs the validations defined on CustomerView using the
// "default" view.
func ValidateCustomerView(result *CustomerView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Src == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("src", "result"))
	}
	if result.Mobile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mobile", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	if result.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "result"))
	}
	if result.Industry == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("industry", "result"))
	}
	if result.Level == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("level", "result"))
	}
	if result.Note == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("note", "result"))
	}
	if result.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "result"))
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
