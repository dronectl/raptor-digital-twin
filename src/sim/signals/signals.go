
package signals

import (
    "sync"
    "time"
    "math"
    "math/rand"
)

type NoiseParams struct {
    offset float64
    gain float64
}

// Signals update as you observe them
type Signal struct {
    m sync.Mutex
    // ramp rate
    tau float64
    // temporal memory
    t time.Time
    // current value (where is the signal now)
    a float64
    // reference value (where the signal is going)
    b float64
    // noise params
    np NoiseParams
}

func (s *Signal)computeExp() float64 {
    t := time.Now()
    tdiff := float64(t.Sub(s.t).Nanoseconds()) / 1e9
    s.t = t
    s.a = (s.a - s.b)*(1.0 - math.Exp(-tdiff / s.tau)) + s.b
    return s.a
}

func (s *Signal)noiseGen() {
    s.a += s.np.offset + s.np.gain * rand.Float64()
}

func (s *Signal) SetReference(b float64) {
    s.m.Lock()
    s.b = b
    s.m.Unlock()
}

func (s *Signal) Process() float64 {
    s.m.Lock()
    defer s.m.Unlock()
    s.computeExp()
    s.noiseGen()
    return s.a
}

func NewSignal(freq int, initial, tau float64, np NoiseParams) *Signal {
    return &Signal{a: initial, tau: tau, np: np}
}


