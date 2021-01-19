// Code generated by goa v3.2.6, DO NOT EDIT.
//
// User HTTP server encoders and decoders
//
// Command:
// $ goa gen crm/design

package server

import (
	"context"
	user "crm/gen/user"
	"io"
	"net/http"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeLoginByUsernameResponse returns an encoder for responses returned by
// the User LoginByUsername endpoint.
func EncodeLoginByUsernameResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*user.LoginByUsernameResult)
		enc := encoder(ctx, w)
		body := NewLoginByUsernameResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginByUsernameRequest returns a decoder for requests sent to the User
// LoginByUsername endpoint.
func DecodeLoginByUsernameRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body LoginByUsernameRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateLoginByUsernameRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewLoginByUsernamePayload(&body)

		return payload, nil
	}
}

// EncodeLoginByUsernameError returns an encoder for errors returned by the
// LoginByUsername User endpoint.
func EncodeLoginByUsernameError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_server_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewLoginByUsernameInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal_server_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewLoginByUsernameBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdatePasswordResponse returns an encoder for responses returned by
// the User UpdatePassword endpoint.
func EncodeUpdatePasswordResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*user.UpdatePasswordResult)
		enc := encoder(ctx, w)
		body := NewUpdatePasswordResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdatePasswordRequest returns a decoder for requests sent to the User
// UpdatePassword endpoint.
func DecodeUpdatePasswordRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdatePasswordRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdatePasswordRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token string
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdatePasswordPayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeUpdatePasswordError returns an encoder for errors returned by the
// UpdatePassword User endpoint.
func EncodeUpdatePasswordError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_server_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdatePasswordInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal_server_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdatePasswordBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetCaptchaImageResponse returns an encoder for responses returned by
// the User GetCaptchaImage endpoint.
func EncodeGetCaptchaImageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*user.GetCaptchaImageResult)
		enc := encoder(ctx, w)
		body := NewGetCaptchaImageResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetCaptchaImageError returns an encoder for errors returned by the
// GetCaptchaImage User endpoint.
func EncodeGetCaptchaImageError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_server_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetCaptchaImageInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal_server_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetCaptchaImageBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalUserSessionToSessionResponseBody builds a value of type
// *SessionResponseBody from a value of type *user.Session.
func marshalUserSessionToSessionResponseBody(v *user.Session) *SessionResponseBody {
	if v == nil {
		return nil
	}
	res := &SessionResponseBody{}
	if v.User != nil {
		res.User = marshalUserUserToUserResponseBody(v.User)
	}
	if v.Credentials != nil {
		res.Credentials = marshalUserCredentialsToCredentialsResponseBody(v.Credentials)
	}

	return res
}

// marshalUserUserToUserResponseBody builds a value of type *UserResponseBody
// from a value of type *user.User.
func marshalUserUserToUserResponseBody(v *user.User) *UserResponseBody {
	res := &UserResponseBody{
		ID:       v.ID,
		Username: v.Username,
		Mobile:   v.Mobile,
	}

	return res
}

// marshalUserCredentialsToCredentialsResponseBody builds a value of type
// *CredentialsResponseBody from a value of type *user.Credentials.
func marshalUserCredentialsToCredentialsResponseBody(v *user.Credentials) *CredentialsResponseBody {
	res := &CredentialsResponseBody{
		Token:     v.Token,
		ExpiresIn: v.ExpiresIn,
	}

	return res
}

// marshalUserSuccessResultToSuccessResultResponseBody builds a value of type
// *SuccessResultResponseBody from a value of type *user.SuccessResult.
func marshalUserSuccessResultToSuccessResultResponseBody(v *user.SuccessResult) *SuccessResultResponseBody {
	if v == nil {
		return nil
	}
	res := &SuccessResultResponseBody{
		OK: v.OK,
	}

	return res
}

// marshalUserCaptchaToCaptchaResponseBody builds a value of type
// *CaptchaResponseBody from a value of type *user.Captcha.
func marshalUserCaptchaToCaptchaResponseBody(v *user.Captcha) *CaptchaResponseBody {
	if v == nil {
		return nil
	}
	res := &CaptchaResponseBody{
		Image:     v.Image,
		CaptchaID: v.CaptchaID,
	}

	return res
}
