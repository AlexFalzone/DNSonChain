package util

import (
	"log"
	"net"
)

//take the ip from the current machine
func GetIP() []net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var ips []net.IP

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips
}
