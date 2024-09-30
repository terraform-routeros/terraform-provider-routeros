data "routeros_wifi_easy_connect" "test" {
  type     = "WPA2"
  ssid     = "test"
  password = "password12345"
}

output "qrcode" {
  value = data.routeros_wifi_easy_connect.test.qr_code
}

# We can disable the QR code output and view it in the state file if needed.
# terraform.exe state show data.routeros_wifi_easy_connect.test
