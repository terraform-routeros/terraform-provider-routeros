package routeros

import (
	"fmt"
	"maps"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Simple renames belong in mikrotik_resource_drift.yaml. These settings also
// changed their value representation and need conversion at the API boundary.
func versionedWriteFields(item MikrotikItem, path string, d *schema.ResourceData) {
	version, err := parseRouterOSVersion(RouterOSVersion)
	if err != nil {
		return
	}
	switch path {
	case "/interface/bridge":
		if version < 7<<16|23<<8 {
			return
		}
		legacy, present := item["add-dhcp-option82"]
		delete(item, "add-dhcp-option82")
		if !present || !attributeConfigured(d, "add_dhcp_option82") {
			return
		}
		if legacy == "yes" {
			item["dhcp-agent-circuit-id"] = "$(INTERFACE):$(VID)"
			item["dhcp-agent-remote-id"] = "$(BRIDGEMAC)"
		} else {
			delete(item, "dhcp-agent-circuit-id")
			delete(item, "dhcp-agent-remote-id")
			item["!dhcp-agent-circuit-id"] = ""
			item["!dhcp-agent-remote-id"] = ""
		}
	case "/routing/bgp/connection", "/routing/bgp/template":
		// RouterOS 7.20 moved the connection router-id to the instance and
		// removed the deprecated input.ignore-as-path-len parameter.
		if version >= 7<<16|20<<8 {
			delete(item, "cluster-id")
			delete(item, "router-id")
			delete(item, "input.ignore-as-path-len")
		}
		if version < 7<<16|22<<8 {
			return
		}
		legacy, present := item["add-path-out"]
		delete(item, "add-path-out")
		if !present || !attributeConfigured(d, "add_path_out") {
			return
		}
		if legacy == "all" {
			item["output.add-path"] = "ip,ipv6"
		} else {
			item["output.add-path"] = ""
		}
	case "/interface/bridge/port":
		// RouterOS 7.24 only accepts the STP port priority in hexadecimal
		// form (0x10..0xF0) and rejects decimal values.
		if version < 7<<16|24<<8 {
			return
		}
		if priority, ok := item["priority"]; ok && priority != "" &&
			!strings.HasPrefix(priority, "0x") && !strings.HasPrefix(priority, "0X") {
			if value, err := strconv.ParseInt(priority, 10, 64); err == nil {
				item["priority"] = fmt.Sprintf("0x%X", value)
			}
		}
	case "/system/logging/action":
		// RouterOS 7.24 removed the deprecated bsd-syslog parameter.
		if version >= 7<<16|24<<8 {
			delete(item, "bsd-syslog")
		}
	}
}

func versionedReadFields(item MikrotikItem, path string, d *schema.ResourceData) MikrotikItem {
	version, err := parseRouterOSVersion(RouterOSVersion)
	if err != nil {
		return item
	}
	switch path {
	case "/interface/bridge":
		if version >= 7<<16|23<<8 && attributeConfigured(d, "add_dhcp_option82") {
			item = maps.Clone(item)
			item["add-dhcp-option82"] = BoolToMikrotikJSON(item["dhcp-agent-circuit-id"] != "" || item["dhcp-agent-remote-id"] != "")
		}
	case "/routing/bgp/connection", "/routing/bgp/template":
		if version >= 7<<16|22<<8 && attributeConfigured(d, "add_path_out") {
			item = maps.Clone(item)
			item["add-path-out"] = "none"
			if strings.Contains(item["output.add-path"], "ip") {
				item["add-path-out"] = "all"
			}
			delete(item, "output.add-path")
		}
	}
	return item
}

func attributeConfigured(d *schema.ResourceData, name string) bool {
	config := d.GetRawConfig()
	return config.IsKnown() && !config.IsNull() && config.Type().HasAttribute(name) &&
		!config.GetAttr(name).IsNull()
}
