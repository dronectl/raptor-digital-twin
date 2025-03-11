package signals

import (
    "math"
    "testing"
)

func TestSignalProcessExp(t *testing.T) {
    const (
        initial float64 = 0.0
        tau float64 = 0.05 // fast signal response to make utests fast
        offset float64 = 0.001
        gain float64 = 1.1
    )
    exp := &SignalExp{
        signal: Signal{
            np: NoiseParams{
                Enable: true,
                Offset: offset,
                Gain: gain,
            },
            out: initial,
            Enable: true,
            UpdateFreq: 1000,
        },
        Tau: tau,
    }
    for _, ref := range []float64{10.0, 0.0} {
        t.Logf("Signal Reference Step %f -> %f", exp.Ref, ref)
        exp.Ref = ref
        exp.signal.Reset()
        for range 10 {
            // run process
            v := Process(exp, &exp.signal)
            // calculate expected signal value
            base := (exp.Ref - exp.signal.out)*(1.0 - math.Exp(-exp.signal.t/ exp.Tau)) + exp.signal.out
            noiseLower := -exp.signal.np.Gain + exp.signal.np.Offset + base
            noiseUpper := exp.signal.np.Gain + exp.signal.np.Offset + base
            if v != exp.signal.out {
                t.Errorf("Expected %f got %f", exp.signal.out, v)
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
        frequency float64 = 1000.0
        noiseGain float64 = 1.0
        noiseOffset float64 = 0.1
    ) 
    sin := &SignalSin{
        signal: Signal{
            np: NoiseParams{Enable: true, Offset: noiseOffset, Gain: noiseGain},
            out: initial,
            UpdateFreq: 10000,
            Enable: true,
        },
        Amplitude: amplitude,
        Offset: offset,
        Frequency: frequency,
        Phase: phase,
    }
    for range 10 {
        // run process
        v := Process(sin, &sin.signal)
        base := sin.Amplitude * math.Sin((math.Pi * 2 * sin.Frequency * sin.signal.t) + sin.Phase) + sin.Offset
        noiseLower := -sin.signal.np.Gain + sin.signal.np.Offset + base
        noiseUpper := sin.signal.np.Gain + sin.signal.np.Offset + base
        if v != sin.signal.out {
            t.Errorf("Expected %f got %f", sin.signal.out, v)
            break
        }
        t.Logf("signal: %f expected: %f, range: [%f, %f]", v, base, noiseLower, noiseUpper)
        if v > noiseUpper || v < noiseLower {
            t.Errorf("Signal out of tolerance: got %f expected within range [%f %f]", v, noiseLower, noiseUpper)
            break
        }
    }
}
