package api

import (
	"net/http"

	"github.com/KleinSpeedy/language-helper-backend/types"
	"github.com/gin-gonic/gin"
)

// Get a user with by id
func (s *Server) GetUser(c *gin.Context) {
	me := types.NewUser("Jonas", 1)

	c.JSON(http.StatusOK, gin.H{
		"name": me.Username,
		"id":   me.ID,
	})
}

// create new user
func (s *Server) CreateUser(c *gin.Context) {
	c.JSON(200, nil)
}

// update user
func (s *Server) UpdateUser(c *gin.Context) {
	c.JSON(200, nil)
}

// delete user by id
func (s *Server) DeleteUser(c *gin.Context) {
	c.JSON(200, nil)
}
