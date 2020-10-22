// Package delivery implement web server
package delivery

import (
	"github.com/gin-gonic/gin"
	"path"
)

// NewEngine create a Engine
func NewEngine() *Engine {
	e := new(Engine)
	e.engine = gin.Default()
	e.root = e.engine.Group("/")
	e.views = e.root
	e.api = e.root.Group("/api")
	return e
}

// Engine shortcuts to gin.Engine
type Engine struct {
	engine *gin.Engine
	root *gin.RouterGroup
	views *gin.RouterGroup
	api *gin.RouterGroup
}

// SetResourceRoot load html templates & add route for static assets
func (e *Engine) SetResourceRoot(root string) {
	e.engine.Static("/public", path.Join(root, "public"))
	e.engine.LoadHTMLGlob(path.Join(root, "templates/**/*"))
}

// Run start engine
func (e *Engine) Run(addr string) error {
	return e.engine.Run(addr)
}