package main

import (
	"fmt"
	"net"
	"time"
)

// We'll check the status of the address (domain+port) within the duration of 5 second and return the status of the said address
func Check(domain string, port string) string {
	var status string
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		status = fmt.Sprintf("The site \"%v\" is unreachable \n Error: %v", domain, err)
	} else {
		status = fmt.Sprintf("The site \"%v\" is up and running, \n From: %v\n To: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
