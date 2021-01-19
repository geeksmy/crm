package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

func newTextEncoder(w io.Writer, ct string) goahttp.Encoder {
	return &textEncoder{w, ct}
}

type textEncoder struct {
	w  io.Writer
	ct string
}

func (e *textEncoder) Encode(v interface{}) error {
	var err error

	switch c := v.(type) {
	case string:
		_, err = e.w.Write([]byte(c))
	case *string: // v may be a string pointer when the Response Body is set to the field of a custom response type.
		_, err = e.w.Write([]byte(*c))
	case []byte:
		_, err = e.w.Write(c)
	default:
		err = fmt.Errorf("can't encode %T as %s", c, e.ct)
	}

	return err
}

func ResponseEncoder(ctx context.Context, w http.ResponseWriter) goahttp.Encoder {
	var ct string
	{
		if a := ctx.Value(goahttp.ContentTypeKey); a != nil {
			ct = a.(string)
		}
	}

	if ct == "application/octet-stream" || ct == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		return newTextEncoder(w, ct)
	}

	return goahttp.ResponseEncoder(ctx, w)
}
