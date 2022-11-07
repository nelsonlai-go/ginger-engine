package rest

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/nelsonlai-go/ginger-engine/ginger"
	"github.com/nelsonlai-go/sql"
)

type statementRequest struct {
	Statement string `form:"statement"`
}

func GetStatementRequest(ctx *gin.Context) *sql.Statement {
	var f statementRequest
	if err := ctx.ShouldBindQuery(&f); err != nil {
		return nil
	}

	if f.Statement == "" {
		return nil
	}

	var clause sql.Statement
	if err := json.Unmarshal([]byte(f.Statement), &clause); err != nil {
		return nil
	}

	return &clause
}

func GetPaginationRequest(ctx *gin.Context) *sql.Pagination {
	return ginger.GetRequest[sql.Pagination](ctx)
}

func GetSortRequest(ctx *gin.Context) *sql.Sort {
	return ginger.GetRequest[sql.Sort](ctx)
}
