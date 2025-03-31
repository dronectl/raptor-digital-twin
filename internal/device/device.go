package device

import (
	"time"
	"log/slog"
	"github.com/google/uuid"

	"github.com/dronectl/rdt/internal/common"
	com "github.com/dronectl/rdt/internal/device/com"
)

type DeviceOpts struct {
	UpdateFrequency uint32
}

type DeviceVersion struct {
    Major uint32
    Minor uint32
    Patch uint32
}

type Device struct {
	logger *slog.Logger

	name            string
	uuid            string
	updateFrequency uint32 // Hz
    hwVersion      DeviceVersion
    fwVersion      DeviceVersion

    uhciHandle *com.UHCIHandle
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
    d.uhciHandle.Start()
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
        uhciHandle:         com.NewUHCIHandle(),
	}
	device.logger.Info("New device created", "uuid", device.uuid)
	return &device
}
