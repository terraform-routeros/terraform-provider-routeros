resource "routeros_snmp_community" "test" {
	authentication_password = "authpasswd"
	authentication_protocol = "MD5"
	comment                 = "Comment"
	disabled                = true
	encryption_password     = "encpassword"
	encryption_protocol     = "DES"
	name                    = "private"
	read_access             = true
	security                = "private"
	write_access            = true
}
