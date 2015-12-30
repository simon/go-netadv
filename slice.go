package netadv

// SliceShiftRight return a slive of bytes shifted by "bits" bits
func SliceShiftRight(s1 []byte, bits uint) []byte {
	l := len(s1)
	s := make([]byte, l)

	r := bits / 8
	rs := bits - r*8

	for i := 0; i < l; i++ {
		// relative index
		ri := i - int(r)
		bl := byte(0)
		if ri >= 1 {
			bl = s1[ri-1] << (8 - rs)
		}

		bu := byte(0)
		if ri >= 0 {
			bu = s1[ri] >> rs
		}
		s[i] = bl | bu
	}

	return s
}

// SliceSetOne set 1 into the b'th byte
func SliceSetOne(s1 []byte, b uint) []byte {
	r := b / 8
	rs := b - r*8

	mb := byte(0x80)
	mb = mb >> rs

	s1[r] = s1[r] | mb

	return s1
}
