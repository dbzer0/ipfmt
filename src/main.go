package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/dbzer0/ipfmt/src/ipfmt"
)

var (
	Version string
	Build   string
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: \n  %s 127.0.0.1\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	var ip net.IP
	if ip = net.ParseIP(os.Args[1]); ip == nil {
		fmt.Printf("ERROR: invalid IP address: %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("ip              : %s\n", ip)
	fmt.Printf("hex             : %s\n", ipfmt.ToHex(ip))
	fmt.Printf("octal           : %s\n", ipfmt.ToOctal(ip))
	fmt.Printf("integer         : %s\n", ipfmt.ToInt(ip))
	fmt.Printf("hex (no dotted) : %s\n", ipfmt.ToSingleHex(ip))
	fmt.Printf("combo (h.d.o.h) : %s\n", ipfmt.Combo(ip))
}
