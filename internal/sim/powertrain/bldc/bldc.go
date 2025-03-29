package bldc

import (
	"sync"

	"github.com/dronectl/rdt/internal/common"
	"github.com/dronectl/rdt/internal/sim/powertrain/propeller"
)

type Bldc struct {
    m sync.Mutex
    params *common.BldcParameters
    env *common.EnvParameters
    propeller *propeller.PropellerCtx
    Current float64 // A
    Torque float64
    Speed uint32
}

func (b *Bldc) ProcessState(voltage float64) {
    // TODO: implement newton raphson solver
    b.m.Lock()
    b.Current = 0.0
    b.Torque = 0.0
    b.Speed = 0
    b.m.Unlock()
}

func NewBldc(bldcParams *common.BldcParameters, propParams *common.PropellerParameters, envParams *common.EnvParameters) *Bldc {
    return &Bldc{
        params: bldcParams,
        propeller: propeller.NewPropeller(propParams, envParams.Rho),
        env: envParams,
    }
}

