package common

import (
    "math"
)

type BldcParameters struct {
    MaxSpeed uint32 // RPM
    MaxVoltage float64 // V
    KV float64 // rpm / V
    Resistance float64 // Ohm
    kt float64 // Nm/A
    ke float64 // V/(rad/s)
}

func CalculateKt(kv float64) float64 {
    return 60.0 / (2 * math.Pi * kv)
}

func CalculateKe(kv float64) float64 {
    return (1.0 / kv) * (60.0 / (2 * math.Pi))
}

var DefaultBldcParameters = BldcParameters{
    MaxSpeed:   10000,
    MaxVoltage: 12.0,
    KV:         1000.0,
    Resistance: 0.1,
    kt:         CalculateKt(1000.0),
    ke:         CalculateKe(1000.0),
}

