package netadv

import "net"

// ConvertIPv4InIPv6Network convert ipv6 network followed by ipv4
func ConvertIPv4InIPv6Network(network, ip string) (*net.IPNet, error) {
	// parse original network
	_, pNet1, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}

	// create an empty network
	_, pNet2, err := net.ParseCIDR("0::/63")
	if err != nil {
		return nil, err
	}

	// set 2 last bytes in new network
	ip2 := net.ParseIP(ip)
	pNet2.IP[0] = ip2[14]
	pNet2.IP[1] = ip2[15]

	// we shift network to follow the first ip
	shift, _ := pNet1.Mask.Size()
	pNet2.IP = SliceShiftRight(pNet2.IP, uint(shift))

	// return the OR of joined network
	return IPNetOr(pNet1, pNet2), nil
}
