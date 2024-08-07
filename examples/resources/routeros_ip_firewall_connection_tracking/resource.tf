
resource "routeros_ip_firewall_connection_tracking" "data" {
  enabled                  = "yes"
  generic_timeout          = "3m"
  icmp_timeout             = "3m"
  loose_tcp_tracking       = "false"
  tcp_close_timeout        = "3m"
  tcp_close_wait_timeout   = "3m"
  tcp_established_timeout  = "3m"
  tcp_fin_wait_timeout     = "3m"
  tcp_last_ack_timeout     = "3m"
  tcp_max_retrans_timeout  = "3m"
  tcp_syn_received_timeout = "3m"
  tcp_syn_sent_timeout     = "3m"
  tcp_time_wait_timeout    = "3m"
  tcp_unacked_timeout      = "3m"
  udp_stream_timeout       = "3m"
  udp_timeout              = "3m"
}