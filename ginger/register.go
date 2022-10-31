package ginger

import (
	"log"
)

type RegisterHandlerFunc func(e Ginger, opt RegisterOption)

type RegisterOption interface {
	Param(key string, required bool) any
	SetParam(m map[string]any)
}

func NewRegisterOption(m map[string]any) RegisterOption {
	return &registerOption{
		param: m,
	}
}

type registerOption struct {
	param map[string]any
}

func (o *registerOption) Param(key string, required bool) any {
	if v, ok := o.param[key]; ok {
		return v
	}
	if required {
		log.Fatalf("missing required param %s", key)
	}
	return nil
}

func (o *registerOption) SetParam(m map[string]any) {
	for k, v := range m {
		o.param[k] = v
	}
}
