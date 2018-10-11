package ipfmt

import (
	"fmt"
	"net"
)

func ToHex(ip net.IP) string {
	return fmt.Sprintf("%#x.%#x.%#x.%#x", ip[12], ip[13], ip[14], ip[15])
}

func ToOctal(ip net.IP) string {
	return fmt.Sprintf("%#o.%#o.%#o.%#o", ip[12], ip[13], ip[14], ip[15])
}

func ToInt(ip net.IP) string {
	i := int(ip[12]) * 16777216
	i += int(ip[13]) * 65536
	i += int(ip[14]) * 256
	i += int(ip[15])
	return fmt.Sprintf("%d", i)
}

// no dotted hex
func ToSingleHex(ip net.IP) string {
	i := int(ip[12]) * 16777216
	i += int(ip[13]) * 65536
	i += int(ip[14]) * 256
	i += int(ip[15])
	return fmt.Sprintf("%#x", i)
}

func Combo(ip net.IP) string {
	return fmt.Sprintf("%#x.%d.%#o.%#x", ip[12], ip[13], ip[14], ip[15])
}
