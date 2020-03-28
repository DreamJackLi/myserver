package http

import (
	"context"
	"net/http"
)

type Context struct {
	context.Context

	Request *http.Request
	Writer  http.ResponseWriter

	index    int8
	handlers []HandlerFunc

	Keys map[string]interface{}

	Error error

	method string
	engine *Engine

	RoutePath string

	Params Params
}