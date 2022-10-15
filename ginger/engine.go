package ginger

import "github.com/gin-gonic/gin"

type Engine struct {
	GinEngine   *gin.Engine
	pluginSetup *pluginSetup
}

type pluginSetup struct {
	InitializeFunctions []func()
	GinMiddleware       []gin.HandlerFunc
	Routes              []route
	PreRunFunctions     []func()
	RegisterKeys        map[string]bool
}

type route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

func New() *Engine {
	return &Engine{
		GinEngine: gin.Default(),
		pluginSetup: &pluginSetup{
			InitializeFunctions: make([]func(), 0),
			GinMiddleware:       make([]gin.HandlerFunc, 0),
			Routes:              make([]route, 0),
			PreRunFunctions:     make([]func(), 0),
			RegisterKeys:        make(map[string]bool),
		},
	}
}

// panic if the required plugins have not been registered
func (e *Engine) RequireDependencies(pluginIDs ...string) {
	for _, pluginID := range pluginIDs {
		if _, ok := e.pluginSetup.RegisterKeys[pluginID]; !ok {
			panic("The plugin (" + pluginID + ") has not been registered")
		}
	}
}

func (e *Engine) RegisterInitializeFunctions(f ...func()) {
	e.pluginSetup.InitializeFunctions = append(e.pluginSetup.InitializeFunctions, f...)
}

func (e *Engine) RegisterMiddleware(middleware ...gin.HandlerFunc) {
	e.pluginSetup.GinMiddleware = append(e.pluginSetup.GinMiddleware, middleware...)
}

func (e *Engine) PreRunFunctions(f ...func()) {
	e.pluginSetup.PreRunFunctions = append(e.pluginSetup.PreRunFunctions, f...)
}

func (e *Engine) Run(addr string) {
	for _, f := range e.pluginSetup.InitializeFunctions {
		f()
	}
	for _, m := range e.pluginSetup.GinMiddleware {
		e.GinEngine.Use(m)
	}
	for _, r := range e.pluginSetup.Routes {
		e.GinEngine.Handle(r.method, r.path, r.handlers...)
	}
	for _, f := range e.pluginSetup.PreRunFunctions {
		f()
	}
	e.GinEngine.Run(addr)
}
