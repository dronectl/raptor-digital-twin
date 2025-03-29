package common

type SimControl int

const (
    SIM_CMD_NULL SimControl = iota
    SIM_CMD_START
    SIM_CMD_STOP
)

