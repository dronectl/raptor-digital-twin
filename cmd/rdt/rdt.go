package main

import (
    "fmt"

    "github.com/dronectl/rdt/internal/common"
    "github.com/dronectl/rdt/internal/device"
    "github.com/dronectl/rdt/internal/sim"
)

func main() {
   fmt.Println("dronectl Raptor Digital Twin");
   controlChan := make(chan common.SimControl)
   powertrainReadingsChan := make(chan common.PowertrainReadings)
   device := device.NewDevice(1, powertrainReadingsChan, controlChan)
   sim := sim.NewSimCtx(10, 5, powertrainReadingsChan, controlChan)
   sim.Start()
   device.Start()
}
