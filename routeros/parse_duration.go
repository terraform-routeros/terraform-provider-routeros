package routeros

import (
	"errors"
	"fmt"
	"time"
)

// https://github.com/xhit/go-str2duration
var unitMap = map[string]int64{
	"ms": int64(time.Millisecond),
	"s":  int64(time.Second),
	"m":  int64(time.Minute),
	"h":  int64(time.Hour),
	"d":  int64(time.Hour) * 24,
	"w":  int64(time.Hour) * 168,
}

func ParseDuration(s string) (time.Duration, error) {
	// ([0-9]*([0-9]*)?[a-z]+)+
	orig := s
	var d int64

	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, fmt.Errorf(`time: invalid duration "%v"`, orig)
	}
	for s != "" {
		var (
			v, f  int64       // integers before, after decimal point
			scale float64 = 1 // value = v + f/scale
		)

		var err error

		// The next character must be [0-9]
		if !('0' <= s[0] && s[0] <= '9') {
			return 0, fmt.Errorf(`time: invalid duration "%v"`, orig)
		}
		// Consume [0-9]*
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, fmt.Errorf(`time: invalid duration "%v"`, orig)
		}

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if '0' <= c && c <= '9' {
				break
			}
		}
		var unit int64
		if i != 0 {
			u := s[:i]
			s = s[i:]

			var ok bool
			unit, ok = unitMap[u]
			if !ok {
				return 0, fmt.Errorf(`time: unknown unit "%v" in duration "%v"`, u, orig)
			}
		} else {
			// missing unit in duration
			unit = int64(time.Second)
		}
		if v > (1<<63-1)/unit {
			// overflow
			return 0, fmt.Errorf(`time: invalid duration "%v"`, orig)
		}
		v *= unit
		if f > 0 {
			// float64 is needed to be nanosecond accurate for fractions of hours.
			// v >= 0 && (f*unit/scale) <= 3.6e+12 (ns/h, h is the largest unit)
			v += int64(float64(f) * (float64(unit) / scale))
			if v < 0 {
				// overflow
				return 0, fmt.Errorf(`time: invalid duration "%v"`, orig)
			}
		}
		d += v
		if d < 0 {
			// overflow
			return 0, fmt.Errorf(`time: invalid duration "%v"`, orig)
		}
	}

	return time.Duration(d), nil
}

var errLeadingInt = errors.New("time: bad [0-9]*") // never printed

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x int64, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > (1<<63-1)/10 {
			// overflow
			return 0, "", errLeadingInt
		}
		x = x*10 + int64(c) - '0'
		if x < 0 {
			// overflow
			return 0, "", errLeadingInt
		}
	}
	return x, s[i:], nil
}
