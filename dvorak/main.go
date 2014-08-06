package main

import (
	"fmt"
)

var DvorakQwertMapper = map[rune]rune{
	'[':'-', '{':'_',
	']':'=', '}':'+',
	'\'':'q', '"':'Q',
	',':'w', '<':'W',
	'.':'e', '>':'E',
	'p':'r', 'P':'R',
	'y':'t', 'Y':'T',
	'f':'y', 'F':'Y',
	'g':'u', 'G':'U',
	'c':'i', 'C':'I',
	'r':'o', 'R':'O',
	'l':'p', 'L':'P',
	'/':'[', '?':'{',
	'=':']', '+':'}',
	'o':'s', 'O':'S',
	'e':'d', 'E':'D',
	'u':'f', 'U':'F',
	'i':'g', 'I':'G',
	'd':'h', 'D':'H',
	'h':'j', 'H':'J',
	't':'k', 'T':'K',
	'n':'l', 'N':'L',
	's':';', 'S':':',
	'-':'\'', '_':'"',

	';':'z', ':':'Z',
	'q':'x', 'Q':'X',
	'j':'c', 'J':'C',
	'k':'v', 'K':'V',
	'x':'b', 'X':'B',
	'b':'n', 'B':'N',
	'w':',', 'W':'<',
	'v':'.', 'V':'>',
	'z':'/', 'Z':'?',
}

//var QwertDvorakMapper = map[string]string{
//	'-':'[', '_':'{',
//	'=':']', '+':'}',
//	'q':''', 'Q':'\'',
//	'w':',', 'W':'<',
//	'e':'.', 'E':'>',
//	'r':'p', 'R':'P',
//	't':'y', 'T':'Y',
//	'y':'f', 'Y':'F',
//	'u':'g', 'U':'G',
//}

func DvorakMaptoQwert(dvorak string) string {
	newstr := make([]rune, len(dvorak))
	for i, v := range dvorak {
		newv, ok := DvorakQwertMapper[v]
		if ok {
			newstr[i] = newv
		}else {
			newstr[i] = v
		}
	}
	return string(newstr)

}

func main(){
	var input = `macb() ? lpcbyu(&gbcq/_\021%ocq\012\0_=w(gbcq)/_dak._=}_ugb_[0q60)s+`

	fmt.Print(DvorakMaptoQwert(input))
}
