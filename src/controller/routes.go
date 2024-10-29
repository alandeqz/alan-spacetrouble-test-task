package controller

import (
	"github.com/gin-gonic/gin"

	bookingsController "alan-tabeo-test-task/src/controller/booking"
)

// RegisterRoutes registers all the routes for the application.
func RegisterRoutes(
	httpServer *gin.Engine,
	bookingsController *bookingsController.BookingController,
) {
	httpServer.POST("/bookings", bookingsController.CreateBooking)
	httpServer.GET("/bookings", bookingsController.GetBookings)
}
