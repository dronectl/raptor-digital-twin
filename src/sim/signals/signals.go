package signals

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type SignalProcessor interface {
    compute() float64
}

type NoiseParams struct {
    Enabled bool
    Gain float64
    Offset float64
}

type Signal struct {
    m sync.Mutex
    t time.Time
    np NoiseParams
    out float64
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
    if s.np.Enabled {
        s.out += s.np.Offset + s.np.Gain * rand.Float64()
    }
}

// compute next sine signal valuu
func (s *SignalSin)compute() float64 {
    tdiff := time.Since(s.signal.t).Seconds() 
    s.signal.out = s.Amplitude * math.Sin(s.Frequency * tdiff + s.Phase) + s.Offset
    s.signal.t = time.Now()
    return s.signal.out
}

// compute next exponential signal valuu
func (s *SignalExp)compute() float64 {
    tdiff := time.Since(s.signal.t).Seconds() 
    s.signal.out = (s.Ref - s.signal.out)*(1.0 - math.Exp(-tdiff / s.Tau)) + s.signal.out
    s.signal.t = time.Now()
    return s.signal.out
}

// Process the signal with thread safety
func Process(sp SignalProcessor, signal *Signal) float64 {
	signal.m.Lock()
	defer signal.m.Unlock()
	signal.out = sp.compute()
	signal.noiseGen()
	return signal.out
}

