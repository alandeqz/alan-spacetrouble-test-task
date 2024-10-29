package services

import (
	"context"

	"alan-tabeo-test-task/src/models"
)

var _ BookingServiceI = &BookingService{}

type BookingServiceI interface {
	// CreateBooking creates a new booking.
	CreateBooking(ctx context.Context, booking *models.Booking) error

	// GetAllBookings returns all bookings.
	GetAllBookings(ctx context.Context, paging *models.Paging) ([]*models.Booking, error)

	// DeleteBooking deletes a booking.
	DeleteBooking(ctx context.Context, id uint64) error
}

type BookingService struct {
	bookingsRepository models.BookingRepositoryI
}

// NewBookingService creates and returns a new instance of the BookingService.
func NewBookingService(
	bookingsRepository models.BookingRepositoryI,
) *BookingService {
	return &BookingService{
		bookingsRepository: bookingsRepository,
	}
}

// CreateBooking creates a new booking.
func (bs *BookingService) CreateBooking(ctx context.Context, booking *models.Booking) error {
	return bs.bookingsRepository.Create(ctx, booking)
}

// GetAllBookings returns all bookings.
func (bs *BookingService) GetAllBookings(ctx context.Context, paging *models.Paging) ([]*models.Booking, error) {
	return bs.bookingsRepository.GetAll(ctx, paging)
}

// DeleteBooking deletes a booking.
func (bs *BookingService) DeleteBooking(ctx context.Context, id uint64) error {
	return bs.bookingsRepository.Delete(ctx, id)
}
