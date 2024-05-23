package routes

import (
	"net/http"
	"strconv"

	"github.com/alerebal/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	userId := ctx.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the event",
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}
	if userId != event.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update the event"})
	}
	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data required"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "The event has been updated", "event": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
	}
	if userId != event.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete the event"})
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
