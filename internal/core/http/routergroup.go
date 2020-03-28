package http

// RouterGroup RouterGroup
type RouterGroup struct {
	Handlers []HandlerFunc
	basePath string
	engine *Engine
	root bool
	baseConfig *MethodConfig
}