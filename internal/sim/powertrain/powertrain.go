package powertrain

import (
	"sync"

	sg "github.com/dronectl/rdt/internal/sim/signals"
	common "github.com/dronectl/rdt/internal/common"
	bldc "github.com/dronectl/rdt/internal/sim/powertrain/bldc"
)


type PowertrainCtx struct {
	m            sync.Mutex
    bldc *bldc.Bldc
    // state
	tripped        bool
	armed        bool
	speedRef     float64
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

func (pt *PowertrainCtx) ProcessState() common.PowertrainReadings {
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
    readings.Temperature = 0.0
    readings.L1Voltage = pt.l1Signal.Process()
    readings.L2Voltage = pt.l2Signal.Process()
    readings.L3Voltage = pt.l3Signal.Process()
    pt.m.Unlock()
    return readings
}

func NewPowertrain(bldcParams *common.BldcParameters, propParams *common.PropellerParameters, envParams *common.EnvParameters) *PowertrainCtx {
	return &PowertrainCtx{
        bldc: bldc.NewBldc(bldcParams, propParams, envParams),
		armed:    false,
		speedRef: 0.0,
        dcVoltageSignal: sg.NewSignalExp(0.1, 0.0, sg.NoiseParams{Offset: 0.1, Gain: 0.1}, 10000),
        l1Signal: sg.NewSignalSin(12,12,100,0, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
        l2Signal: sg.NewSignalSin(12,12,100, 120, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
        l3Signal: sg.NewSignalSin(12, 12, 100, 240, sg.NoiseParams{Offset: 0.001, Gain: 0.01}, 10000),
	}
}

