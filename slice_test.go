package netadv

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSliceShiftRight(t *testing.T) {
	one := []byte{0, 0, 1, 0}
	last := []byte{0, 0, 0, 1}

	t1 := SliceShiftRight(one, 8)
	if !bytes.Equal(t1, last) {
		fmt.Printf("%#v\n", t1)
		t.Fatal("the 32th byte is not to 1")
	}
}

func TestSliceSetOne(t *testing.T) {
	zero := []byte{0, 0, 0, 0}
	one := []byte{0, 0, 1, 0}

	t1 := SliceSetOne(zero, 23)
	if !bytes.Equal(t1, one) {
		fmt.Printf("%#v\n", t1)
		t.Fatal("the 24th byte is not to 1")
	}
}
