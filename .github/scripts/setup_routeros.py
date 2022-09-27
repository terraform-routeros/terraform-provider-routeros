import routeros_api
import os


def main():
    user = os.environ.get("ROS_USERNAME")
    pswd = os.environ.get("ROS_PASSWORD")
    ip_addr = os.environ.get("ROS_IP_ADDRESS")

    connection = routeros_api.RouterOsApiPool(
        ip_addr, username=user, password=pswd, port=8728, plaintext_login=True)
    api = connection.get_api()

    # Set up bridge

    bridge = api.get_resource("/interface/bridge")
    bridge.add(name="bridge")

    # Set up certificates to allow use of REST API

    certificate = api.get_resource("/certificate")
    certificate.add(name="root-cert", common_name="MyRouter",
                    days_valid="3650", key_usage="key-cert-sign,crl-sign")
    certificate.add(name="https-cert",
                    common_name="MyRouter", days_valid="3650")
    certs = certificate.get()
    root_cert_id = [x['id'] for x in certs if x['name'] == "root-cert"][0]
    http_cert_id = [x['id'] for x in certs if x['name'] == "https-cert"][0]
    api.get_binary_resource("/").call("certificate/sign",
                                      {"id": bytes(root_cert_id, "utf-8")})
    api.get_binary_resource("/").call("certificate/sign",
                                      {"id": bytes(http_cert_id, "utf-8"), "ca": b"root-cert"})
    services = api.get_resource("/ip/service")
    for x in services.get():
        if x['name'] in ['www-ssl', 'api-ssl']:
            services.set(id=x['id'],certificate="https-cert", disabled="false")

    # Create a DHCP pool

    pools = api.get_resource("/ip/pool")
    pools.add(name="dhcp", ranges="192.168.88.100-192.168.88.200")

    wireguard = api.get_resource("/interface/wireguard")
    wireguard.add(name="wg1")

    # Output the list of interfaces

    print(api.get_resource("/interface").get())


if __name__ == "__main__":
    main()
