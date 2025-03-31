package main

import (
	"log/slog"
	"os"

	"github.com/dronectl/rdt/internal/common"
	"github.com/dronectl/rdt/internal/device"
	"github.com/dronectl/rdt/internal/sim"
)

func configureLogger() {
	slogOpts := &slog.HandlerOptions{Level: slog.LevelDebug}
	jsonHandler := slog.NewJSONHandler(os.Stdout, slogOpts)
	slog.SetDefault(slog.New(jsonHandler))
}

func main() {
	configureLogger()
	logger := slog.Default()
	logger.Info("dronectl - Raptor Digital Twin")
	// build sync primitives
	control := make(chan common.SimControl)
	powertrainReadings := make(chan common.PowertrainReadings)
	envReadings := make(chan common.EnvironmentReadings)
	// init device and sim emulation engines
	deviceOpts := device.DeviceOpts{UpdateFrequency: 1}
	simOpts := sim.SimOpts{UpdateFrequency: 10, SamplePrescaler: 10}
	device := device.NewDevice(control, powertrainReadings, envReadings, &deviceOpts)
	sim := sim.NewSimCtx(control, powertrainReadings, envReadings, &simOpts)
	sim.Start()
	// main thread of execution
	device.Start()
}
