package ginger

import "github.com/gin-gonic/gin"

type Context[T any] struct {
	GinCtx   *gin.Context
	Request  *T
	Response interface{}

	// additional parameters
	// you may store auth credentials, session
	param map[string]interface{}
}

func (ctx *Context[T]) Param(key string) interface{} {
	return ctx.param[key]
}

func (ctx *Context[T]) SetParam(key string, value interface{}) {
	ctx.param[key] = value
}
