package handlers

import (
	"net/http"
	"server/internal/auth"
	db "server/internal/database"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUserHandler(c *gin.Context) {
	user := db.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := s.dbManager.CreateUser(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := auth.CreateToken(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token:": token})
}

func (s *Server) GetUserHandler(c *gin.Context) {
	idAny, ok := c.Get("id")
	if !ok {
		c.String(http.StatusInternalServerError, "Not found context key")
		return
	}

	var id int
	if id, ok = idAny.(int); !ok {
		c.String(http.StatusBadRequest, "Bad token")
		return
	}

	user, err := s.dbManager.GetUser(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
