package vehicle

// Code generated by github.com/andig/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/andig/evcc/api"
)

func decorateVehicle(base api.Vehicle, vehicleStatus func() (api.ChargeStatus, error), vehicleRange func() (int64, error)) api.Vehicle {
	switch {
	case vehicleRange == nil && vehicleStatus == nil:
		return base

	case vehicleRange == nil && vehicleStatus != nil:
		return &struct {
			api.Vehicle
			api.VehicleStatus
		}{
			Vehicle: base,
			VehicleStatus: &decorateVehicleVehicleStatusImpl{
				vehicleStatus: vehicleStatus,
			},
		}

	case vehicleRange != nil && vehicleStatus == nil:
		return &struct {
			api.Vehicle
			api.VehicleRange
		}{
			Vehicle: base,
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case vehicleRange != nil && vehicleStatus != nil:
		return &struct {
			api.Vehicle
			api.VehicleRange
			api.VehicleStatus
		}{
			Vehicle: base,
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
			VehicleStatus: &decorateVehicleVehicleStatusImpl{
				vehicleStatus: vehicleStatus,
			},
		}
	}

	return nil
}

type decorateVehicleVehicleRangeImpl struct {
	vehicleRange func() (int64, error)
}

func (impl *decorateVehicleVehicleRangeImpl) Range() (int64, error) {
	return impl.vehicleRange()
}

type decorateVehicleVehicleStatusImpl struct {
	vehicleStatus func() (api.ChargeStatus, error)
}

func (impl *decorateVehicleVehicleStatusImpl) Status() (api.ChargeStatus, error) {
	return impl.vehicleStatus()
}