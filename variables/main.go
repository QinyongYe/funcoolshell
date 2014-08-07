package main

import (
	"fmt"
	"net/http"
)

const seed = "http://fun.coolshell.cn/n/2014"

func main(){
	url := seed
	buf := make([]byte, 30)

	for {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := res.Body.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		newnum := string(buf[0:n])
		fmt.Println(newnum)
		url = "http://fun.coolshell.cn/n/" + newnum

	}
}

