package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   int
	router *gin.Engine
}

// creates a new api server with the specified port
func NewServer(port int) *Server {
	return &Server{
		Port:   port,
		router: nil,
	}
}

// tells server to listen on port
// returns if allocation failed or server exits
func (s *Server) Start() error {
	// create gin engine
	if s.router = gin.Default(); s.router == nil {
		return fmt.Errorf("Allocation failed!")
	}

	// disable trusting all proxies
	s.router.SetTrustedProxies(nil)

	// configure http server
	// TODO: TLS config
	server := http.Server{
		Addr:           s.getAddress(),
		Handler:        s.router,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	// crud user, routes:
	user := s.router.Group("/user")
	{
		user.GET("/get/:id", s.GetUser)
		user.POST("/create", s.CreateUser)
		user.PUT("/update", s.UpdateUser)
		user.DELETE("/delete/:id", s.DeleteUser)
	}

	fmt.Printf("Starting server on port %s", server.Addr)
	return server.ListenAndServe()
}

// get the port as a string
func (s *Server) getAddress() string {
	return fmt.Sprintf(":%d", s.Port)
}
