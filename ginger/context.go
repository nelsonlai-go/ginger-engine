package ginger

import "github.com/gin-gonic/gin"

type Context interface {
	Param(string) interface{}
	SetParam(string, interface{})
	ClientIP() string
	RequestPath() string
	ShouldBindUri(interface{}) error
	ShouldBindQuery(interface{}) error
	ShouldBindJSON(interface{}) error
}

type gingerContext struct {
	gin.Context
	param map[string]interface{}
}

func (g *gingerContext) Param(key string) interface{} {
	return g.param[key]
}

func (g *gingerContext) SetParam(key string, val interface{}) {
	g.param[key] = val
}

func (g *gingerContext) RequestPath() string {
	p := g.FullPath()
	if len(p) > 0 {
		return p
	}
	return g.Request.URL.Path
}
