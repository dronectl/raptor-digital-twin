package signals

import (
    "time"
    "math"
    "testing"
)

func TestSignalSetReference(t *testing.T) {
    exp := NewExp(0.0, 1.2, NoiseParams{Offset: 0.0, Gain: 0.0})
    const ref = 1.0
    exp.signal.SetReference(ref)
    if exp.signal.ref != ref {
        t.Errorf("Expected %f got %f", ref, exp.signal.ref)
    }
}

func TestSignalProcessExp(t *testing.T) {
    const (
        initial float64 = 0.0
        tau float64 = 0.0005 // fast signal response to make utests fast
        offset float64 = 0.0
        gain float64 = 0.0
    )
    exp := NewExp(initial, tau, NoiseParams{Offset: offset, Gain: gain})
    const ref = 10.0 
    exp.signal.SetReference(ref)
    for range 10 {
        tmp := 0.0
        v := Process(exp, &exp.signal)
        t.Logf("t: %v v: %f", exp.signal.t, v)
        if v != exp.signal.out {
            t.Errorf("Expected %f got %f", exp.signal.out, v)
            break
        }
        if v < tmp {
            t.Errorf("Expected positive increments")
            break
        }
    }
    // start from s.a and approach 0
    exp.signal.SetReference(0.0)
    for range 10 {
        tmp := exp.signal.out
        v := Process(exp, &exp.signal)
        t.Logf("t: %v v: %f", exp.signal.t, v)
        if v != exp.signal.out {
            t.Errorf("Expected %f got %f", exp.signal.out, v)
            break
        }
        if v > tmp {
            t.Errorf("Expected negative increments")
            break
        }
    }
}

func TestSignalProcessSin(t *testing.T) {
    const (
        initial float64 = 0.0
        amplitude float64 = 12.0
        offset float64 = 12.0
        phase float64 = math.Pi
        frequency float64 = 2 * math.Pi * 10
        noiseGain float64 = 0.0
        noiseOffset float64 = 0.0
    ) 
    sin := SignalSin{
        signal: Signal{
            t: time.Now(),
            np: NoiseParams{Enabled: true, Offset: noiseOffset, Gain: noiseGain},
            out: initial,
        },
        Amplitude: amplitude,
        Offset: offset,
        Frequency: frequency,
        Phase: phase,
    }
    const ref = 10.0 
    for range 10 {
        v := Process(sin, &sin.signal)
        t.Logf("t: %v v: %f", sin.signal.t, v)
        if v != sin.signal.out {
            t.Errorf("Expected %f got %f", sin.signal.out, v)
            break
        }
        if v < amplitude - offset || v > amplitude {
            t.Errorf("Expected positive increments")
            break
        }
    }
}

func TestSignalNewExp(t *testing.T) {
    const (
        initial float64 = 34.0
        tau float64 = 0.02
        offset float64 = 7.0
        gain float64 = 0.2
    ) 
    exp := NewExp(initial, tau, NoiseParams{Offset: offset, Gain: gain})
    if exp.signal.out != initial {
        t.Errorf("Initial - expected %f got %f", initial, exp.signal.out)
    }
    if exp.tau != tau {
        t.Errorf("Tau - expected %f got %f", tau, exp.tau)
    }
    if exp.signal.np.Offset != offset {
        t.Errorf("Noise Offset - expected %f got %f", offset, exp.signal.np.Offset)
    }
    if exp.signal.np.Gain != gain {
        t.Errorf("Noise Gain - expected %f got %f", gain, exp.signal.np.Gain)
    }
}

func TestSignalNewSin(t *testing.T) {
    const (
        initial float64 = 34.0
        amplitude float64 = 0.02
        offset float64 = 7.0
        phase float64 = math.Pi
        frequency float64 = 2 * math.Pi 
        noiseGain float64 = 7.0
        noiseOffset float64 = 0.2
    ) 
    sin := NewSin(initial, amplitude, frequency, phase, offset, NoiseParams{Offset: noiseOffset, Gain: noiseGain})
    if sin.signal.out != initial {
        t.Errorf("Initial - expected %f got %f", initial, sin.signal.out)
    }
    if sin.amplitude != amplitude {
        t.Errorf("Amplitude - expected %f got %f", amplitude, sin.amplitude)
    }
    if sin.phase != phase {
        t.Errorf("Phase - expected %f got %f", phase, sin.phase)
    }
    if sin.offset != offset {
        t.Errorf("Phase - expected %f got %f", offset, sin.offset)
    }
    if sin.signal.np.Offset != noiseOffset {
        t.Errorf("Noise Offset - expected %f got %f", offset, sin.signal.np.Offset)
    }
    if sin.signal.np.Gain != noiseGain {
        t.Errorf("Noise Gain - expected %f got %f", noiseGain, sin.signal.np.Gain)
    }
}
