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
    noise NoiseParams
    out float64
    t float64
    offset float64
    enable bool
    UpdateFreq int // Hz
}

// superimpose noise on output signal
func (s *BaseSignal)applyNoise() {
    // range [-1, 1]
    gainSeed := rand.Float64() - rand.Float64()
    s.out += s.noise.Offset + (s.noise.Gain * gainSeed)
}

func (s *BaseSignal)Enable(enable bool) {
    s.t = 0.0
    s.out = s.offset
    s.enable = enable
}

// Sine Signal Generator
type SignalSin struct {
    BaseSignal

    Amplitude float64
    Frequency float64 // rads
    Phase float64 // rads
}

// compute next sine signal valuu
func (s *SignalSin)compute() {
    if !s.enable {
        return
    }
    s.t += 1.0 / float64(s.UpdateFreq)
    s.out = s.Amplitude * math.Sin((2 * math.Pi * s.Frequency * s.t) + s.Phase) + s.offset
}

func (s *SignalSin)Process() float64 {
    s.compute()
	s.applyNoise()
    return s.out
}

func NewSignalSin(amplitude, offset, freq, phase float64, noiseParams NoiseParams, updateFreq int) *SignalSin {
    return &SignalSin{
        BaseSignal: BaseSignal{
            noise: noiseParams,
            out: 0.0,
            t: 0.0,
            enable: false,
            offset: offset,
            UpdateFreq: updateFreq,
        },
        Amplitude: amplitude,
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

func NewSignalExp(tau, offset float64, np NoiseParams, updateFreq int) *SignalExp {
    return &SignalExp{
        BaseSignal: BaseSignal{
            noise: np,
            out: 0.0,
            t: 0.0,
            enable: false,
            offset: offset,
            UpdateFreq: updateFreq,
        },
        Tau: tau,
        Ref: offset,
    }
}

