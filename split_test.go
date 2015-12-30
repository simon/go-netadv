package netadv

import (
	"math/rand"
	"net"
	"testing"
)

func BenchmarkSetOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		size := rand.Intn(128)
		SliceSetOne(net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, uint(size))
	}
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		size := rand.Intn(128)
		byte1 := net.IP{0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		SliceShiftRight(byte1, uint(size))
	}
}
