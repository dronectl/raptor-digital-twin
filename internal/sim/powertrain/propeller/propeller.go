package propeller

import (
	"github.com/dronectl/rdt/internal/common"
	"math"
)

type PropellerCtx struct {
	params  *common.PropellerParameters
	rho     float64 // air density
	torque  float64 // Nm
	cTAlpha float64
	cTBeta  float64
	cT      float64 // coefficient of thrust
}

func omegaToRpm(omega float64) float64 {
	return omega * 30.0 / math.Pi
}

func estimateCt(p *PropellerCtx) {
	p.cT = p.params.CtGain*(p.params.Pitch/p.params.Diameter) + p.params.CtOffset
}

func (p *PropellerCtx) ComputeTorque(omega float64) float64 {
	return p.cT * p.rho * math.Pow(p.params.Diameter, 5) * math.Pow(omega, 2)
}

func NewPropeller(params *common.PropellerParameters, rho float64) *PropellerCtx {
	return &PropellerCtx{
		params: params,
		rho:    rho,
	}
}
