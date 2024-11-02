package services

import (
	"context"
	"log/slog"
	"time"

	"alan-tabeo-test-task/src/errors"
	"alan-tabeo-test-task/src/logging"
	"alan-tabeo-test-task/src/models"
	"alan-tabeo-test-task/src/services/spacex_client"
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
	spaceXClient       spacex_client.SpaceXClientI
}

// NewBookingService creates and returns a new instance of the BookingService.
func NewBookingService(
	bookingsRepository models.BookingRepositoryI,
	spaceXClient spacex_client.SpaceXClientI,
) *BookingService {
	return &BookingService{
		bookingsRepository: bookingsRepository,
		spaceXClient:       spaceXClient,
	}
}

// CreateBooking creates a new booking.
func (bs *BookingService) CreateBooking(ctx context.Context, booking *models.Booking) error {
	bookingForLaunchpad, err := bs.bookingsRepository.FindBookingForLaunchpad(ctx, booking.LaunchpadID, booking.LaunchDate)
	if err != nil {
		slog.Error("failed to check the existing bookings for the launchpad", logging.Error, err.Error())

		return err
	}

	if bookingForLaunchpad == nil {
		return errors.ErrBookingAlreadyExists
	}

	launchpadUsedBySpaceX, err := bs.spaceXClient.DoFlightsExistForLaunchpadAndDate(booking.LaunchpadID, booking.LaunchDate)
	if err != nil {
		slog.Error("failed to check the planned flights for SpaceX", logging.Error, err.Error())

		return err
	}

	if launchpadUsedBySpaceX {
		return errors.ErrBookingAlreadyExists
	}

	from, to := getWeekRange(booking.LaunchDate)

	bookingsForDestination, err := bs.bookingsRepository.FindBookingsForDestination(ctx, booking.DestinationID, from, to)
	if err != nil {
		slog.Error("failed to check the existing booking for the destination")

		return err
	}

	if len(bookingsForDestination) == 0 {
		return errors.ErrBookingAlreadyExists
	}

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

// getWeekRange returns the start and end times of the current week.
func getWeekRange(launchDate time.Time) (time.Time, time.Time) {
	launchDateUTC := launchDate.UTC()

	weekday := int(launchDateUTC.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	monday := launchDateUTC.AddDate(0, 0, -weekday+1).Truncate(24 * time.Hour).UTC()

	sunday := monday.AddDate(0, 0, 6).
		Add(time.Hour*23 + time.Minute*59 + time.Second*59 + time.Nanosecond*999999999)

	return monday, sunday
}
