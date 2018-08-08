package main

import (
	"unicode/utf8"
	"fmt"
	"io"
	"bytes"
	"os"
)

func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

func init(){
	fmt.Printf("%#v\n", []rune("Hello, 世界"))
	fmt.Printf("%#v\n", []rune("Hello, 世"))
	fmt.Printf("%#v\n", []rune("Hello, "))
}
func f(){

}

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}


//func main(){
//	a = []int{1, 2, 3}
//	a = a[:copy(a, a[1:])] // 删除开头1个元素
//	a = a[:copy(a, a[N:])] // 删除开头N个元素
//	return
//}