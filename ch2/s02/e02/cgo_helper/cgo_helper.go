package cgo_helper

import "C"

type CChar C.char

func (p *CChar) GoString() string {
	return C.GoString((*C.char)(p))
}

//func PrintCString(cs *C.char) {
//	print(((*CChar)(cs.GoString())))
//}


func PrintCString(cs *CChar) {
	print(((cs.GoString())))
	//print(((*CChar)(cs.GoString())))
}
