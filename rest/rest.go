package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonlai-go/errs"
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

func OK(ctx *gin.Context, data interface{}, p *sql.Pagination) {
	resp := &Response{
		Success:    true,
		Data:       data,
		Pagination: p,
	}
	ctx.Set("response", resp) // write to ctx for perform tests
	ctx.JSON(200, resp)
}

func ERR(ctx *gin.Context, err errs.Error, data any) {
	resp := &Response{
		Success: false,
		Error: &Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
		Data: data,
	}
	ctx.Set("response", resp) // write to ctx for perform tests
	ctx.JSON(200, resp)
}
