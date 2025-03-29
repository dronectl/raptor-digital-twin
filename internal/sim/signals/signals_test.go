package signals

import (
    "math"
    "testing"
)

func TestSignalProcessExp(t *testing.T) {
    const (
        initial float64 = 0.0
        tau float64 = 0.05 // fast signal response to make utests fast
        offset float64 = 23.4
        noiseGain float64 = 1.1
        noiseOffset float64 = 0.001
    )
    exp := NewSignalExp(tau, offset, NoiseParams{gain: noiseGain, offset: noiseOffset}, 1000)
    for _, ref := range []float64{offset + 10.0, offset} {
        t.Logf("Signal Reference Step %f -> %f", exp.Ref, ref)
        exp.Ref = ref
        exp.Enable(true)
        for range 10 {
            // run process
            v := exp.Process()
            // calculate expected signal value
            base := (exp.Ref - exp.out)*(1.0 - math.Exp(-exp.t/ exp.Tau)) + exp.out
            noiseLower := -exp.noise.gain + exp.noise.offset + base
            noiseUpper := exp.noise.gain + exp.noise.offset + base
            if v != exp.out {
                t.Errorf("Expected %f got %f", exp.out, v)
                break
            }
            t.Logf("signal: %f expected: %f, range: [%f, %f]", v, base, noiseLower, noiseUpper)
            if v > noiseUpper || v < noiseLower {
                t.Errorf("Signal out of tolerance: got %f expected within range [%f %f]", v, noiseLower, noiseUpper)
                break
            }
        }
    }
}

func TestSignalProcessSin(t *testing.T) {
    const (
        initial float64 = 0.0
        amplitude float64 = 12.0
        offset float64 = 12.0
        phase float64 = 0.0 
        frequency float64 = 10.0
        noiseGain float64 = 1.0
        noiseOffset float64 = 0.1
    ) 
    sin := NewSignalSin(amplitude, offset, frequency, phase, NoiseParams{gain: noiseGain, offset: noiseOffset}, 1000)
    sin.Enable(true)
    for range 10 {
        // run process
        v := sin.Process()
        base := sin.Amplitude * math.Sin((math.Pi * 2 * sin.Frequency * sin.t) + sin.Phase) + sin.offset
        noiseLower := -sin.noise.gain + sin.noise.offset + base
        noiseUpper := sin.noise.gain + sin.noise.offset + base
        if v != sin.out {
            t.Errorf("Expected %f got %f", sin.out, v)
            break
        }
        t.Logf("signal: %f expected: %f, range: [%f, %f]", v, base, noiseLower, noiseUpper)
        if v > noiseUpper || v < noiseLower {
            t.Errorf("Signal out of tolerance: got %f expected within range [%f %f]", v, noiseLower, noiseUpper)
            break
        }
    }
}
