package main

import (
	"fmt"
	"crypto/sha1"
	"encoding/hex"
)

type nqueen struct {
	solutions [][]int
	currentSolution []int
	fullflags int
}

func (nq *nqueen) findSolutions(row, ld, rd, depth int) {
	if row != nq.fullflags {
		availables := nq.fullflags &^ (row | ld | rd )
		for availables != 0 {
			rightbit := availables & (^availables + 1)  // get the right most setted bit
			availables = availables-rightbit;
			nq.currentSolution[depth] = bitpos(rightbit)
			nq.findSolutions(row|rightbit, (ld|rightbit)<<1, (rd|rightbit)>>1, depth+1);
		}
	} else {
		solution := make([]int, len(nq.currentSolution))
		copy(solution, nq.currentSolution)
		nq.solutions = append(nq.solutions, solution)
	}
}

func bitpos(n int) int {
	switch(n) {
	case 1:
		return 1
	case 2:
		return 2
	case 4:
		return 3
	case 8:
		return 4
	case 16:
		return 5
	case 32:
		return 6
	case 64:
		return 7
	case 128:
		return 8
	case 256:
		return 9
	case 512:
		return 10
	case 1024:
		return 11
	default:
		panic("not implemented")
	}
}

func newNqueen(n uint) *nqueen{
	nq := &nqueen{}
	nq.fullflags = (1 << n) -1
	nq.solutions = make([][]int, 0)
	nq.currentSolution = make([]int, n)
	return nq
}

func NqueenSolutions(n uint) [][]int {
	nq := newNqueen(n)
	nq.findSolutions(0,0,0,0)
	return nq.solutions
}



func main(){
	solutions := NqueenSolutions(9)
	for _,s := range solutions {
		sum := sha1.Sum([]byte("zWp8LGn01wxJ7" + toString(s) + "\n"))
		if hex.EncodeToString(sum[:]) == "e48d316ed573d3273931e19f9ac9f9e6039a4242" {
			fmt.Println(sum)
			fmt.Println(toString(s))
		}
	}
}

func toString(solution []int) string {
	result := make([]byte, len(solution))
	for i,v := range solution {
		result[i] = byte(v + '0')
	}
	return string(result)
}
