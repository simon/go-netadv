package netadv

import "net"

// IPNetOr make OR bitwise operation on net.IPNet
func IPNetOr(n1, n2 *net.IPNet) *net.IPNet {
	oip := SliceOr(n1.IP, n2.IP)
	omask := SliceOr(n1.Mask, n2.Mask)
	return &net.IPNet{IP: oip, Mask: omask}
}
