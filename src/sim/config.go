
package sim

import (

    sg "github.com/dronectl/raptor-digital-twin/sim/signals"
)

var (
    // signals
    ThrustSignal *sg.Signal = sg.NewSignal(0.0, 0.0005, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    TorqueSignal *sg.Signal = sg.NewSignal(0.0, 0.005, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    EscL1Signal *sg.Signal = sg.NewSignal(0.0, 0.00001, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    EscL2Signal *sg.Signal = sg.NewSignal(0.0, 0.00001, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    EscL3Signal *sg.Signal = sg.NewSignal(0.0, 0.00001, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    TorqueSignal *sg.Signal = sg.NewSignal(0.0, 0.005, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    TorqueSignal *sg.Signal = sg.NewSignal(0.0, 0.005, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    RpmSignal *sg.Signal = sg.NewSignal(0.0, 0.0005, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.1})
    BldcTempSignal *sg.Signal = sg.NewSignal(0.0, 1.0, sg.NoiseParams{Enabled: true, Offset: 0.0, Gain: 0.2})
)

