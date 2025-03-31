package common

type PowertrainReadings struct {
	Armed           bool
	Thrust          float64 // Nm
	Torque          float64 // Nm
	Speed           uint32  // RPM
	MechanicalPower float64 // W
	// Power train
	Current float64 // A
	Voltage float64 // V
	Power   float64 // W
	// ESC output signals
	L1Voltage   float64 // V
	L2Voltage   float64 // V
	L3Voltage   float64 // V
	Efficiency  float64 // %
	Temperature float64 // C
}
