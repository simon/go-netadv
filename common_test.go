package netadv

import "testing"

func TestMinI(t *testing.T) {
	one := 1
	two := 2

	if MinI(one, two) != one {
		t.Fatal("minimum between 1 and 2 is 1")
	}
}

func TestMaxI(t *testing.T) {
	one := 1
	two := 2

	if MaxI(one, two) != two {
		t.Fatal("maximum between 1 and 2 is 2")
	}
}
