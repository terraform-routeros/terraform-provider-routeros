resource "routeros_disk_settings" "test" {
  auto_smb_sharing     = false
  auto_smb_user        = "guest"
  auto_media_sharing   = false
  auto_media_interface = "lo"
}
