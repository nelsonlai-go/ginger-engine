package ginger

import "github.com/gin-gonic/gin"

type Context interface {
	Param(string) any
	SetParam(string, any)
	ClientIP() string
	RequestPath() string
	ShouldBindUri(any) error
	ShouldBindQuery(any) error
	ShouldBindJSON(any) error
	JSON(code int, obj any)
}

type gingerContext struct {
	gin.Context
	param map[string]any
}

func (g *gingerContext) Param(key string) any {
	return g.param[key]
}

func (g *gingerContext) SetParam(key string, val any) {
	g.param[key] = val
}

func (g *gingerContext) RequestPath() string {
	p := g.FullPath()
	if len(p) > 0 {
		return p
	}
	return g.Request.URL.Path
}
