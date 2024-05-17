
resource "routeros_ipv6_neighbor_discovery" "test" {
  interface 						= "ether1"
  hop_limit 						= 33
  advertise_dns 					= false
  advertise_mac_address 		    = true
  disabled 						= false
  managed_address_configuration	= true
  mtu 							= 9000
  other_configuration				= true
  pref64_prefixes					= []
  ra_delay						= "3s"
  ra_interval						= "3m20s-10m"
  ra_lifetime						= "30m"
  ra_preference					= "high"
  reachable_time					= "10m"
  retransmit_interval				= "12m"
}