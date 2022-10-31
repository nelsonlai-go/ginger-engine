package rest

import (
	"github.com/nelsonlai-go/ginger-engine/ginger"
	"github.com/nelsonlai-go/sql"
)

func SortRequest(ctx ginger.Context) *sql.Sort {
	return ginger.GetRequest[sql.Sort](ctx)
}
