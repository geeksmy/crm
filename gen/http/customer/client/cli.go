// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Customer HTTP client CLI support package
//
// Command:
// $ goa gen crm/design

package client

import (
	customer "crm/gen/customer"
	"encoding/json"
	"fmt"
	"strconv"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetPayload builds the payload for the Customer Get endpoint from CLI
// flags.
func BuildGetPayload(customerGetID string, customerGetToken string) (*customer.GetPayload, error) {
	var id string
	{
		id = customerGetID
	}
	var token string
	{
		token = customerGetToken
	}
	v := &customer.GetPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildListPayload builds the payload for the Customer List endpoint from CLI
// flags.
func BuildListPayload(customerListCursor string, customerListLimit string, customerListToken string) (*customer.ListPayload, error) {
	var err error
	var cursor *int
	{
		if customerListCursor != "" {
			var v int64
			v, err = strconv.ParseInt(customerListCursor, 10, 64)
			val := int(v)
			cursor = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for cursor, must be INT")
			}
		}
	}
	var limit *int
	{
		if customerListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(customerListLimit, 10, 64)
			val := int(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
		}
	}
	var token string
	{
		token = customerListToken
	}
	v := &customer.ListPayload{}
	v.Cursor = cursor
	v.Limit = limit
	v.Token = token

	return v, nil
}

// BuildUpdatePayload builds the payload for the Customer Update endpoint from
// CLI flags.
func BuildUpdatePayload(customerUpdateBody string, customerUpdateToken string) (*customer.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(customerUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": \"咸阳\",\n      \"email\": \"adb@adb.com\",\n      \"id\": \"519151ca-6250-4eec-8016-1e14a68dc448\",\n      \"industry\": 1,\n      \"level\": 1,\n      \"mobile\": \"18000001234\",\n      \"name\": \"张三\",\n      \"note\": \"备注\",\n      \"src\": 1,\n      \"url\": \"http://www.baidu.com\"\n   }'")
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
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = customerUpdateToken
	}
	v := &customer.UpdatePayload{
		ID:       body.ID,
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

	return v, nil
}

// BuildCreatePayload builds the payload for the Customer Create endpoint from
// CLI flags.
func BuildCreatePayload(customerCreateBody string, customerCreateToken string) (*customer.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(customerCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": \"咸阳\",\n      \"email\": \"adb@adb.com\",\n      \"industry\": 1,\n      \"level\": 1,\n      \"mobile\": \"18000001234\",\n      \"name\": \"张三\",\n      \"note\": \"备注\",\n      \"src\": 1,\n      \"url\": \"http://www.baidu.com\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.mobile", body.Mobile, goa.FormatRegexp))

		err = goa.MergeErrors(err, goa.ValidatePattern("body.mobile", body.Mobile, "^1(?:3\\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\\d|9\\d)\\d{8}$"))
		if utf8.RuneCountInString(body.Mobile) < 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", body.Mobile, utf8.RuneCountInString(body.Mobile), 11, true))
		}
		if utf8.RuneCountInString(body.Mobile) > 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", body.Mobile, utf8.RuneCountInString(body.Mobile), 11, false))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.url", body.URL, goa.FormatURI))

		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = customerCreateToken
	}
	v := &customer.CreatePayload{
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

	return v, nil
}

// BuildDeletePayload builds the payload for the Customer Delete endpoint from
// CLI flags.
func BuildDeletePayload(customerDeleteBody string, customerDeleteToken string) (*customer.DeletePayload, error) {
	var err error
	var body DeleteRequestBody
	{
		err = json.Unmarshal([]byte(customerDeleteBody), &body)
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
		token = customerDeleteToken
	}
	v := &customer.DeletePayload{}
	if body.Ids != nil {
		v.Ids = make([]string, len(body.Ids))
		for i, val := range body.Ids {
			v.Ids[i] = val
		}
	}
	v.Token = token

	return v, nil
}
