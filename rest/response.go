package rest

import (
	"github.com/nelsonlai-go/errs"
	"github.com/nelsonlai-go/ginger-engine/ginger"
	"github.com/nelsonlai-go/sql"
)

type Response struct {
	Success    bool            `json:"success"`
	Error      *Error          `json:"error,omitempty"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
}

type Error struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
}

func OK(ctx ginger.Context, data interface{}, p *sql.Pagination) {
	resp := &Response{
		Success:    true,
		Data:       data,
		Pagination: p,
	}
	ctx.SetParam("response", resp)
	ctx.JSON(200, resp)
}

func ERR(ctx ginger.Context, err errs.Error, data any) {
	resp := &Response{
		Success: false,
		Error: &Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
		Data: data,
	}
	ctx.JSON(200, resp)
}
