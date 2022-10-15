package ginger

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

type RegisterHandlerFunc[T any] func(e *Engine, option *T)

func Register[T any](e *Engine, pluginID string, f RegisterHandlerFunc[T], option *T) {
	if _, ok := e.pluginSetup.RegisterKeys[pluginID]; ok {
		return // already registered, skip
	}
	e.pluginSetup.RegisterKeys[pluginID] = true
	f(e, option)
}

type HandlerFunc[T any] func(*Context[T])

func GET[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "GET", path, handler, middleware...)
}

func POST[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "POST", path, handler, middleware...)
}

func PUT[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "PUT", path, handler, middleware...)
}

func DELETE[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "DELETE", path, handler, middleware...)
}

func _setupRoute[T any](e *Engine, method string, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	ginHandlers := append(middleware, _toGinHandler(e, handler))
	e.pluginSetup.Routes = append(e.pluginSetup.Routes, route{
		method:   method,
		path:     path,
		handlers: ginHandlers,
	})
}

func _toGinHandler[T any](e *Engine, handler HandlerFunc[T]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(T)
		tagKeys := _parseRequestTagKeys(request)

		for _, key := range tagKeys {
			switch key {
			case "uri":
				if err := ctx.ShouldBindUri(request); err != nil {
					log.Println(err)
				}
			case "json":
				if err := ctx.ShouldBindJSON(request); err != nil {
					log.Println(err)
				}
			case "form":
				if err := ctx.ShouldBindQuery(request); err != nil {
					log.Println(err)
				}
			}
		}

		context := &Context[T]{
			GinCtx:   ctx,
			Request:  request,
			Response: nil,
			param:    make(map[string]interface{}),
		}

		handler(context)
	}
}

func _parseRequestTagKeys[T any](request *T) []string {
	var keyMap = make(map[string]bool)

	numOfField := reflect.TypeOf(request).Elem().NumField()
	for i := 0; i < numOfField; i++ {
		tag := reflect.TypeOf(request).Elem().Field(i).Tag

		check := tag.Get("uri")
		if check != "" {
			keyMap["uri"] = true
		}

		check = tag.Get("json")
		if check != "" {
			keyMap["json"] = true
		}

		check = tag.Get("form")
		if check != "" {
			keyMap["form"] = true
		}
	}

	var keys []string
	for key := range keyMap {
		keys = append(keys, key)
	}

	return keys
}
