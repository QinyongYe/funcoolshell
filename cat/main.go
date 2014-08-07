package main

import (
	"regexp"
	"fmt"
)

func main(){
	var pattern, err = regexp.Compile("([A-Z][0-9][a-z])|([0-9][A-Z][a-z])")
	if(err != nil) {
		fmt.Println(err)
		return
	}


}

