package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	bookingsController "alan-tabeo-test-task/src/controller/booking"
	_ "alan-tabeo-test-task/src/docs"
)

// RegisterRoutes registers all the routes for the application.
func RegisterRoutes(
	httpServer *gin.Engine,
	bookingsController *bookingsController.BookingController,
) {
	v1HTTPServer := httpServer.Group("/v1")

	v1HTTPServer.POST("/bookings", bookingsController.CreateBooking)
	v1HTTPServer.GET("/bookings", bookingsController.GetBookings)
	v1HTTPServer.DELETE("/bookings/:id", bookingsController.DeleteBooking)
	v1HTTPServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
