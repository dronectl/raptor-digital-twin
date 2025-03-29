package main

import (
    "fmt"

    "github.com/dronectl/rdt/internal/device"
    "github.com/dronectl/rdt/internal/sim"
)

func main() {
   fmt.Println("dronectl Raptor Digital Twin");
   device := device.Device{Name: "MyDevice", UUID: "1234"}
   // build power train instance
   // load into simulation backend
   sim.Start(device)
}
