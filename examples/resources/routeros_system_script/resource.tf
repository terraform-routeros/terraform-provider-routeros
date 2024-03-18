resource "routeros_system_script" "script" {
  name   = "my_script"
  source = <<EOF
    :log info "This is a test script created by Terraform."
    EOF
  policy = ["read", "write", "test", "policy"]
}