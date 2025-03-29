package device

import (
    "fmt"
    "time"

    "github.com/dronectl/rdt/internal/common"
)


type Device struct {
    Name string // user configured name
    UUID string // serial id
    control chan common.SimControl
    updateFrequency uint32 // Hz
    powertrainReadings chan common.PowertrainReadings
}

func (d *Device) runner() {
    // start a new device
    for {
        // fetch data
        select {
            case ptReadings := <- d.powertrainReadings:
                fmt.Println("received powertrain readings", ptReadings)
        }
        time.Sleep(time.Duration(1000/d.updateFrequency) * time.Millisecond)
    }
}

// main thread of execution
func (d *Device) Start() {
    // start keep alive message stream (1 Hz)
    // fetch conditions data
    d.runner()
}

func NewDevice(updateFrequency uint32, powertrainChannel chan common.PowertrainReadings, control chan common.SimControl) *Device {
    return &Device{
        updateFrequency: updateFrequency,
        control: control,
        powertrainReadings: powertrainChannel,
    }
}
