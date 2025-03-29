package bldc

import (
	"sync"

	"github.com/dronectl/rdt/internal/common"
	"github.com/dronectl/rdt/internal/sim/propeller"
)

type Bldc struct {
    m sync.Mutex
    params common.BldcParameters
    propeller *propeller.PropellerCtx
    Current float64 // A
    Torque float64
    Speed uint32
    kt float64 // Nm/A
    ke float64 // V/(rad/s)
}

func (b *Bldc) ProcessState(voltage float64) {
    // TODO: implement newton raphson solver
    b.m.Lock()
    b.Current = 0.0
    b.Torque = 0.0
    b.Speed = 0
    b.m.Unlock()
}

func NewBldc(bldcParams common.BldcParameters, propParams common.PropellerParameters, kt float64, ke float64, rho float64, cTAlpha float64, cTBeta float64) *Bldc {
    return &Bldc{
        params: bldcParams,
        propeller: propeller.NewPropeller(propParams, rho, cTAlpha, cTBeta),
        kt: kt,
        ke: ke,
    }
}

