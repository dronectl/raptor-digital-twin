package signals

import (
	"math"
	"math/rand"
)

type Signal interface {
    Process() float64
    Enable(enable bool)
}

type NoiseParams struct {
    Gain float64
    Offset float64
}

type BaseSignal struct {
    Noise NoiseParams
    out float64
    t float64
    UpdateFreq int // Hz
    enable bool
}

// superimpose noise on output signal
func (s *BaseSignal)applyNoise() {
    // range [-1, 1]
    gainSeed := rand.Float64() - rand.Float64()
    s.out += s.Noise.Offset + (s.Noise.Gain * gainSeed)
}

func (s *BaseSignal)Enable(enable bool) {
    s.t = 0.0
    s.out = 0.0
    s.enable = enable
}

// Sine Signal Generator
type SignalSin struct {
    BaseSignal

    Amplitude float64
    Offset float64
    Frequency float64 // rads
    Phase float64 // rads
}

// compute next sine signal valuu
func (s *SignalSin)compute() {
    if !s.enable {
        return
    }
    s.t += 1.0 / float64(s.UpdateFreq)
    s.out = s.Amplitude * math.Sin((2 * math.Pi * s.Frequency * s.t) + s.Phase) + s.Offset
}

func (s *SignalSin)Process() float64 {
    s.compute()
	s.applyNoise()
    return s.out
}

func NewSignalSin(amplitude, offset, freq, phase float64, updateFreq int) *SignalSin {
    return &SignalSin{
        BaseSignal: BaseSignal{
            Noise: NoiseParams{Offset: 0.0, Gain: 0.0},
            out: 0.0,
            t: 0.0,
            enable: false,
            UpdateFreq: updateFreq,
        },
        Amplitude: amplitude,
        Offset: offset,
        Frequency: freq,
        Phase: phase,
    }
}

// Exponential Signal Generator
type SignalExp struct {
    BaseSignal

    // ramp rate
    Tau float64
    Ref float64
}

// compute next exponential signal valuu
func (s *SignalExp)compute() {
    if !s.enable {
        return
    }
    s.t += 1.0 / float64(s.UpdateFreq)
    s.out = (s.Ref - s.out)*(1.0 - math.Exp(-s.t / s.Tau)) + s.out
}

func (s *SignalExp)Process() float64 {
    s.compute()
	s.applyNoise()
    return s.out
}


func NewSignalExp(tau float64, updateFreq int) *SignalExp {
    return &SignalExp{
        BaseSignal: BaseSignal{
            Noise: NoiseParams{Offset: 0.0, Gain: 0.0},
            out: 0.0,
            t: 0.0,
            enable: false,
            UpdateFreq: updateFreq,
        },
        Tau: tau,
        Ref: 0.0,
    }
}

