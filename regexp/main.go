package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
)

func check(err error) {
	if err!=nil {
		panic(err)
	}
}

func main(){
	text, err := ioutil.ReadFile("F:/GitHub/go/src/github.com/windinn/goUtils/regexp/cat.txt")
	check(err)

	// regexp don't support back reference:\1
	// so we use 2 steps to find all the candidates
	regx := regexp.MustCompile(`[A-Z][0-9][a-z][0-9][A-Z]|[0-9][]A-Z][a-z][A-Z][0-9]`)
	result := regx.FindAll(text, -1)

	// filter
	for _, s := range result {
		if s[0] == s[4] && s[1] == s[3] {
			fmt.Println(string(s))
		}
	}
}

