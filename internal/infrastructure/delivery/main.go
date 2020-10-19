// Package delivery implement web server
package delivery

import (
	"github.com/gin-gonic/gin"
)

// NewEngine create a Engine
func NewEngine() *Engine {
	e := new(Engine)
	e.root = gin.Default()
	e.api = e.root.Group("/api")
	return e
}

// Engine shortcuts to gin.Engine
type Engine struct {
	root *gin.Engine
	api *gin.RouterGroup
}

// Run start engine
func (e *Engine) Run(addr string) error {
	return e.root.Run(addr)
}