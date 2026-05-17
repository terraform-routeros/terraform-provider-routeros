package routeros

import (
	"fmt"
	"strings"
)

func routerOSVersionAtLeast(minVersion string) (bool, error) {
	if RouterOSVersion == "" {
		return false, fmt.Errorf("RouterOS version is not set")
	}

	currentVersion, err := parseRouterOSVersion(strings.TrimSpace(RouterOSVersion))
	if err != nil {
		return false, err
	}

	minimumVersion, err := parseRouterOSVersion(minVersion)
	if err != nil {
		return false, err
	}

	return currentVersion >= minimumVersion, nil
}
