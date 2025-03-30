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

    control chan common.SimControl
    channels common.IPCChannels
}

func (d *Device) runner() {
    for {
        select {
            case ptReadings := <- d.channels.PowertrainReadings:
                d.logger.Println("received powertrain readings", ptReadings)
            case envReadings := <- d.channels.EnvironmentReadings:
                d.logger.Println("received env readings", envReadings)
        }
        time.Sleep(time.Duration(1000/d.updateFrequency) * time.Millisecond)
    }
}

func (d *Device) Start() {
    d.runner()
}

func NewDevice(updateFrequency uint32, channels common.IPCChannels) *Device {
    return &Device{
        name: "hungry-velociraptor",
        uuid: uuid.NewString(),
        logger: log.New(os.Stdout, "", log.Lshortfile | log.Lmicroseconds),
        updateFrequency: updateFrequency,
        channels: channels,
    }
}
