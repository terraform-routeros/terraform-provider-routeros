package routeros

import (
	"fmt"
	"strconv"
)

var bitUnitMap = map[byte]uint64{
	// RouterOS accepts either case for kilo (`max-limit=30k` stores 30000) but
	// rejects lowercase for the larger units: `30m` is an "invalid trailing
	// char" error while `30M` is accepted. This follows SI, where the kilo
	// prefix is lower case and the rest are upper case.
	'k': 1e3,
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
		return 0, fmt.Errorf(`bits: invalid value %q`, s)
	}

	// Cut 'bps'.
	if l := len(s); l > 4 && s[l-3:] == "bps" {
		s = s[:l-3]
	}

	// Verifying the value.
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			// The last symbol.
			if i == len(s)-1 {
				var ok bool
				unit, ok = bitUnitMap[s[i]]
				if !ok {
					return 0, fmt.Errorf(`bits: unknown unit %q in value %q`, string(s[i]), s)
				}
				s = s[:i]
			} else {
				return 0, fmt.Errorf(`bits: invalid value %q`, s)
			}
		}
	}

	d, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(`bits: invalid value %q`, s)
	}

	// We do not control overflow on multiplication!
	return d * unit, nil
}

var byteUnitMap = map[byte]uint64{
	// See the note on bitUnitMap: lowercase kilo only.
	'k': 1 << 10,
	'K': 1 << 10,
	'M': 1 << 20,
	'G': 1 << 30,
	'T': 1 << 40,
	'P': 1 << 50,
	'E': 1 << 60,
}

func ParseByteValues(s string) (uint64, error) {
	var unit uint64 = 1

	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, fmt.Errorf(`bytes: invalid value %q`, s)
	}

	// Cut 'Bps'.
	// if l := len(s); l > 4 && s[l-3:] == "Bps" {
	// 	s = s[:l-3]
	// }

	// Verifying the value.
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			// The last symbol.
			if i == len(s)-1 {
				var ok bool
				unit, ok = byteUnitMap[s[i]]
				if !ok {
					return 0, fmt.Errorf(`bytes: unknown unit %q in value %q`, string(s[i]), s)
				}
				s = s[:i]
			} else {
				return 0, fmt.Errorf(`bytes: invalid value %q`, s)
			}
		}
	}

	d, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(`bytes: invalid value %q`, s)
	}

	// We do not control overflow on multiplication!
	return d * unit, nil
}
