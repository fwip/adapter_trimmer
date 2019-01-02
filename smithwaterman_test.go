package main

import (
	"testing"
)

type alignTest struct {
	a        string
	b        string
	expected alignment
}

var alignmentTests = []alignTest{
	{
		a: "GGTTGACTA",
		b: "TGTTACGG",
		expected: alignment{
			matches:    5,
			mismatches: 0,
			gaps:       1,
			starta:     1,
			enda:       5,
			startb:     1,
			endb:       6,
		},
	},
}

func TestAlignment(t *testing.T) {

	for _, at := range alignmentTests {
		result := align([]byte(at.a), []byte(at.b))
		if result != at.expected {
			t.Errorf("Expected:\n%#v\nGot:\n%#v\n", at.expected, result)
		}
	}

}
