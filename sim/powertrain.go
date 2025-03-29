package sim

import (
	"sync"
    "math"

	sg "github.com/dronectl/raptor-digital-twin/sim/signals"
	tz "github.com/dronectl/raptor-digital-twin/sim/tripzones"
)

type BldcParameters struct {
    rpm uint32
    max_voltage float64 // V
    kv float64 // rpm / V
    resistance float64 // Ohm
}

type PropellerParameters struct {
    diameter float64 // m
    pitch float64 // m
}

type Propeller struct {
    PropellerParameters
    rho float64 // air density
    torque float64 // Nm
    cT float64 // coefficient of thrust
}

func omegaToRpm(omega float64) float64 {
    return omega * 30.0 / math.Pi
}

func (p *Propeller) ComputeTorque(omega float64) float64 {
    return p.cT * p.rho * math.Pow(p.diameter, 5) * math.Pow(omega, 2)
}

type Bldc struct {
    BldcParameters
    current float64 // A
    k_t float64 // Nm/A
    k_e float64 // V/(rad/s)
}

type PowertrainReadings struct {
	armed   bool
	thrust  float64 // Nm
	torque  float64 // Nm
	rpm  uint32 // RPM
    mechanicalPower float64 // W
    // Power train
	dcCurrent float64 // A
	dcVoltage float64 // V
	dcPower float64 // W
    // ESC output signals
	l1Voltage float64 // V
	l2Voltage float64 // V
	l3Voltage float64 // V
    efficiency float64 // %
	temperature  float64 // C
}

type PowertrainCtx struct {
	m            sync.Mutex
    // state
    bldcRatings PowertrainBldcRatings
	tripped        bool
	armed        bool
	speedRef     float64
    readings    PowertrainReadings
    // tripzones
	currentTz   *tz.Tripzone[float64]
	tempTz       *tz.Tripzone[float64]
    // electrical signals
	dcCurrentSignal *sg.SignalExp
	dcVoltageSignal *sg.SignalExp
	tempSignal   *sg.SignalExp
	l1Signal   *sg.SignalSin
	l2Signal   *sg.SignalSin
	l3Signal   *sg.SignalSin
}


func computeF(i float64) float64 {
    w = (V - (** r_m)) / k_e
}

func (pt *PowertrainCtx) Arm(arm bool) {
	pt.m.Lock()
	pt.armed = arm
	pt.m.Unlock()
}

func (pt *PowertrainCtx) SetSpeed(duty float64) {
	pt.m.Lock()
	pt.speedRef = duty
	pt.dcCurrentSignal.Ref = duty * pt.bldcRatings.current
    if pt.armed == true {
        pt.dcVoltageSignal.Enable(true)
        pt.dcCurrentSignal.Enable(true)
        pt.l1Signal.Enable(true)
        pt.l2Signal.Enable(true)
        pt.l3Signal.Enable(true)
    }
	pt.m.Unlock()
}

func (pt *PowertrainCtx) Init() {
    pt.tempSignal.Enable(true)
}

// runs the pt simulation for a single time step
func (pt *PowertrainCtx) ProcessState() {
    pt.m.Lock()
    // heat := pt.readings.dcCurrent ** 2 * pt.bldcRatings
    pt.readings.dcCurrent = pt.dcCurrentSignal.Process()
    pt.readings.dcVoltage = pt.dcVoltageSignal.Process()
    pt.readings.dcPower = pt.readings.dcCurrent * pt.readings.dcVoltage
    pt.readings.rpm = uint32(pt.bldcRatings.kv * pt.readings.dcVoltage)
    pt.readings.torque = pt.readings.dcCurrent * pt.bldcRatings.kv
    pt.readings.mechanicalPower = pt.readings.torque * float64(pt.readings.rpm)
    losses := math.Pow(pt.readings.dcCurrent, 2) * 1.0
    pt.readings.efficiency = (pt.readings.mechanicalPower + losses) / pt.readings.dcPower;
    pt.readings.temperature = pt.tempSignal.Process()
    pt.readings.l1Voltage = pt.l1Signal.Process()
    pt.readings.l2Voltage = pt.l2Signal.Process()
    pt.readings.l3Voltage = pt.l3Signal.Process()
    pt.m.Unlock()
}

func (pt *PowertrainCtx) GetReadings() *PowertrainReadings {
	return &pt.readings
}

func NewPowertrain(env *Environment, bldc *PowertrainBldcRatings) *PowertrainCtx {
	return &PowertrainCtx{
        bldcRatings: *bldc,
		armed:    false,
		speedRef: 0.0,
		currentTz:  tz.New(0.0, 0.0),
		tempTz:   tz.New(0.0, 0.0),
        dcCurrentSignal: sg.NewSignalExp(0.1, 0.0, sg.NoiseParams{Offset: 0.1, Gain: 0.1}, 10000),
        dcVoltageSignal: sg.NewSignalExp(0.1, 0.0, sg.NoiseParams{Offset: 0.1, Gain: 0.1}, 10000),
        tempSignal: sg.NewSignalExp(10, env.AmbientTemperature, sg.NoiseParams{Offset: 0.1, Gain: 0.1}, 10000),
        l1Signal: sg.NewSignalSin(12,12,100,0, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
        l2Signal: sg.NewSignalSin(12,12,100, 120, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
        l3Signal: sg.NewSignalSin(12, 12, 100, 240, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
	}
}


