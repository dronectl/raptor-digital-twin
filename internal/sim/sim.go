package sim

import (
    "os"
	"log"
	"time"

	common "github.com/dronectl/rdt/internal/common"
	powertrain "github.com/dronectl/rdt/internal/sim/powertrain"
)

type SimCtx struct {
    logger *log.Logger
    // configurability layer
    envParameters common.EnvParameters
    propellerParameters common.PropellerParameters
    bldcParameters common.BldcParameters

    updateFrequency uint32 // Hz
    samplePrescaler uint8
    powertrainCtx *powertrain.PowertrainCtx
    prescaleCounter uint8

    // channels
    control <-chan common.SimControl
    powertrainReadings chan<- common.PowertrainReadings
    envReadings chan<- common.EnvironmentReadings
}

func (s *SimCtx) runner() {
    defer close(s.powertrainReadings)
    defer close(s.envReadings)
    s.prescaleCounter = s.samplePrescaler
    for {
        cmd := common.SIM_CMD_NULL
        // non-blocking read
        select {
            case cmd = <- s.control:
                s.logger.Println("Sim: Received command ", cmd)
            default:
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
    s.logger.Println("Exited simulation")
}

func (s *SimCtx) Start() {
    s.logger.Println("Starting simulation backend")
    go s.runner()
}

func NewSimCtx(updateFrequency uint32, samplePrescaler uint8, control <-chan common.SimControl, powertrainReadings chan<- common.PowertrainReadings, envReadings chan<- common.EnvironmentReadings) *SimCtx {
    simCtx:= SimCtx{
        logger: log.New(os.Stdout, "", log.Lshortfile | log.Lmicroseconds),
        bldcParameters: common.DefaultBldcParameters,
        envParameters: common.DefaultEnvParameters,
        propellerParameters: common.DefaultPropellerParameters,
        samplePrescaler: samplePrescaler,
        updateFrequency: updateFrequency,
        control: control,
        powertrainReadings: powertrainReadings,
        envReadings: envReadings,
    }
    // pass references to the configurability layer so changes are parameterically applied
    simCtx.powertrainCtx = powertrain.NewPowertrain(&simCtx.bldcParameters, &simCtx.propellerParameters, &simCtx.envParameters)
    return &simCtx
}

