package ginger

import (
	"github.com/gin-gonic/gin"
)

type Ginger interface {
	Run(addr string)
	GET(path string, handler HandlerFunc, middleware ...gin.HandlerFunc)
	POST(path string, handler HandlerFunc, middleware ...gin.HandlerFunc)
	PUT(path string, handler HandlerFunc, middleware ...gin.HandlerFunc)
	DELETE(path string, handler HandlerFunc, middleware ...gin.HandlerFunc)
	Register(handler RegisterHandlerFunc, option RegisterOption)
	InitFunc(f ...func())
	Middleware(middleware ...gin.HandlerFunc)
}

type gingerEngine struct {
	gin.Engine

	initFuncs  []func()
	middleware []gin.HandlerFunc
	routes     []route
}

type route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

func New() Ginger {
	return &gingerEngine{
		Engine:     *gin.New(),
		initFuncs:  make([]func(), 0),
		middleware: make([]gin.HandlerFunc, 0),
		routes:     make([]route, 0),
	}
}

func (e *gingerEngine) Run(addr string) {
	for _, f := range e.initFuncs {
		f()
	}
	for _, m := range e.middleware {
		e.Use(m)
	}
	for _, r := range e.routes {
		e.Handle(r.method, r.path, r.handlers...)
	}
	e.Engine.Run(addr)
}

func (e *gingerEngine) Register(handler RegisterHandlerFunc, option RegisterOption) {
	handler(e, option)
}

func (e *gingerEngine) InitFunc(f ...func()) {
	e.initFuncs = append(e.initFuncs, f...)
}

func (e *gingerEngine) Middleware(middleware ...gin.HandlerFunc) {
	e.middleware = append(e.middleware, middleware...)
}

func (e *gingerEngine) GET(path string, handler HandlerFunc, middleware ...gin.HandlerFunc) {
	e.setupRoute("GET", path, handler, middleware...)
}

func (e *gingerEngine) POST(path string, handler HandlerFunc, middleware ...gin.HandlerFunc) {
	e.setupRoute("POST", path, handler, middleware...)
}

func (e *gingerEngine) PUT(path string, handler HandlerFunc, middleware ...gin.HandlerFunc) {
	e.setupRoute("PUT", path, handler, middleware...)
}

func (e *gingerEngine) DELETE(path string, handler HandlerFunc, middleware ...gin.HandlerFunc) {
	e.setupRoute("DELETE", path, handler, middleware...)
}

func (e *gingerEngine) setupRoute(method string, path string, handler HandlerFunc, middleware ...gin.HandlerFunc) {
	hs := append(middleware, handler.GinHandler())
	e.routes = append(e.routes, route{method, path, hs})
}
