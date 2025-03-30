package common

type IPCChannels struct {
    Control chan SimControl
    PowertrainReadings chan PowertrainReadings
    EnvironmentReadings chan EnvironmentReadings
}
