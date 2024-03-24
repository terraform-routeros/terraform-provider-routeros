data "routeros_ip_services" "router" {
  provider = routeros.router
}

resource "routeros_ip_service" "router-disabled" {
  provider = routeros.router
  for_each = { for s in data.routeros_ip_services.router.services : s.name => s if s.name != "www-ssl" }

  disabled = true
  numbers  = each.value.name
  port     = each.value.port
}