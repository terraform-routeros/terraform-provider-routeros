package routeros

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// https://gist.github.com/vaerh/65b30353239cd17ee14f2e68983fbbf0

var pow = []int{24, 16, 8, 0}

func ipToULong(ip string) (res uint32) {
	var oct uint64

	for i, v := range strings.SplitN(ip, ".", 4) {
		oct, _ = strconv.ParseUint(v, 10, 32)
		res |= uint32(oct << pow[i])
	}
	return
}

func IpRangeToCIDR(ip1, ip2 string) (string, error) {
	a1 := ipToULong(ip1)
	a2 := ipToULong(ip2)

	if a1 > a2 {
		return "", errors.New("start must be less than end")
	}

	var m uint32 = 0xFFFFFFFF
	l := 32

	for m > 0 {
		m1 := m << 1
		if (a1&m1) != a1 || (a1|^m1) > a2 {
			break
		}
		m = m1
		l--
	}

	var addr = make([]byte, 4)
	binary.BigEndian.PutUint32(addr, a1)

	a1 |= ^m

	if a2 < a1+1 {
		return fmt.Sprintf("%d.%d.%d.%d/%v", addr[0], addr[1], addr[2], addr[3], l), nil
	}

	return ip1 + "-" + ip2, nil
}
