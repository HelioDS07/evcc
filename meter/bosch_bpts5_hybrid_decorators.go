package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateBoschBpts5Hybrid(base api.Meter, meterEnergy func() (float64, error), battery func() (float64, error)) api.Meter {
	switch {
	case battery == nil && meterEnergy == nil:
		return base

	case battery == nil && meterEnergy != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
		}{
			Meter: base,
			MeterEnergy: &decorateBoschBpts5HybridMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && meterEnergy == nil:
		return &struct {
			api.Meter
			api.Battery
		}{
			Meter: base,
			Battery: &decorateBoschBpts5HybridBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && meterEnergy != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateBoschBpts5HybridBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateBoschBpts5HybridMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decorateBoschBpts5HybridBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateBoschBpts5HybridBatteryImpl) SoC() (float64, error) {
	return impl.battery()
}

type decorateBoschBpts5HybridMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateBoschBpts5HybridMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}