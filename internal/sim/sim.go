package sim

import (
	"fmt"
	"time"

	common "github.com/dronectl/rdt/internal/common"
	powertrain "github.com/dronectl/rdt/internal/sim/powertrain"
)

type SimCtx struct {
    // configurability layer
    envParameters common.EnvParameters
    propellerParameters common.PropellerParameters
    bldcParameters common.BldcParameters

    control chan common.SimControl
    powertrainReadings chan common.PowertrainReadings
    updateFrequency uint32 // Hz
    samplePrescaler uint8
    powertrainCtx *powertrain.PowertrainCtx
    prescaleCounter uint8
}

func (s *SimCtx) runner() {
    defer close(s.powertrainReadings)
    defer close(s.control)
    s.prescaleCounter = s.samplePrescaler
    for {
        cmd := common.SIM_CMD_NULL
        select {
            case cmd := <- s.control:
                fmt.Println("received command", cmd)
        }
        if cmd == common.SIM_CMD_STOP {
            break
        }
        readings := s.powertrainCtx.ProcessState()
        // apply sample prescaling
        if s.prescaleCounter == 1 {
            s.powertrainReadings <- readings
            s.prescaleCounter = s.samplePrescaler
        } else {
            s.prescaleCounter--
        }
        time.Sleep(time.Duration(1000/s.updateFrequency) * time.Millisecond)
    }
    fmt.Println("Exited simulation")
}

func (s *SimCtx) Start() {
    fmt.Println("Starting simulation backend")
    go s.runner()
}

func NewSimCtx(updateFrequency uint32, samplePrescaler uint8, powertrainChannel chan common.PowertrainReadings, control chan common.SimControl) *SimCtx {
    simCtx:= SimCtx{
        bldcParameters: common.DefaultBldcParameters,
        envParameters: common.DefaultEnvParameters,
        propellerParameters: common.DefaultPropellerParameters,
        samplePrescaler: samplePrescaler,
        updateFrequency: updateFrequency,
        control: control,
        powertrainReadings: powertrainChannel,
    }
    // pass references to the configurability layer so changes are parameterically applied
    simCtx.powertrainCtx = powertrain.NewPowertrain(&simCtx.bldcParameters, &simCtx.propellerParameters, &simCtx.envParameters)
    return &simCtx
}

