package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-routeros/routeros"
)

var (
	commands = []string{
		"/certificate/add name=ssl common-name=router key-size=prime256v1",
		"/certificate/sign .id=ssl",
		"/ip/service/set .id=www-ssl disabled=no certificate=ssl",
		"/ip/service/set .id=api-ssl disabled=no certificate=ssl",
		"/interface/bridge/add name=bridge",
		"/ip/pool/add name=dhcp ranges=192.168.88.100-192.168.88.200",
		"/interface/wireguard/add name=wg1",
		"/interface/list/add name=list",
	}
)

func main() {
	username := os.Getenv("ROS_USERNAME")
	password := os.Getenv("ROS_PASSWORD")
	host := os.Getenv("ROS_IP_ADDRESS")

	var err error
	var client *routeros.Client
	for i := 0; i < 12; i++ {
		log.Printf("Connection attempt #%v... ", i)
		client, err = routeros.Dial(host+":8728", username, password)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatal("Host is not available")
	}
	defer client.Close()

	for _, command := range commands {
		cmd := strings.Split(command, " ")

		for i, c := range cmd {
			if strings.ContainsRune(c, '=') && len(c) > 0 && c[0] != '=' {
				cmd[i] = "=" + c
			}
		}

		log.Println(cmd)

		res, err := client.RunArgs(cmd)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(res)
	}
}
