package main


import (
	"fmt"
)

type fucker struct {
	Cells [3000]int
	Pointer int
}

func (f *fucker) Fuck(instruction string) {
	for i := 0; i < len(instruction); i++{
		c := instruction[i]
		switch c {
		case '+':
			f.Cells[f.Pointer]++
		case '-':
			f.Cells[f.Pointer]--
		case '>':
			f.Pointer++
		case '<':
			f.Pointer--
		case '.':
			fmt.Print(string(f.Cells[f.Pointer]))
		case ',':
			fmt.Scan(&f.Cells[f.Pointer])
		case '[':
			if f.Cells[f.Pointer] == 0 {
				for ; i < len(instruction) && instruction[i] != ']'; i++ {
					// do nothing
				}
			}
		case ']':
			if f.Cells[f.Pointer] != 0 {
				for ; instruction[i] != '['; i-- {
					// do nothing
				}
			}
		}
	}
}

var instruction = `
++++++++[>+>++>+++>++++>+++++>++++++>+++++++>++++++++>+++++++++>++++++++++>+++++++++++>++++++++++++>+++++++++++++>++++++++++++++>+++++++++++++++>++++++++++++++++<<<<<<<<<<<<<<<<-]>>>>>>>>>>>>>>>-.+<<<<<<<<<<<<<<<>>>>>>>>>>>>>---.+++<<<<<<<<<<<<<>>>>>>>>>>>>>>----.++++<<<<<<<<<<<<<<>>>>>>>>>>>>+++.---<<<<<<<<<<<<>>>>>>>>>>>>>>-.+<<<<<<<<<<<<<<>>>>>>>>>>>>>>---.+++<<<<<<<<<<<<<<>>>>>>>>>>>>>---.+++<<<<<<<<<<<<<>>>>>>--.++<<<<<<>>>>>>>>>>>>>.<<<<<<<<<<<<<>>>>>>>>>>>>>>>----.++++<<<<<<<<<<<<<<<>>>>>>>>>>>>>>---.+++<<<<<<<<<<<<<<>>>>>>>>>>>>>>----.++++<<<<<<<<<<<<<<.
`
func main() {
	f := new(fucker)
	f.Fuck(instruction)
}
