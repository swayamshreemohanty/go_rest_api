package routes

import (
	"go_rest_api/models"
	"go_rest_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context)  {
	id,err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEvent(id)

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func getEvents(c *gin.Context)  {

	events , err:= models.GetAllEvents()

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context)  {
	// initialize an empty event
	var event models.Event

	// Bind the event from the request
	err := c.ShouldBindJSON(&event)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the userId from the context
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	event.UserId = userId.(int64)

	// Save the event to the database
	event.Save()

	c.JSON(http.StatusCreated, event)
}

func updateEvent(c *gin.Context)  {
	id,err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	//Check the event exists
	event, err := models.GetEvent(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	
	//Check if the user is the owner of the event
	if ok:=utils.CheckUserOwnership(c, event.UserId); !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to update this event"}) 
		return
	}

	// initialize an empty event
	var updatedEvent models.Event

	// Bind the event from the request
	err=c.ShouldBindJSON(&updatedEvent)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the ID of the event
	updatedEvent.ID = int(id)

	err = updatedEvent.Update()

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context)  {
	id,err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Check the event exists
	event, err := models.GetEvent(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}


	//Check if the user is the owner of the event
	if ok:=utils.CheckUserOwnership(c, event.UserId); !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized to delete this event"})
		return
	}

	// Delete the event
	err = event.Delete()

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}