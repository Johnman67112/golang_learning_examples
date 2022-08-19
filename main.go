package main

import (
	"strings"

	"github.com/Johnman67112/golang_learning/datagrouping"
	"github.com/Johnman67112/golang_learning/flowcontrol"
)

func main() {
	GeneralHub("data_ex")
}

func GeneralHub(option string) {
	switch {
	case strings.Contains(option, "flow"):
		flowcontrol.ControlFlow(option)
	case strings.Contains(option, "data"):
		datagrouping.DataGrouping(option)
	}
}
