package booking

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"alan-tabeo-test-task/src/models"
	"alan-tabeo-test-task/src/services"
)

type BookingController struct {
	bookingService services.BookingServiceI
}

// NewBookingController creates and returns a new instance of the BookingController.
func NewBookingController(
	bookingService services.BookingServiceI,
) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	booking := new(models.Booking)

	if err := c.ShouldBindJSON(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := bc.bookingService.CreateBooking(c, booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, booking)
}

func (bc *BookingController) GetBookings(c *gin.Context) {
	paging := new(models.Paging)

	if err := c.ShouldBindQuery(paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	bookings, err := bc.bookingService.GetAllBookings(c, paging)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, bookings)
}
