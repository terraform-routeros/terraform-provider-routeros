package routeros

import (
	"testing"
)

// queue/interface rows (one per interface) are auto-created by RouterOS and can only be
// modified, not added/deleted, so there is no self-contained acceptance test.
// Validated manually on real switches:
//   - Marvell-98DX8216 / CRS317-1G-16S+ / RouterOS 7.20.8: queue modified on a down port
//     (sfp-sfpplus8) and reverted; verified via REST; switch restored to baseline.
//   - Marvell-98DX226S / CRS310-8G+2S+ / RouterOS 7.16.2 and
//     Marvell-98DX8208 / CRS309-1G-8S+ / RouterOS 7.18.2: import + plan round-trip (no diff).
func TestAccQueueInterface_basic(t *testing.T) {
	t.Log("Test skipped, the resource modifies an existing interface queue and requires real hardware.")
}
