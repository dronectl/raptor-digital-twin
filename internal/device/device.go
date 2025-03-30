package device

import (
    "os"
    "log"
    "time"
    "github.com/google/uuid"

    "github.com/dronectl/rdt/internal/common"
)

type Device struct {
    logger *log.Logger

    name string
    uuid string
    updateFrequency uint32 // Hz

    // channels
    control chan<- common.SimControl
    powertrainReadings <-chan common.PowertrainReadings
    envReadings <-chan common.EnvironmentReadings
}

func (d *Device) runner() {
    defer close(d.control)
    for {
        select {
            case ptReadings := <- d.powertrainReadings:
                d.logger.Println("received powertrain readings", ptReadings)
            case envReadings := <- d.envReadings:
                d.logger.Println("received env readings", envReadings)
        }
        time.Sleep(time.Duration(1000/d.updateFrequency) * time.Millisecond)
    }
}

func (d *Device) Start() {
    d.logger.Println("Starting device emulator")
    d.runner()
}

func NewDevice(updateFrequency uint32, control chan<- common.SimControl, powertrainReadings <-chan common.PowertrainReadings, envReadings <-chan common.EnvironmentReadings) *Device {
    device := Device{
        name: "hungry-velociraptor",
        uuid: uuid.NewString(),
        logger: log.New(os.Stdout, "", log.Lshortfile | log.Lmicroseconds),
        updateFrequency: updateFrequency,
        control: control,
        powertrainReadings: powertrainReadings,
        envReadings: envReadings,
    }
    device.logger.Println("New device created with UUID", device.uuid)
    return &device
}
