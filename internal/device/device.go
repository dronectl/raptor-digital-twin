package device

import (
	"github.com/google/uuid"
	"log/slog"
	"time"

	"github.com/dronectl/rdt/internal/common"
)

type DeviceOpts struct {
	UpdateFrequency uint32
}

type Device struct {
	logger *slog.Logger

	name            string
	uuid            string
	updateFrequency uint32 // Hz

	// channels
	control            chan<- common.SimControl
	powertrainReadings <-chan common.PowertrainReadings
	envReadings        <-chan common.EnvironmentReadings
}

func (d *Device) runner() {
	defer close(d.control)
	for {
		select {
		case ptReadings := <-d.powertrainReadings:
			d.logger.Info("Received powertrain payload", "powertrain", ptReadings)
		case envReadings := <-d.envReadings:
			d.logger.Info("Received environment payload", "environment", envReadings)
		}
		time.Sleep(time.Duration(1000/d.updateFrequency) * time.Millisecond)
	}
}

func (d *Device) Start() {
	d.logger.Info("Starting device emulator")
	d.runner()
}

func NewDevice(control chan<- common.SimControl, powertrainReadings <-chan common.PowertrainReadings, envReadings <-chan common.EnvironmentReadings, opts *DeviceOpts) *Device {
	device := Device{
		name:               "hungry-velociraptor",
		uuid:               uuid.NewString(),
		logger:             slog.Default(),
		updateFrequency:    opts.UpdateFrequency,
		control:            control,
		powertrainReadings: powertrainReadings,
		envReadings:        envReadings,
	}
	device.logger.Info("New device created", "uuid", device.uuid)
	return &device
}
