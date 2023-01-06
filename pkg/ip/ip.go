package ip

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
)

func Port(protocol string, host string, port int) (int, bool) {
	address := host + strconv.Itoa(port)

	_, err := net.Dial("tcp", address)

	if err == nil {
		return port, true
	} else {
		return port, false
	}
}

func Ping(host string) {
	cmd := exec.Command("ping", host, "-c 10", "-w 5", "-i 0.1")
	stats := cmd.Run()
	fmt.Println(stats)
}
