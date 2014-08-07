package main

import (
	"fmt"
)

func decode26(str string) int64 {
	var result int64 = 0
	for _,c := range str {
		result = result * 26 + int64(c - 'A')
	}
	return result
}

func encode26(n int64) string{
	result := make([]byte, 0)
	for {
		result = append(result, byte(n % 26 + 'A'))
		n = n/26
		if n==0 {
			break
		}
	}
	reverse(result)
	return string(result)
}

func reverse(str []byte) {
	n := len(str)
	for i := 0; i < n/2;i++{
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
}

func main(){

	result := decode26("COOLSHELL")/decode26("SHELL")
	fmt.Println(encode26(result))
	// fuck, coolshell must must has a bug in this problem
}
