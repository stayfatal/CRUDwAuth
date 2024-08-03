package handlers

import (
	"net/http"
	"server/internal/database"

	"github.com/gin-gonic/gin"
)

func (s *Server) AddEventHandler(c *gin.Context) {
	event := database.Event{}
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = s.dbManager.AddEvent(event)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "your event was succesfuly added"})
}

func (s *Server) AddSubscriptionHandler(c *gin.Context) {
	sub := database.Subscription{}
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ctxId, ok := c.Get("id")
	if !ok {
		c.String(http.StatusBadRequest, "Context not found")
		return
	}

	sub.UserID = ctxId.(int)

	err = s.dbManager.AddSubscription(sub)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "your succesfuly subscribed at event"})
}
