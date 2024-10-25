package routes

import (
	"go_rest_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Register for the event
func registerForEvent(c *gin.Context)  {
	// Get the event ID from the URL
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the userId from the context
	requestByUserId := c.GetInt64("userId")

	//check the event exists
	event, err := models.GetEvent(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}

	// Register the user for the event
	err = event.RegisterForEvent(requestByUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for the event"})
}

// Cancel the registration for the event
func cancelRegistrationForEvent(c *gin.Context)  {
	// Get the event ID from the URL
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the userId from the context
	requestByUserId := c.GetInt64("userId")

	//check the event exists
	event, err := models.GetEvent(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}

	// Unregister the user from the event
	err = event.UnregisterFromEvent(requestByUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from the event"})
}