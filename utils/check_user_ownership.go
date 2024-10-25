package utils

import (
	"github.com/gin-gonic/gin"
)

//Check if the user owns the resource
func CheckUserOwnership(c *gin.Context, ownerId int64) bool{
	// Get the userId from the context
	userId, exists := c.Get("userId")
	if !exists {
		return false
	}
	// Parse the userId to int64
	parsedUserId := userId.(int64)

	//Check if the user is the owner of the event
	return parsedUserId == ownerId
}