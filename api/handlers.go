package api

import (
	"fmt"
	"net/http"

	"github.com/KleinSpeedy/language-helper-backend/database"
	"github.com/KleinSpeedy/language-helper-backend/datatypes"
	"github.com/gin-gonic/gin"
)

var (
	dbc *database.Controller
	err error
)

// create new user with pw and username
// header returns 200 if ok, 400 on error
func (s *Server) createUser(c *gin.Context) {
	var json datatypes.User

	// retrieve json information
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
			},
		)
		return
	}

	// check if user exists
	err, ok := dbc.UserExists(json.Username)
	if err != nil || ok {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
			},
		)
		return
	}
	// new user, hash and salt pw
	err, ok = dbc.SaveNewUser(json.Username, json.Password)
	if err != nil || !ok {
		// Only send 400 for security reasons
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
			},
		)
		return
	}
	// creation successfull
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
		},
	)
}

func (s *Server) authenticateUser(c *gin.Context) (*datatypes.User, error) {
	var json datatypes.User
	// retrieve json information
	if err := c.ShouldBind(&json); err != nil {
		return nil, fmt.Errorf("Invalid Input")
	}
	// authenticate user
	err, ok := dbc.PasswordMatches(json.Username, json.Password)
	if err != nil || !ok {
		return nil, fmt.Errorf("Wrong password")
	}
	return &json, nil
}

// returns status not found and error message if
// route has no handler registered
func (s *Server) routeNotFound(c *gin.Context) {
	c.JSON(
		http.StatusNotFound,
		gin.H{
			"status":  http.StatusNotFound,
			"message": "Page not found",
		},
	)
}

func (s *Server) getHiragana(c *gin.Context) {

}

func (s *Server) hiraganaDone(c *gin.Context) {

}

func (s *Server) getExtHiragana(c *gin.Context) {

}

func (s *Server) extHiraganaDone(c *gin.Context) {

}
