package routeros

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRouterOSVersion(ros string) (version uint64, err error) {
	for i, p := range strings.Split(strings.TrimSpace(ros), ".") {
		var u uint64
		if u, err = strconv.ParseUint(p, 10, 64); err != nil {
			err = fmt.Errorf("RouterOS version parts parsing error, %v", err)
			return
		} else {
			version += u << ((2 - i) * 8)
		}

		if i > 3 {
			break
		}
	}

	if version == 0 {
		err = fmt.Errorf("RouterOS version parsing error, version is zero")
	}
	return
}

func routerOSVersionAtLeast(minVersion string) (bool, error) {
	if RouterOSVersion == "" {
		return false, fmt.Errorf("RouterOS version is not set")
	}

	currentVersion, err := parseRouterOSVersion(RouterOSVersion)
	if err != nil {
		return false, err
	}

	minimumVersion, err := parseRouterOSVersion(minVersion)
	if err != nil {
		return false, err
	}

	return currentVersion >= minimumVersion, nil
}
