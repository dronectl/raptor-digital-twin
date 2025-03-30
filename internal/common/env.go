
package common

type EnvParameters struct {
    Rho float64 // air density kg/m^3
    Temperature float64 // Celsius
}

var DefaultEnvParameters = EnvParameters{
    Rho: 1.225,
    Temperature: 25.0,
}

type EnvironmentReadings struct {
	Temperature  float64 // C
    RelativeHumidity float64 // %
    Pressure float64 // hPa
}

