package main

import (
	"github.com/adzsx/gonets/pkg/ip"
	"fmt"
	"os"
)

func main() {
	up := ip.Ping("1.1.1.1")

	if up{
		fmt.Println("Host is up")
	}else{
		fmt.Println("Could not establish stable connection to host.")
		os.Exit(1)
	}
}
