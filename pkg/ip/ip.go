package ip

import (
	"net"
	"os/exec"
	"strconv"
	"strings"
)

var(
	packets int
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

func Ping(host string) bool {
	cmd := exec.Command("ping", host, "-c 10", "-w 5", "-i 0.01")
	output, _ := cmd.Output()
	stats := string(output)

	list := strings.Split(stats, " ")
	for index, element := range list{
	
		if element == "received,"{
			packets, _ = strconv.Atoi(list[index-1])
		}
	}	

	return packets == 10
}
