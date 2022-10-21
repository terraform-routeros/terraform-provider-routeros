terraform {
  required_providers {
    routeros = {
      source = "GNewbury1/routeros"
    }
  }
}

provider "routeros" {
  hosturl  = "https://127.0.0.1"
  username = "admin"
  password = ""
  insecure = true
}