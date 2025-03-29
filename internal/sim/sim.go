
package sim

import (
    "time"
    "github.com/dronectl/rdt/sim/common"
)

type SimCtx struct {
    updateFrequency uint32 // Hz
    powertrainReadings chan common.PowertrainReadings
    powertrainCtx PowertrainCtx
}

func (s *SimCtx) runner() {
    for {
        s.powertrainCtx.ProcessState()
        s.powertrainReadings <- s.powertrainCtx.readings
        time.Sleep(time.Duration(1000/s.updateFrequency) * time.Millisecond)
    }
}

func (s *SimCtx) Start(updateFrequency uint32) {
    go s.runner()
}

func NewSimCtx(updateFrequency uint32, powertrainChannel chan common.PowertrainReadings) *SimCtx {
    return &SimCtx{
        updateFrequency: updateFrequency,
        powertrainReadings: powertrainChannel,
    }
}

