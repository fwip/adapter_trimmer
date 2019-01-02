package main

import (
	"fmt"

	"github.com/shenwei356/bio/seqio/fastx"
)

const errorRate = 0.2

func trim(r1, r2 *fastx.Record, a1, a2 []byte) (o1, o2 *fastx.Record) {
	s1 := r1.Seq.Seq
	s2 := r2.Seq.Seq
	s2rc := r2.Seq.RevCom().Seq
	aln := align(s1, s2rc)

	dist1 := aln.mismatches + aln.gaps + aln.starta + (len(s2) - aln.endb - 1)

	fmt.Printf("%s, %s, %s\n", s1, s2, s2rc)
	fmt.Printf("%#v, %d\n", aln, dist1)
	if float64(dist1) <= float64(aln.enda)*errorRate && (aln.starta == len(s2)-aln.endb-1) {
		fmt.Println("Doing a trim")
		// Trim

		// Extract the putative adapter sequences from the sequencing read
		r1adapter := s1[aln.enda+1:]
		r2adapter := s2[len(s2)-aln.startb:]

		// The read length might be longer than the adapter itself, and
		// in this case trim the putatitive adapter sequence to the length of the
		// adapter.

		// Note: read1 will contain Adapter #2, while read2 will contain Adapter #1

		if len(r1adapter) > len(a2) {
			r1adapter = r1adapter[:len(a2)]
		}
		if len(r2adapter) > len(a1) {
			r2adapter = r2adapter[:len(a1)]
		}
		// if (read1_adapter.size() > adapter2.size()) {
		// read1_adapter = read1_adapter.substr(0, adapter2.size());
		// }
		// if (read2_adapter.size() > adapter1.size()) {
		// read2_adapter = read2_adapter.substr(0, adapter1.size());
		// }

		// Now trim the adapters, as in many cases they are longer than the
		// putative adapter size.

		er1 := a2[:len(r1adapter)]
		er2 := a1[:len(r2adapter)]
		// expected_read1_adapter = adapter2.substr(0, read1_adapter.size());
		// expected_read2_adapter = adapter1.substr(0, read2_adapter.size());

		// If an exact match occurs set the trimming coordinates, else
		// perform two more local sequence alignments.

		enda := aln.enda + 1
		endb := aln.endb + 1
		if string(r1adapter) != string(er1) {
			aln2 := align(r1adapter, er1)
			dist := len(er1) - aln.matches
			if float64(dist) <= float64(len(r1adapter))*errorRate {
				enda = aln.enda + 1 + aln2.starta - aln2.startb
			}
		}
		if string(r2adapter) != string(er2) {
			aln2 := align(r2adapter, er2)
			dist := len(er2) - aln.matches
			if float64(dist) <= float64(len(r2adapter))*errorRate {
				endb = len(s2) - aln.startb + aln2.starta - aln2.startb
			}
		}
		// if (read1_adapter != expected_read1_adapter) {

		// alignment2 = sw2.align(DNA.unmask(read1_adapter), expected_read1_adapter);

		// dist2 = expected_read1_adapter.size() - alignment2.matches;

		// if (dist2 <= read1_adapter.length() * errorRate) {

		// enda = alignment1.enda + 1 + alignment2.starta - alignment2.startb;

		// }

		//} else {

		//enda = alignment1.enda + 1;
		//}

		//				if (read2_adapter != expected_read2_adapter ) {
		//
		//					alignment3 = sw3.align(DNA.unmask(read2_adapter), expected_read2_adapter);
		//
		//					dist3 = expected_read2_adapter.size() - alignment3.matches;
		//
		//					if (dist3 <= read2_adapter.length() * errorRate) {
		//
		//						endb = lenb - alignment1.startb + alignment3.starta - alignment3.startb;
		//					}
		//
		//				} else {
		//
		//					endb = lenb - alignment1.startb;
		//				}
		//
		//

		// Do trimming
		//o1 := r1.Seq.SubSeq(0, enda)
		//2 := r2.Seq.SubSeq(0, endb)
		o1 := r1.Clone()
		o2 := r2.Clone()
		o1.Seq.SubSeqInplace(0, enda)
		o2.Seq.SubSeqInplace(0, endb)
		return o1, o2

		//				if ( (enda != lena and lena - enda >= min_overlap) and (endb != lenb and lenb - endb >= min_overlap) ) {
		//
		//					read1.sequence = read1.sequence.substr(0, enda);
		//					read1.quality = read1.quality.substr(0, enda);
		//
		//					read2.sequence = read2.sequence.substr(0, endb);
		//					read2.quality = read2.quality.substr(0, endb);
		//
		//#pragma omp atomic
		//					++ntrimmed;
		//
		//				}
		//			}

	}
	return r1, r2
}
