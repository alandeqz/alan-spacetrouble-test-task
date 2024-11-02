package models

import (
	"time"
)

type Booking struct {
	ID            uint64    `json:"id" gorm:"primaryKey;column:id" example:"9"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at" example:"2024-10-29T13:55:28.897Z"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at" example:"2024-10-30T12:13:37.374Z"`
	FirstName     string    `json:"first_name" gorm:"column:first_name" example:"John"`
	LastName      string    `json:"last_name" gorm:"column:last_name" example:"Doe"`
	Gender        Gender    `json:"gender" gorm:"column:gender"`
	Birthday      time.Time `json:"birthday" gorm:"column:birthday" example:"1999-09-01T00:00:00Z"`
	LaunchpadID   string    `json:"launchpad_id" gorm:"column:launchpad_id" example:"1"`
	DestinationID string    `json:"destination_id" gorm:"column:destination_id" example:"2"`
	LaunchDate    time.Time `json:"launch_date" gorm:"column:launch_date" example:"2024-12-01T15:00:00.000Z"`
}

// TableName returns the table name of the Booking model.
func (b *Booking) TableName() string {
	return "alan_tabeo_test_task.bookings"
}
