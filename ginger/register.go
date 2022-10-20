package ginger

import (
	"log"
)

type RegisterHandlerFunc func(e Ginger, opt RegisterOption)

type RegisterOption interface {
	Param(key string, required bool) string
	SetParam(m map[string]string)
}

func NewRegisterOption(m map[string]string) RegisterOption {
	return &registerOption{}
}

type registerOption struct {
	param map[string]string
}

func (o *registerOption) Param(key string, required bool) string {
	if v, ok := o.param[key]; ok {
		return v
	}
	if required {
		log.Fatalf("missing required param %s", key)
	}
	return ""
}

func (o *registerOption) SetParam(m map[string]string) {
	if o.param == nil {
		o.param = make(map[string]string)
	}
	for k, v := range m {
		o.param[k] = v
	}
}
