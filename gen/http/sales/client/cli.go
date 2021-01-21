// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Sales HTTP client CLI support package
//
// Command:
// $ goa gen crm/design

package client

import (
	sales "crm/gen/sales"
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetPayload builds the payload for the Sales Get endpoint from CLI flags.
func BuildGetPayload(salesGetID string, salesGetToken string) (*sales.GetPayload, error) {
	var id string
	{
		id = salesGetID
	}
	var token string
	{
		token = salesGetToken
	}
	v := &sales.GetPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildListPayload builds the payload for the Sales List endpoint from CLI
// flags.
func BuildListPayload(salesListToken string) (*sales.ListPayload, error) {
	var token string
	{
		token = salesListToken
	}
	v := &sales.ListPayload{}
	v.Token = token

	return v, nil
}

// BuildUpdatePayload builds the payload for the Sales Update endpoint from CLI
// flags.
func BuildUpdatePayload(salesUpdateBody string, salesUpdateToken string) (*sales.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(salesUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"consignment_date\": \"2021-01-01\",\n      \"customer_id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"head_id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"is_sales_return\": false,\n      \"money\": 123,\n      \"name\": \"xx销售\",\n      \"note\": \"备注\"\n   }'")
		}
		if !(body.IsSalesReturn == false || body.IsSalesReturn == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.is_sales_return", body.IsSalesReturn, []interface{}{false, true}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = salesUpdateToken
	}
	v := &sales.UpdatePayload{
		ID:              body.ID,
		Name:            body.Name,
		CustomerID:      body.CustomerID,
		Money:           body.Money,
		ConsignmentDate: body.ConsignmentDate,
		IsSalesReturn:   body.IsSalesReturn,
		Note:            body.Note,
		HeadID:          body.HeadID,
	}
	v.Token = token

	return v, nil
}

// BuildCreatePayload builds the payload for the Sales Create endpoint from CLI
// flags.
func BuildCreatePayload(salesCreateBody string, salesCreateToken string) (*sales.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(salesCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"code\": \"123asd123123asd\",\n      \"consignment_date\": \"2021-01-01\",\n      \"customer_id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"founder_id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"head_id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"is_sales_return\": false,\n      \"money\": 123,\n      \"name\": \"xx销售\",\n      \"note\": \"备注\"\n   }'")
		}
		if !(body.IsSalesReturn == false || body.IsSalesReturn == true) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.is_sales_return", body.IsSalesReturn, []interface{}{false, true}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = salesCreateToken
	}
	v := &sales.CreatePayload{
		Name:            body.Name,
		Code:            body.Code,
		CustomerID:      body.CustomerID,
		Money:           body.Money,
		ConsignmentDate: body.ConsignmentDate,
		IsSalesReturn:   body.IsSalesReturn,
		Note:            body.Note,
		HeadID:          body.HeadID,
		FounderID:       body.FounderID,
	}
	v.Token = token

	return v, nil
}

// BuildDeletePayload builds the payload for the Sales Delete endpoint from CLI
// flags.
func BuildDeletePayload(salesDeleteBody string, salesDeleteToken string) (*sales.DeletePayload, error) {
	var err error
	var body DeleteRequestBody
	{
		err = json.Unmarshal([]byte(salesDeleteBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ids\": [\n         \"91cc3eb9-ddc0-4cf7-a62b-c85df1a9166f\",\n         \"91cc3eb9-ddc0-4cf7-a62b-c85df1a9166f\",\n         \"91cc3eb9-ddc0-4cf7-a62b-c85df1a9166f\"\n      ]\n   }'")
		}
		if body.Ids == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("ids", "body"))
		}
		if len(body.Ids) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.ids", body.Ids, len(body.Ids), 100, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = salesDeleteToken
	}
	v := &sales.DeletePayload{}
	if body.Ids != nil {
		v.Ids = make([]string, len(body.Ids))
		for i, val := range body.Ids {
			v.Ids[i] = val
		}
	}
	v.Token = token

	return v, nil
}
