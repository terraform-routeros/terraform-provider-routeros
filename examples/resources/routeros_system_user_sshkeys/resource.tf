resource "routeros_system_user_sshkeys" "test" {
  user    = "test-user-1"
  key     = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCyJ1EvW98veNVzR3VamNgmu0xOd/JK9YNvP/pa4WC5eT90UbX4TN7dKEK/x2FCwnnG9u0FQhzG2qa/Cg8meUvlfydn6uxc0/WCeXTKSu6sT63noPO6m4fHY7gu3Zt+fOc/WYGch9sBeWjZlCS1mA2lajkWhM3J8TFWCFm2Zk4/S3s5mt6VLbwpQnH2LhE41+azzDEVhcR6i3FfdgOF/J+j2fYYHJsBEKoQA5zUac2zWmz7X4Rv/g11ZBRqdMpHSD58o5F9lBb13antu5GcEs5RXpXp08OyXuRV9qhFpDBC8DOMALSOgT3vnu8uJLgo8QIulERofj/cRXbLCsmvMbpioBuGFXWx3ha4Ntd6z07kUh2KVbaIQLd/629UHNvgIhoBLlREJ8E5vllsX+jh8hRITHcCiEwXcDO+gG3hvJt0+jm8S8SObE/IHk8VuwWdhIsSku5vd+wVlxm8VeJzjc0cjdIiytvsq8VpLudKEUiqR0f2tHcoq8H+xcJv3Ycu1i8="
  comment = "Test User"
}
