package rest

import (
	"encoding/json"

	"github.com/nelsonlai-go/ginger-engine/ginger"
	"github.com/nelsonlai-go/sql"
)

type statementRequest struct {
	Statement string `form:"statement"`
}

func StatementRequest(ctx ginger.Context) *sql.Statement {
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
