package main

import "fmt"

var src =  "abcdefghijklmnopqrstuvwxyz"
var dest = "pvwdgazxubqfsnrhocitlkeymj"

var mapper [26]byte

func initMapper() {
	for i, v := range dest {
		mapper[v - 'a'] = src[i]
	}
}

func mapping(input string) string {
	var newstr = make([]byte, len(input))
	for i, v := range input {
		if v >= 'a' && v <= 'z' {
			newstr[i] = mapper[v-'a']
		}else {
			newstr[i] = byte(v)
		}
	}
	return string(newstr)
}


func main(){
	var input = `Wxgcg txgcg ui p ixgff, txgcg ui p epm. I gyhgwt mrl lig txg ixgff wrsspnd tr irfkg txui hcrvfgs, nre, hfgpig tcm liunz txg crt13 ra "ixgff" tr gntgc ngyt fgkgf.`
	initMapper()
	fmt.Print(mapping(input))
}
