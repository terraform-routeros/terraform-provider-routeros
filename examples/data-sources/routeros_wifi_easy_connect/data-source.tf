data "routeros_wifi_easy_connect" "test" {
  type     = "WPA2"
  ssid     = "test"
  password = "password12345"
}

output "qrcode" {
  value = data.routeros_wifi_easy_connect.test.qr_code
}
