package spacex_client

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/alandeqz/alan-spacetrouble-test-task/src/logging"
)

const (
	getUpcomingSpaceXLaunchesURL = "https://api.spacexdata.com/v5/launches"
)

var _ SpaceXClientI = &SpaceXClient{}

type SpaceXClientI interface {
	// DoFlightsExistForLaunchpadAndDate requests the SpaceX database to check
	// if the intended launchpad is already planned to be used by SpaceX for the specific date.
	DoFlightsExistForLaunchpadAndDate(launchpadID string, launchDate time.Time) (bool, error)
}

type SpaceXClient struct{}

// NewSpaceXClient creates and returns a new instance of *SpaceXClient.
func NewSpaceXClient() *SpaceXClient {
	return &SpaceXClient{}
}

type SpaceXResponse struct {
	Launchpad string    `json:"launchpad"`
	DateUTC   time.Time `json:"date_utc"`
}

// DoFlightsExistForLaunchpadAndDate requests the SpaceX database to check
// if the intended launchpad is already planned to be used by SpaceX for the specific date.
func (sxc *SpaceXClient) DoFlightsExistForLaunchpadAndDate(launchpadID string, launchDate time.Time) (bool, error) {
	response := make([]*SpaceXResponse, 0)

	req, err := http.NewRequest(http.MethodGet, getUpcomingSpaceXLaunchesURL, nil)
	if err != nil {
		slog.Error("failed to form the request to get the upcoming SpaceX launches", logging.Error, err.Error())

		return false, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to get the upcoming SpaceX launches", logging.Error, err.Error())

		return false, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		slog.Error("failed to decode the response from the SpaceX API", logging.Error, err.Error())

		return false, err
	}

	for _, r := range response {
		if r.Launchpad == launchpadID && r.DateUTC == launchDate {
			return true, nil
		}
	}

	return false, nil
}
