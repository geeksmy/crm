// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the Supplier service.
//
// Command:
// $ goa gen crm/design

package client

import (
	"fmt"
)

// GetSupplierPath returns the URL path to the Supplier service Get HTTP endpoint.
func GetSupplierPath(id string) string {
	return fmt.Sprintf("/api/supplier/%v", id)
}

// ListSupplierPath returns the URL path to the Supplier service List HTTP endpoint.
func ListSupplierPath() string {
	return "/api/supplier"
}

// UpdateSupplierPath returns the URL path to the Supplier service Update HTTP endpoint.
func UpdateSupplierPath() string {
	return "/api/supplier/update"
}

// CreateSupplierPath returns the URL path to the Supplier service Create HTTP endpoint.
func CreateSupplierPath() string {
	return "/api/supplier/create"
}

// DeleteSupplierPath returns the URL path to the Supplier service Delete HTTP endpoint.
func DeleteSupplierPath() string {
	return "/api/supplier/delete"
}