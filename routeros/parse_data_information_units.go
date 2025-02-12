package routeros

import (
	"fmt"
	"strconv"
)

var bitUnitMap = map[byte]uint64{
	'K': 1e3,
	'M': 1e6,
	'G': 1e9,
	'T': 1e12,
	'P': 1e15,
	'E': 1e18,
}

func ParseBitValues(s string) (uint64, error) {
	var unit uint64 = 1

	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, fmt.Errorf(`bits: invalid value "%v"`, s)
	}

	// Verifying the value.
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			// The last symbol.
			if i == len(s)-1 {
				var ok bool
				unit, ok = bitUnitMap[s[i]]
				if !ok {
					return 0, fmt.Errorf(`bits: unknown unit "%v" in value "%v"`, s[i], s)
				}
				s = s[:i]
			} else {
				return 0, fmt.Errorf(`bits: invalid value "%v"`, s)
			}
		}
	}

	d, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(`bits: invalid value "%v"`, s)
	}

	// We do not control overflow on multiplication!
	return d * unit, nil
}
