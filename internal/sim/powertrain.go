package sim

import (
	"sync"

	"github.com/dronectl/rdt/internal/sim/bldc"
	"github.com/dronectl/rdt/internal/common"
	sg "github.com/dronectl/rdt/internal/sim/signals"
)


type PowertrainCtx struct {
	m            sync.Mutex
    bldc *bldc.Bldc
    // state
	tripped        bool
	armed        bool
	speedRef     float64
    // environment signals
	tempSignal   *sg.SignalExp
    // electrical signals
	dcVoltageSignal *sg.SignalExp
	l1Signal   *sg.SignalSin
	l2Signal   *sg.SignalSin
	l3Signal   *sg.SignalSin
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
    readings := common.PowertrainReadings{}
    voltage := pt.dcVoltageSignal.Process()
    pt.bldc.ProcessState(voltage)
    pt.m.Lock()
    readings.Current = 0.0;
    readings.Voltage = voltage
    readings.Power = pt.bldc.Current * readings.Voltage
    readings.Speed = pt.bldc.Speed
    readings.Torque = pt.bldc.Torque
    readings.MechanicalPower = readings.Torque * float64(readings.Speed)
    readings.Efficiency = 100.0
    readings.Temperature = pt.tempSignal.Process()
    readings.L1Voltage = pt.l1Signal.Process()
    readings.L2Voltage = pt.l2Signal.Process()
    readings.L3Voltage = pt.l3Signal.Process()
    pt.m.Unlock()
}

func NewPowertrain(env *Environment, bldc *PowertrainBldcRatings) *PowertrainCtx {
	return &PowertrainCtx{
        bldc: bldc.NewBldc(),
		armed:    false,
		speedRef: 0.0,
        dcVoltageSignal: sg.NewSignalExp(0.1, 0.0, sg.NoiseParams{Offset: 0.1, Gain: 0.1}, 10000),
        tempSignal: sg.NewSignalExp(10, env.AmbientTemperature, sg.NoiseParams{Offset: 0.1, Gain: 0.1}, 10000),
        l1Signal: sg.NewSignalSin(12,12,100,0, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
        l2Signal: sg.NewSignalSin(12,12,100, 120, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
        l3Signal: sg.NewSignalSin(12, 12, 100, 240, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
	}
}

