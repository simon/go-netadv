package netadv

import "net"

// SplitNetworks split network in a
func SplitNetworks(network *net.IPNet, newsize uint) ([]*net.IPNet, error) {
	var ret []*net.IPNet
	oldsize, _ := network.Mask.Size()

	// use a stack to avoid a recurvive function
	ret = append(ret, network)
	for i := oldsize; i < int(newsize); i++ {
		var tmp []*net.IPNet

		// parse network array to split all of them in two
		for _, n := range ret {
			one, two := SplitNetworkInTwo(n)
			tmp = append(tmp, one)
			tmp = append(tmp, two)
		}
		ret = tmp
	}

	return ret, nil
}

// SplitNetworkInTwo split a network in two. return two sub network
func SplitNetworkInTwo(network *net.IPNet) (*net.IPNet, *net.IPNet) {
	size, _ := network.Mask.Size()
	newMask := net.CIDRMask(size+1, 8*net.IPv6len)

	ip2 := SliceSetOne(net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, uint(size))
	ip2 = SliceOr(network.IP, ip2)

	return &net.IPNet{IP: network.IP, Mask: newMask}, &net.IPNet{IP: ip2, Mask: newMask}
}
