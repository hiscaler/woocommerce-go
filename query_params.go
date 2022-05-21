package woocommerce

import (
	"github.com/google/go-querystring/query"
	"net/url"
	"strings"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

const (
	ViewContext = "view"
	EditContext = "edit"
)

type queryParams struct {
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Order   string `url:"order,omitempty"`
	OrderBy string `url:"order_by,omitempty"`
	Context string `url:"context,omitempty"`
}

func (q *queryParams) TidyVars() *queryParams {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PerPage <= 0 {
		q.PerPage = 10
	}
	if q.Offset < 0 {
		q.Offset = 0
	}

	if q.Order == "" {
		q.Order = SortAsc
	} else {
		q.Order = strings.ToLower(q.Order)
		if q.Order != SortDesc {
			q.OrderBy = SortAsc
		}
	}

	if q.Context == "" {
		q.Context = ViewContext
	} else {
		q.Context = strings.ToLower(q.Context)
		if q.Context != EditContext {
			q.Context = ViewContext
		}
	}
	return q
}

func toValues(i interface{}) (values url.Values) {
	values, _ = query.Values(i)
	return
}
