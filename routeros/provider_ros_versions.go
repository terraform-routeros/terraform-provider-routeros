package routeros

import (
	"fmt"
	"strings"
)

var ROSVersion rosVersionType = v7_10

type rosVersionType int

const (
	v7_8 rosVersionType = iota
	v7_9
	v7_10
)

var ROSSupportedVersions = []string{
	"7.8", "7.9", "7.10",
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
