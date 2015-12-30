package netadv

// SliceAnd return an AND on two slice of bytes
func SliceAnd(s1, s2 []byte) []byte {
	mini := MinI(len(s1), len(s2))
	s := make([]byte, mini)
	for i := 0; i < mini; i++ {
		s[i] = s1[i] & s2[i]
	}
	return s
}

// SliceOr return an OR on two slice of bytes
func SliceOr(s1, s2 []byte) []byte {
	mini := MinI(len(s1), len(s2))
	s := make([]byte, mini)
	for i := 0; i < mini; i++ {
		s[i] = s1[i] | s2[i]
	}
	return s
}

// SliceXor return an XOR on two slice of bytes
func SliceXor(s1, s2 []byte) []byte {
	mini := MinI(len(s1), len(s2))
	s := make([]byte, mini)
	for i := 0; i < mini; i++ {
		s[i] = s1[i] ^ s2[i]
	}
	return s
}

// SliceNot return the NOT of bytes slice
func SliceNot(s1 []byte) []byte {
	mini := len(s1)
	s := make([]byte, mini)
	for i := 0; i < mini; i++ {
		s[i] = ^s1[i]
	}
	return s
}
