package routeros

import (
	"testing"
)

// qos_port is a per-port row auto-created by the switch chip; it can only be modified, not
// added/deleted, so it has no self-contained acceptance test (it would mutate an existing
// hardware port). Validated manually on real switches:
//   - Marvell-98DX8216 / CRS317-1G-16S+ / RouterOS 7.20.8: trust_l2/trust_l3 modified on a
//     down port (sfp-sfpplus8) and reverted; verified via REST; switch restored to baseline.
//   - Marvell-98DX226S / CRS310-8G+2S+ / RouterOS 7.16.2 and
//     Marvell-98DX8208 / CRS309-1G-8S+ / RouterOS 7.18.2: import + plan round-trip (no diff).
func TestAccInterfaceEthernetSwitchQosPort_basic(t *testing.T) {
	t.Log("Test skipped, the resource modifies an existing switch port and requires real hardware.")
}
