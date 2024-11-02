package models

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"gorm.io/gorm"

	serviceErrors "github.com/alandeqz/alan-spacetrouble-test-task/src/errors"
	"github.com/alandeqz/alan-spacetrouble-test-task/src/logging"
)

var _ BookingRepositoryI = &BookingRepository{}

type BookingRepositoryI interface {
	// Create creates a new booking.
	Create(ctx context.Context, booking *Booking) error

	// GetAll returns all bookings.
	GetAll(ctx context.Context, paging *Paging) ([]*Booking, error)

	// Delete deletes a booking.
	Delete(ctx context.Context, id uint64) error

	// FindBookingForLaunchpad finds and returns the booking for the specific launchpad for the specific date.
	FindBookingForLaunchpad(ctx context.Context, launchpadID string, launchDate time.Time) (*Booking, error)

	// FindBookingsForDestination finds and returns the list of bookings
	// flying to the specific destination for the passed timeframe.
	FindBookingsForDestination(ctx context.Context, destinationID string, from, to time.Time) ([]*Booking, error)
}

type BookingRepository struct {
	*gorm.DB
}

// NewBookingRepository creates and returns a new instance of *BookingRepository.
func NewBookingRepository(
	db *gorm.DB,
) *BookingRepository {
	return &BookingRepository{
		db,
	}
}

// Create creates a new booking.
func (br *BookingRepository) Create(_ context.Context, booking *Booking) error {
	if err := br.DB.Create(booking).Error; err != nil {
		slog.Error("failed to create a new booking", logging.Error, err.Error())

		return err
	}

	return nil
}

// GetAll returns all bookings.
func (br *BookingRepository) GetAll(_ context.Context, paging *Paging) ([]*Booking, error) {
	bookings := make([]*Booking, 0)

	db := br.DB

	if paging.Limit != nil {
		db = db.Limit(*paging.Limit)
	}

	if paging.Offset != nil {
		db = db.Offset(*paging.Offset)
	}

	if err := db.Find(&bookings).Error; err != nil {
		slog.Error("failed to get all bookings", logging.Error, err.Error())

		return nil, err
	}

	return bookings, nil
}

// Delete deletes a booking.
func (br *BookingRepository) Delete(_ context.Context, id uint64) error {
	res := br.DB.Where("id = ?", id).Delete(&Booking{})

	if res.Error != nil {
		slog.Error("failed to delete the booking", logging.Error, res.Error.Error())

		return res.Error
	}

	if res.RowsAffected == 0 {
		return serviceErrors.ErrBookingNotFound
	}

	return nil
}

// FindBookingForLaunchpad finds and returns the booking for the specific launchpad for the specific date.
func (br *BookingRepository) FindBookingForLaunchpad(_ context.Context, launchpadID string, launchDate time.Time) (*Booking, error) {
	booking := new(Booking)

	if err := br.DB.
		Where("launch_date >= ?", time.Date(launchDate.Year(), launchDate.Month(), launchDate.Day(), 0, 0, 0, 0, time.UTC)).
		Where("launch_date <= ?", time.Date(launchDate.Year(), launchDate.Month(), launchDate.Day(), 23, 59, 59, 999, time.UTC)).
		Where("launchpad_id = ?", launchpadID).
		First(booking).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return booking, nil
}

// FindBookingsForDestination finds and returns the list of bookings
// flying to the specific destination for the passed timeframe.
func (br *BookingRepository) FindBookingsForDestination(_ context.Context, destinationID string, from, to time.Time) ([]*Booking, error) {
	bookings := make([]*Booking, 0)

	if err := br.DB.
		Where("launch_date >= ?", from).
		Where("launch_date <= ?", to).
		Where("destination_id = ?", destinationID).
		Find(&bookings).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return bookings, nil
}
