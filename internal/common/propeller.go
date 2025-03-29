package common

// Propeller parameters Note: (1 inch = 0.0254 m)
type PropellerParameters struct {
    Diameter float64 // inches 
    Pitch float64 // inches

    // coefficient of thrust estimator tuning coefficients (CtGain * (Pitch / Diameter) + CtOffset)
    CtGain float64
    CtOffset float64
}

var DefaultPropellerParameters = PropellerParameters{
    Diameter : 10.0,
    Pitch : 4.5,
    CtGain : 0.01,
    CtOffset: 0.0001,
}
