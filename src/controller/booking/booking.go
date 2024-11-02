package booking

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	serviceErrors "github.com/alandeqz/alan-spacetrouble-test-task/src/errors"
	"github.com/alandeqz/alan-spacetrouble-test-task/src/models"
	"github.com/alandeqz/alan-spacetrouble-test-task/src/services"
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

// CreateBooking
//
//	@Summary		Create the Booking
//	@Description	Create the Booking for the flight by SpaceTrouble
//	@Tags			Bookings
//	@Accept			json
//	@Produce		json
//	@Param			booking	body		BookingRequest	true	"New Booking Request"
//	@Success		201		{object}	models.Booking
//	@Failure		409		{object}	serviceErrors.GenericErrorResponse
//	@Failure		500		{object}	serviceErrors.GenericErrorResponse
//	@Router			/v1/bookings [post]
func (bc *BookingController) CreateBooking(c *gin.Context) {
	bookingRequest := new(BookingRequest)

	if err := c.ShouldBindJSON(bookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	booking := &models.Booking{
		FirstName:     bookingRequest.FirstName,
		LastName:      bookingRequest.LastName,
		Gender:        bookingRequest.Gender,
		Birthday:      bookingRequest.Birthday,
		LaunchpadID:   bookingRequest.LaunchpadID,
		DestinationID: bookingRequest.DestinationID,
		LaunchDate:    bookingRequest.LaunchDate,
	}

	if err := bc.bookingService.CreateBooking(c, booking); err != nil {
		switch {
		case errors.Is(err, serviceErrors.ErrBookingAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}
	}

	c.JSON(http.StatusCreated, booking)
}

// GetBookings
//
//	@Summary		Get the Bookings
//	@Description	Get all the Bookings for the flights by SpaceTrouble
//	@Tags			Bookings
//	@Produce		json
//	@Param			limit	query		int	false	"The limit of the response length"
//	@Param			offset	query		int	false	"The offset of the response"
//	@Success		201		{object}	[]models.Booking
//	@Failure		409		{object}	serviceErrors.GenericErrorResponse
//	@Failure		500		{object}	serviceErrors.GenericErrorResponse
//	@Router			/v1/bookings [get]
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

// DeleteBooking
//
//	@Summary		Delete the Booking
//	@Description	Delete the Booking for the flight by SpaceTrouble
//	@Tags			Bookings
//	@Produce		json
//	@Param			id	path	int	true	"The ID of the Booking"
//	@Success		200
//	@Failure		404	{object}	serviceErrors.GenericErrorResponse
//	@Failure		500	{object}	serviceErrors.GenericErrorResponse
//	@Router			/v1/bookings/{id} [delete]
func (bc *BookingController) DeleteBooking(c *gin.Context) {
	bookingIDParam := c.Param("id")

	bookingID, err := strconv.ParseUint(bookingIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err = bc.bookingService.DeleteBooking(c, bookingID); err != nil {
		switch {
		case errors.Is(err, serviceErrors.ErrBookingNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}
	}

	c.Status(http.StatusOK)
}
