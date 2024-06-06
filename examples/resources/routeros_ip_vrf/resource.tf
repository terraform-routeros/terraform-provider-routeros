resource "routeros_interface_veth" "veth1" {
	name    = "veth1"
}

resource "routeros_interface_veth" "veth2" {
	name    = "veth2"
}

resource "routeros_ip_vrf" "test_vrf_a" {
	disabled 	  = true
	name 		    = "vrf_1"
  comment     = "Custom routing"
	interfaces 	= ["veth1", "veth2"]
	depends_on  = [routeros_interface_veth.veth1, routeros_interface_veth.veth2]
}