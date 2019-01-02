package main

import (
	"testing"

	"github.com/shenwei356/bio/seq"
	"github.com/shenwei356/bio/seqio/fastx"
)

type trimTest struct {
	r1 string
	r2 string
	a1 string
	a2 string
	e1 string
	e2 string
}

var trimTests = []trimTest{
	{
		r1: "AAAAAAAGGGGGGGGGGGGG",
		r2: "AAAAAAACCCCCCCCCCCCC",
		a1: "GGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
		a2: "GGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
		e1: "AAAAAAA",
		e2: "AAAAAAA",
	},
}

func TestTrim(t *testing.T) {
	for _, test := range trimTests {
		r1f, _ := fastx.NewRecord(seq.DNA, []byte("r1"), []byte("r1"), []byte(test.r1))
		r2f, _ := fastx.NewRecord(seq.DNA, []byte("r2"), []byte("r2"), []byte(test.r2))
		result1, result2 := trim(
			r1f,
			r2f,
			[]byte(test.a1),
			[]byte(test.a2),
		)
		if string(result1.Seq.Seq) != test.e1 || string(result2.Seq.Seq) != test.e2 {
			t.Errorf("Expected:\n%s, %s\nGot:\n%s, %s\n", test.e1, test.e2, result1.Seq.Seq, result2.Seq.Seq)
		}

	}
}
