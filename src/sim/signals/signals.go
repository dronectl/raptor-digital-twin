package signals

import (
	"math"
	"math/rand"
)

type SignalProcessor interface {
    compute() float64
}

type NoiseParams struct {
    Enable bool
    Gain float64
    Offset float64
}

type Signal struct {
    np NoiseParams
    out float64
    t float64
    UpdateFreq int // Hz
    Enable bool
}

type SignalSin struct {
    signal Signal

    Amplitude float64
    Offset float64
    Frequency float64 // rads
    Phase float64 // rads
}

type SignalExp struct {
    signal Signal

    // ramp rate
    Tau float64
    Ref float64
}

// superimpose noise on output signal
func (s *Signal)noiseGen() {
    if s.np.Enable {
        // range [-1, 1]
        gainSeed := rand.Float64() - rand.Float64()
        s.out += s.np.Offset + (s.np.Gain * gainSeed)
    }
}

func (s *Signal)Reset() {
    s.t = 0.0
}

// compute next sine signal valuu
func (s *SignalSin)compute() float64 {
    s.signal.t += 1.0 / float64(s.signal.UpdateFreq)
    s.signal.out = s.Amplitude * math.Sin((s.Frequency * s.signal.t) + s.Phase) + s.Offset
    return s.signal.out
}

// compute next exponential signal valuu
func (s *SignalExp)compute() float64 {
    s.signal.t += 1.0 / float64(s.signal.UpdateFreq)
    s.signal.out = (s.Ref - s.signal.out)*(1.0 - math.Exp(-s.signal.t / s.Tau)) + s.signal.out
    return s.signal.out
}

// Process the signal with thread safety
func Process(sp SignalProcessor, signal *Signal) float64 {
    if signal.Enable {
        signal.out = sp.compute()
    } else {
        signal.out = 0.0
    }
	signal.noiseGen()
	return signal.out
}

