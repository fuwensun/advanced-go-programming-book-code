package main

//char* cs = "hello";
import "C"
import "./cgo_helper"

func main() {
	cgo_helper.PrintCString(C.cs)
}