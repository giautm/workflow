package workflow

import "context"

type Context interface {
	Context() context.Context
	Trigger() string
	Param(key string) (interface{}, bool)
}

type SetParamContext interface {
	Context
	SetParam(key string, val interface{}) SetParamContext
}

func NewContext(
	context context.Context,
	trigger string,
) SetParamContext {
	return &wfContext{
		context: context,
		trigger: trigger,
		params:  make(map[string]interface{}),
	}
}

type wfContext struct {
	context context.Context
	trigger string
	params  map[string]interface{}
}

var _ = SetParamContext(&wfContext{})

func (c *wfContext) Context() context.Context {
	return c.context
}

func (c *wfContext) Trigger() string {
	return c.trigger
}

func (c *wfContext) Param(key string) (interface{}, bool) {
	v, ok := c.params[key]
	return v, ok
}

func (c *wfContext) SetParam(key string, val interface{}) SetParamContext {
	c.params[key] = val
	return c
}
