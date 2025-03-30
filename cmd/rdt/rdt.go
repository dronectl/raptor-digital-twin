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
    channels := common.IPCChannels{
        Control: make(chan common.SimControl),
        PowertrainReadings: make(chan common.PowertrainReadings),
        EnvironmentReadings: make(chan common.EnvironmentReadings),
    }
    device := device.NewDevice(10, channels)
    sim := sim.NewSimCtx(100, 10, channels)
    sim.Start()
    device.Start()
}
