package booking

import (
	"time"

	"alan-tabeo-test-task/src/models"
)

type BookingRequest struct {
	ID            uint64        `json:"id" example:"9"`
	CreatedAt     time.Time     `json:"created_at" example:"2024-10-29T13:55:28.897Z"`
	UpdatedAt     time.Time     `json:"updated_at" example:"2024-10-30T12:13:37.374Z"`
	FirstName     string        `json:"first_name" example:"John"`
	LastName      string        `json:"last_name" example:"Doe"`
	Gender        models.Gender `json:"gender"`
	Birthday      time.Time     `json:"birthday" example:"1999-09-01T00:00:00Z"`
	LaunchpadID   string        `json:"launchpad_id" example:"1"`
	DestinationID string        `json:"destination_id" enums:"Mars,Moon,Pluto,Asteroid Belt,Europa,Titan,Ganymede"`
	LaunchDate    time.Time     `json:"launch_date" example:"2024-12-01T15:00:00.000Z"`
}
