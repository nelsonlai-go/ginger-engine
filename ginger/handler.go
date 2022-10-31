package ginger

import "github.com/gin-gonic/gin"

type HandlerFunc func(ctx Context)

func (h HandlerFunc) GinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h(&gingerContext{*ctx, make(map[string]any)})
	}
}
