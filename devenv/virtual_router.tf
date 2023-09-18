
locals {
  iso_dir = "${path.root}/iso"
}


resource "libvirt_network" "chr_network" {
  # the name used by libvirt
  name = "chr_network"

  # mode can be: "nat" (default), "none", "route", "open", "bridge"
  mode = "nat"

  #  the domain used by the DNS server in this network
  domain = "routerosdev.local"

  #  list of subnets the addresses allowed for domains connected
  # also derived to define the host addresses
  # also derived to define the addresses served by the DHCP server
  addresses = ["10.10.10.0/24", "2001:db8:ca2:2::/64"]

  # (optional) the bridge device defines the name of a bridge device
  # which will be used to construct the virtual network.
  # (only necessary in "bridge" mode)
  # bridge = "br7"

  # (optional) the MTU for the network. If not supplied, the underlying device's
  # default is used (usually 1500)
  # mtu = 9000

  # (Optional) DNS configuration
  dns {
    # (Optional, default false)
    # Set to true, if no other option is specified and you still want to
    # enable dns.
    enabled = true
    # (Optional, default false)
    # true: DNS requests under this domain will only be resolved by the
    # virtual network's own DNS server
    # false: Unresolved requests will be forwarded to the host's
    # upstream DNS server if the virtual network's DNS server does not
    # have an answer.
    local_only = true

    # (Optional) one or more DNS forwarder entries.  One or both of
    # "address" and "domain" must be specified.  The format is:
    # forwarders {
    #     address = "my address"
    #     domain = "my domain"
    #  }
    #

    # (Optional) one or more DNS host entries.  Both of
    # "ip" and "hostname" must be specified.  The format is:
    # hosts  {
    #     hostname = "my_hostname"
    #     ip = "my.ip.address.1"
    #   }
    # hosts {
    #     hostname = "my_hostname"
    #     ip = "my.ip.address.2"
    #   }
    #

    # (Optional) one or more static routes.
    # "cidr" and "gateway" must be specified. The format is:
    # routes {
    #     cidr = "10.17.0.0/16"
    #     gateway = "10.18.0.2"
    #   }
  }

  # (Optional) Dnsmasq options configuration
  dnsmasq_options {
    # (Optional) one or more option entries.  Both of
    # "option_name" and "option_value" must be specified.  The format is:
    # options  {
    #     option_name = "server"
    #     option_value = "/base.domain/my.ip.address.1"
    #   }
    # options {
    #     option_name = "address"
    #     ip = "/.api.base.domain/my.ip.address.2"
    #   }
    #
  }
}

resource "libvirt_pool" "chr_image_pool" {
  name = "chr_image_pool"
  type = "dir"
  path = local.iso_dir



  provisioner "local-exec" {
    interpreter = ["bash", "-c"]
    command = "mkdir ${local.iso_dir} && pushd ${local.iso_dir} && wget https://download.mikrotik.com/routeros/${var.router_os_image_version}/chr-${var.router_os_image_version}.img.zip && unzip chr-${var.router_os_image_version}.img.zip"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "rm -rf ${self.path}"
  }
}

resource "libvirt_volume" "router_os_image" {
  name   = "router_os_chr"
  source = "${local.iso_dir}/chr-7.11.2.img"
  pool = libvirt_pool.chr_image_pool.name
}

resource "libvirt_domain" "cloud_hosted_router" {
  name   = "cloud_hosted_router"
  memory = "1024"
  vcpu   = 2
  cmdline = null


  disk {
    volume_id = libvirt_volume.router_os_image.id
  }

  # management interface for api testing
  network_interface {
    network_id     = libvirt_network.chr_network.id
    hostname       = "dev_chr"
    addresses      = ["10.10.10.10"]
    mac            = "AA:AA:AA:AA:AA:10"
    wait_for_lease = true
  }

  # interfaces 1->4 to be subject of tests
  network_interface {
    network_id     = libvirt_network.chr_network.id
    mac            = "AA:AA:AA:AA:AA:11"
  }

  network_interface {
    network_id     = libvirt_network.chr_network.id
    mac            = "AA:AA:AA:AA:AA:12"
  }

  network_interface {
    network_id     = libvirt_network.chr_network.id
    mac            = "AA:AA:AA:AA:AA:13"
  }

  network_interface {
    network_id     = libvirt_network.chr_network.id
    mac            = "AA:AA:AA:AA:AA:14"
  }
}