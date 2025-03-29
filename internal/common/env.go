
package common

type EnvParameters struct {
    Rho float64 // air density kg/m^3
    Temperature float64 // Celsius
}

var DefaultEnvParameters = EnvParameters{
    Rho: 1.225,
    Temperature: 25.0,
}
