package cgo_helper

//#include <stdio.h>
import "C"

type CChar C.char

func (p *CChar) GoString() string {
	return C.GoString((*C.char)(p))
}

func PrintCString(cs *C.char) {
	C.puts(cs)	//ok
}

//
//func PrintCString(cs *C.char) {
//	//print(((*CChar)cs).GoString()))
//	//fmt.Print((CChar(*cs)).GoString())
//	C.puts(cs)	//ok
//}

