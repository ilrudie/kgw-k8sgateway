package syncer

import "github.com/solo-io/go-utils/errors"

var (
	DashboardIdNotFound = func(dashboardJson string) error {
		return errors.Errorf("Could not find dashboard id in dashboard json: %s", dashboardJson)
	}
	DashboardIdConversionError = func(rawDashboardId interface{}) error {
		return errors.Errorf("Could not convert %v to a float64", rawDashboardId)
	}
)
