package woocommerce

import "strings"

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

const (
	ViewContext = "view"
	EditContext = "edit"
)

type Query struct {
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Order   string `url:"order,omitempty"`
	OrderBy string `url:"order_by,omitempty"`
	Context string `url:"context,omitempty"`
}

func (q *Query) TidyVars() *Query {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PerPage <= 0 {
		q.PerPage = 10
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
	if q.Order == "" || !strings.EqualFold(q.Order, SortDesc) {
		q.Order = SortAsc
	}
	if q.Context == "" || !strings.EqualFold(q.Context, EditContext) {
		q.Context = ViewContext
	}
	return q
}
