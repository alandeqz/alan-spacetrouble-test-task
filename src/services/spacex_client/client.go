package spacex_client

import (
	"time"
)

var _ SpaceXClientI = &SpaceXClient{}

type SpaceXClientI interface {
	// DoFlightsExistForLaunchpadAndDate requests the SpaceX database to check
	// if the intended launchpad is already planned to be used by SpaceX for the specific date.
	DoFlightsExistForLaunchpadAndDate(launchpadID string, launchDate time.Time) (bool, error)
}

type SpaceXClient struct {
}

// NewSpaceXClient creates and returns a new instance of *SpaceXClient.
func NewSpaceXClient() *SpaceXClient {
	return &SpaceXClient{}
}

// DoFlightsExistForLaunchpadAndDate requests the SpaceX database to check
// if the intended launchpad is already planned to be used by SpaceX for the specific date.
func (sxc *SpaceXClient) DoFlightsExistForLaunchpadAndDate(launchpadID string, launchDate time.Time) (bool, error) {
	//TODO implement me
	panic("implement me")
}
