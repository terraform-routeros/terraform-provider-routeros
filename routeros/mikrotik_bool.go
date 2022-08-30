package routeros

func boolToMikrotikJSON(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func boolFromMikrotikJSON(s string) bool {
	if s == "true" || s == "yes" {
		return true
	}
	return false
}
