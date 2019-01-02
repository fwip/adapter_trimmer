package main

type sequence []byte

type alignment struct {
	matches    int
	mismatches int
	gaps       int
	starta     int
	enda       int
	startb     int
	endb       int
}

const matchScore = 3
const mismatchScore = -3
const gapScore = -2

func align(a, b sequence) alignment {

	// Initialize matrix
	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
	}

	// Populate matrix
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(b)+1; col++ {
			alnscore := matrix[row-1][col-1]
			if a[row-1] == b[col-1] {
				alnscore += matchScore
			} else {
				alnscore += mismatchScore
			}
			gap1score := matrix[row-1][col] + gapScore
			gap2score := matrix[row][col-1] + gapScore
			if gap1score > alnscore {
				alnscore = gap1score
			}
			if gap2score > alnscore {
				alnscore = gap2score
			}
			if alnscore < 0 {
				alnscore = 0
			}
			matrix[row][col] = alnscore
		}
	}

	// Find best path
	highest := 0
	bestrow := 0
	bestcol := 0
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(b)+1; col++ {
			if matrix[row][col] > highest {
				highest = matrix[row][col]
				bestrow = row
				bestcol = col
			}
		}
	}

	// Print out matrix (debugging)
	//fmt.Print("\t")
	//for _, r := range b {
	//	fmt.Printf("%s\t", string(r))
	//}
	//fmt.Println()

	//for row := 0; row < len(matrix); row++ {
	//	if row > 0 {
	//		fmt.Printf("%s\t", string(rune(a[row-1])))
	//	} else {
	//		fmt.Printf("\t")
	//	}
	//	for col := 0; col < len(b)+1; col++ {
	//		fmt.Printf("%d\t", matrix[row][col])
	//	}
	//	fmt.Println()
	//}

	// Trace back to start
	i, j := bestrow, bestcol
	mismatches := 0
	matches := 0
	gaps := 0
	for {
		val := matrix[i][j]
		//fmt.Printf("%d at %d, %d\n", val, i, j)
		if val == 0 {
			break
		}
		if a[i-1] == b[j-1] {
			if matrix[i-1][j-1] == val-matchScore {
				i, j = i-1, j-1
				matches++
				continue
			}
		}
		if matrix[i-1][j-1] == val-mismatchScore {
			i, j = i-1, j-1
			mismatches++
			continue
		}
		gaps++
		if matrix[i-1][j] == val-gapScore {
			i--
			continue
		}
		j--
	}

	aln := alignment{
		matches:    matches,
		enda:       bestcol - 1,
		endb:       bestrow - 1,
		mismatches: mismatches,
		gaps:       gaps,
		starta:     i,
		startb:     j,
	}

	return aln
}
