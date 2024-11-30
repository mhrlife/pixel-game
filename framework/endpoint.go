package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type EndpointHandler func(ctx *Context) error
type Endpoints struct {
	endpoints   map[string]EndpointHandler
	middlewares []echo.MiddlewareFunc
}

func (e *Endpoints) Register(action string, handler EndpointHandler) {
	e.endpoints[action] = handler
}

func (e *Endpoints) Middleware(middlewareFunc echo.MiddlewareFunc) {
	e.middlewares = append(e.middlewares, middlewareFunc)
}

func (e *Endpoints) Actions() []string {
	return lo.Keys(e.endpoints)
}

type Endpoint interface {
	Endpoints(router *Endpoints)
}
