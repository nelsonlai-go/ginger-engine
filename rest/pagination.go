package rest

import (
	"github.com/nelsonlai-go/ginger-engine/ginger"
	"github.com/nelsonlai-go/sql"
)

func PaginationRequest(ctx ginger.Context) *sql.Pagination {
	return ginger.GetRequest[sql.Pagination](ctx)
}
