package horizon

import (
	u "net/url"
	"strconv"
	"strings"
)

type requestBuilder struct {
	BaseURL  string
	segments []string
	order    OrderDirection
	cursor   Cursor
	limit    Limit
}

func (rb *requestBuilder) AddSegment(segment interface{}) *requestBuilder {
	var stringSegment string

	switch segment := segment.(type) {
	case int32:
		stringSegment = strconv.FormatInt(int64(segment), 32)
	case int64:
		stringSegment = strconv.FormatInt(segment, 64)
	case string:
		stringSegment = segment
	default:
		return rb
	}

	rb.segments = append(rb.segments, stringSegment)
	return rb
}

func (rb *requestBuilder) Params(params ...interface{}) *requestBuilder {
	for _, param := range params {
		switch param := param.(type) {
		case OrderDirection:
			rb.Order(param)
		case Cursor:
			rb.Cursor(param)
		case Limit:
			rb.Limit(param)
		case PageParams:
			rb.Order(param.Order)
			rb.Cursor(param.Cursor)
			rb.Limit(param.Limit)
		}
	}
	return rb
}

func (rb *requestBuilder) Order(order OrderDirection) *requestBuilder {
	rb.order = order
	return rb
}

func (rb *requestBuilder) Cursor(cursor Cursor) *requestBuilder {
	rb.cursor = cursor
	return rb
}

func (rb *requestBuilder) Limit(limit Limit) *requestBuilder {
	rb.limit = limit
	return rb
}

func (rb *requestBuilder) Build() string {
	url := u.URL{
		Path: rb.BaseURL + "/" + strings.Join(rb.segments, "/"),
	}
	values := u.Values{}

	if rb.order != "" {
		values.Set("order", string(rb.order))
	}

	if rb.cursor != "" {
		values.Set("cursor", string(rb.cursor))
	}

	if rb.limit != 0 {
		values.Set("limit", strconv.Itoa(int(rb.limit)))
	}

	url.RawQuery = values.Encode()
	return url.String()
}
