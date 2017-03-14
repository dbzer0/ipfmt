package main

import (
	"./ipfmt"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	Version string
	Build   string
)

func helpMe(prog string) {
	help := fmt.Sprintf("Usage: \n  %s 127.0.0.1", prog)
	fmt.Println(help)
}

func main() {
	flag.Usage = func() {
		helpMe(os.Args[0])
		fmt.Println()
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	// преобразуем строку в IP
	IP := net.ParseIP(os.Args[1])
	if IP == nil {
		fmt.Printf("ERROR: invalid IP address: %s\n", os.Args[1])
		os.Exit(1)
	}

	ipf := ipfmt.NewIpFormat(IP)
	ipf.PrintAll()

	os.Exit(0)
}
