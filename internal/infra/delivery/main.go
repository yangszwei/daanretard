package delivery

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"strings"
)

// NewServer create an instance of Server
func NewServer(data string) *Server {
	s := new(Server)
	s.Engine = gin.Default()
	s.root = s.Engine.Group("/")
	return s
}

// Server wrapper of gin.Engine
type Server struct {
	Engine *gin.Engine      // exposed for testing
	root   *gin.RouterGroup // for middlewares
	web    *gin.RouterGroup // web user interface
	api    *gin.RouterGroup // api
	data   string           // folder path for storing data files
}

// LoadTemplates set html template for server
func (s *Server) LoadTemplates() {
	t := template.New("")
	for name, file := range ui.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".gohtml") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			panic(err)
		}
	}
	s.Engine.SetHTMLTemplate(t)
}

// SetupRouterGroups setup router groups, should be called AFTER middlewares
// are set on server.root so that the middlewares are applied to all route
// groups
func (s *Server) SetupRouterGroups() {
	s.web = s.root
	s.api = s.root.Group("/api")
}

// Run start server
func (s *Server) Run(addr string) error {
	return s.Engine.Run(addr)
}
