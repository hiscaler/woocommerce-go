package woocommerce

type Query struct {
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
	Offset  int    `json:"offset,omitempty"`
	Context string `url:"context,omitempty"`
	Order   string `url:"order,omitempty"`
	OrderBy string `url:"order_by,omitempty"`
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
	if q.Order != "asc" && q.Order != "desc" {
		q.Order = ""
	}
	return q
}
