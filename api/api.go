package api

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// wrapper for gin engine holding router port
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
		return fmt.Errorf("Allocation failed: %s", err.Error())
	}

	// disable trusting all proxies
	err := s.router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	// configure http server
	server := http.Server{
		Addr:           s.getAddress(),
		Handler:        s.router,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	// user autentication and creation routes
	jwtAuth, err := s.addJwtMiddleware("")
	if err != nil {
		return fmt.Errorf("Error creating jwt middleware: %s", err.Error())
	}
	// authentication routes
	auth := s.router.Group("/auth")
	{
		auth.POST("/user/create", s.createUser)
		auth.POST("/user/login", jwtAuth.LoginHandler)
		auth.GET("/user/refresh", jwtAuth.RefreshHandler)
	}
	// game routes
	game := s.router.Group("/game")
	{
		game.GET("/base/:id", s.getHiragana)
		game.POST("/base/:id", s.hiraganaDone)
		game.GET("/ext/:id", s.getExtHiragana)
		game.POST("/ext/:id", s.extHiraganaDone)
	}

	// add default handler for page not found
	s.router.NoRoute(s.routeNotFound)

	fmt.Printf("Starting server on port %s\n", server.Addr)
	return server.ListenAndServe()
}

// get the port as a string
func (s *Server) getAddress() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Server) addJwtMiddleware(key string) (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "language helper",
		Key:         []byte(key),
		Timeout:     time.Minute * 15,
		MaxRefresh:  time.Hour,
		IdentityKey: "UserID",
		// function called on login
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return s.authenticateUser(c)
		},
		// write error message and status code
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(
				code,
				gin.H{
					"status":  code,
					"message": message,
				},
			)
		},
	})
	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
