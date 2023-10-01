package routeros

import (
	"fmt"
	"strings"
)

var ROSVersion rosVersionType

type rosVersionType int

const (
	v7_10 rosVersionType = iota
	v7_9
	v7_8
)

var ROSSupportedVersions = []string{
	"7.10", "7.8", "7.9",
}

func (v rosVersionType) String() string {
	return ROSSupportedVersions[v]
}

func ParseROSVersion(ver string) rosVersionType {
	var res int = -1

	for i, v := range ROSSupportedVersions {
		if v == ver {
			res = i
			break
		}
	}

	if res == -1 {
		panic("Unsupported ROS version: '" + ver + "'")
	}

	return rosVersionType(res)
}

func GetROSSupportedVersions() string {
	return fmt.Sprintf("\nROS supported versions: %v\n\n", strings.Join(ROSSupportedVersions, ", "))
}
