package models

import (
	"context"
	"log/slog"

	"gorm.io/gorm"
)

var _ BookingRepositoryI = &BookingRepository{}

type BookingRepositoryI interface {
	// Create creates a new booking.
	Create(ctx context.Context, booking *Booking) error

	// GetAll returns all bookings.
	GetAll(ctx context.Context, paging *Paging) ([]*Booking, error)

	// Delete deletes a booking.
	Delete(ctx context.Context, id uint64) error
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
		slog.Error("failed to create a new booking", slog.Error, err.Error())

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
		slog.Error("failed to get all bookings", slog.Error, err.Error())

		return nil, err
	}

	return bookings, nil
}

// Delete deletes a booking.
func (br *BookingRepository) Delete(_ context.Context, id uint64) error {
	return nil
}
