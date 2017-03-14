package ipfmt

import (
	"fmt"
	"net"
)

type IPFormat struct {
	net.IP
}

// конструктор
func NewIpFormat(IP net.IP) *IPFormat {
	return &IPFormat{IP}
}

// hex
func (ipFormat IPFormat) ToHex() string {
	return fmt.Sprintf("%#x.%#x.%#x.%#x",
		ipFormat.IP[12],
		ipFormat.IP[13],
		ipFormat.IP[14],
		ipFormat.IP[15])
}

// octal
func (ipFormat IPFormat) ToOctal() string {
	return fmt.Sprintf("%#o.%#o.%#o.%#o",
		ipFormat.IP[12],
		ipFormat.IP[13],
		ipFormat.IP[14],
		ipFormat.IP[15])
}

// integer
func (ipFormat IPFormat) ToInt() string {
	i := int(ipFormat.IP[12]) * 16777216
	i += int(ipFormat.IP[13]) * 65536
	i += int(ipFormat.IP[14]) * 256
	i += int(ipFormat.IP[15])

	return fmt.Sprintf("%d", i)
}

// no dotted hex
func (ipFormat IPFormat) ToSingleHex() string {
	i := int(ipFormat.IP[12]) * 16777216
	i += int(ipFormat.IP[13]) * 65536
	i += int(ipFormat.IP[14]) * 256
	i += int(ipFormat.IP[15])

	return fmt.Sprintf("%#x", i)
}

func (ipFormat IPFormat) Combo() string {
	return fmt.Sprintf("%#x.%d.%#o.%#x",
		ipFormat.IP[12],
		ipFormat.IP[13],
		ipFormat.IP[14],
		ipFormat.IP[15])
}

func (ipFormat IPFormat) PrintAll() {

	fmt.Printf("ip              : %s\n", ipFormat.IP)
	fmt.Printf("hex             : %s\n", ipFormat.ToHex())
	fmt.Printf("octal           : %s\n", ipFormat.ToOctal())
	fmt.Printf("integer         : %s\n", ipFormat.ToInt())
	fmt.Printf("hex (no dotted) : %s\n", ipFormat.ToSingleHex())
	fmt.Printf("combo (h.d.o.h) : %s\n", ipFormat.Combo())

}
