package main

import (
	"flag"
	"fmt"
	"port/port"
)

func main() {
	// Command-line flags
	target := flag.String("target", "localhost", "Target IP address or domain")
	flag.Parse()

	fmt.Printf("Port Scanner Working for %s\n", *target)

	results := port.InitialScan(*target, 1, 65535) // Scan ports from 1 to 65535
	flag:=1
	for _, result := range results {
		if result.State == "Open" {
			fmt.Printf("Port %d Open \n", result.Port)
			flag=0
		}
	}
	if flag==1{
		fmt.Println("No Open Ports")
	}
}
