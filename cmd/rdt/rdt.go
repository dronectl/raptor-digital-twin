package main

import (
    "os"
    "log"

    "github.com/dronectl/rdt/internal/common"
    "github.com/dronectl/rdt/internal/device"
    "github.com/dronectl/rdt/internal/sim"
)

func main() {
    logger := log.New(os.Stdout, "", log.Lshortfile | log.Lmicroseconds)
    logger.Println("dronectl - Raptor Digital Twin");
    // build sync primitives
    control := make(chan common.SimControl)
    powertrainReadings:= make(chan common.PowertrainReadings)
    envReadings:= make(chan common.EnvironmentReadings)
    // init device and sim emulation engines
    device := device.NewDevice(10, control, powertrainReadings, envReadings)
    sim := sim.NewSimCtx(100, 10, control, powertrainReadings, envReadings)
    sim.Start()
    // main thread of execution
    device.Start()
}
