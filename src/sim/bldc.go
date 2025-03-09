
package sim

import (
    "sync"
    "math/rand"

    tz "github.com/dronectl/raptor-digital-twin/sim/tripzones"
    sg "github.com/dronectl/raptor-digital-twin/sim/signals"
)

type BldcReadings struct {
    armed bool
    thrust float64
    torque float64
    current float64
    voltage float64
}

type BldcCtx struct {
    m sync.Mutex
    armed bool
    speedRef float64
    armThreshold float64
    speedTz *tz.Tripzone[float64]
    thrustSignal *sg.Signal
    tempTz *tz.Tripzone[float64]
    tempSignal *sg.Signal
}

const (
    BldcThrustMin = 0.0 // N
    BldcThrustMax = 100.0 // N
)

func (bldc *BldcCtx)Arm(arm bool) {
    bldc.m.Lock()
    bldc.armed = arm
    bldc.m.Unlock()
}

func (bldc *BldcCtx)SetSpeedReference(ref float64) {
    bldc.m.Lock()
    bldc.speedRef = ref
    bldc.thrustSignal.SetReference(bldc.speedRef * 100.0)
    bldc.m.Unlock()
}

func New() *BldcCtx {
    return &BldcCtx{
        armed: false,
        speedRef: 0.0,
        speedTz: tz.New(0.0, 0.0),
        tempTz: tz.New(0.0, 0.0),
    }
}

func (bldc *BldcCtx)GetReadings() *BLDCReadings {
    thrust := bldc.speedSignal.Process()
    return &BLDCReadings{
        armed: bldc.armed,
        thrust: rand.Float64(),
    }
}

