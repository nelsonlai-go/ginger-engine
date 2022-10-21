package ginger

import (
	"log"
)

type RegisterHandlerFunc func(e Ginger, opt RegisterOption)

type RegisterOption interface {
	Param(key string, required bool) interface{}
	SetParam(m map[string]interface{})
}

func NewRegisterOption(m map[string]string) RegisterOption {
	return &registerOption{
		param: make(map[string]interface{}),
	}
}

type registerOption struct {
	param map[string]interface{}
}

func (o *registerOption) Param(key string, required bool) interface{} {
	if v, ok := o.param[key]; ok {
		return v
	}
	if required {
		log.Fatalf("missing required param %s", key)
	}
	return nil
}

func (o *registerOption) SetParam(m map[string]interface{}) {
	for k, v := range m {
		o.param[k] = v
	}
}
