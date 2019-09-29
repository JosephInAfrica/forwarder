package read

import (
	"fmt"
	"io/util"
	"json"
)

type Status struct {
	Data    []Module `json:"data"`
	ErrCode int      `json:"err_code"`
}

type Module struct {
	UTop    Sensor `json:"u_top"`
	UMid    Sensor `json:"u_mid"`
	UBot    Sensor `json:"u_bot"`
	ErrCode int    `json:"err_code"`
}

type Sensor struct {
	T       float64 `json:"t"`
	H       float64 `json:"h"`
	ErrCode int     `json:"err_code"`
}

type BadSensor struct {
	// not sure if this is needed
	ErrCode int `json:"err_code"`
}
