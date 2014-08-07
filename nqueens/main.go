package main

func NqueenAnswers(n int) [][]int {
	answers := make([][]int, 0)
	answer := make([]int, n)
	fullflags := (1 << n) - 1
	func answer(row, ld, rd int) {
		if row != fullflags {
			availables := fullflags &^(row | ld | rd )
			for availables {
				if row==0 {
					answer = answer[:0]
				}
				rightbit := availables & (^availables + 1)  // get the right most setted bit
				availables = availables - rightbit;
				test(row | rightbit, (ld | rightbit) << 1, (rd | rightbit) >> 1);
			}
		}
		else {
			answers := append(answers, row)
		}
	}
}



func main(){

}
