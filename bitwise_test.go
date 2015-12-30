package netadv

import (
	"bytes"
	"testing"
)

func TestSliceOr(t *testing.T) {
	zero := []byte{0, 0, 0, 0}
	one := []byte{0, 0, 1, 0}
	two := []byte{0, 0, 2, 0}
	three := []byte{0, 0, 3, 0}

	t1 := SliceOr(zero, zero)
	if !bytes.Equal(t1, zero) {
		t.Fatal("zero OR zero have to be zero")
	}

	t2 := SliceOr(zero, three)
	if !bytes.Equal(t2, three) {
		t.Fatal("zero OR three have to be three")
	}

	t3 := SliceOr(one, two)
	if !bytes.Equal(t3, three) {
		t.Fatal("one OR two have to be three")
	}
}
