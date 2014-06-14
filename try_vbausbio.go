package main

import (
	"fmt"
	"syscall"
	"log"
	"flag"
	//"unsafe"
)

func main() {
	flag.Parse()
	dll, err := syscall.LoadDLL("vbausbio.dll")
	if err != nil {
		log.Fatal(err, dll)
	}
	defer dll.Release()

	proc, err := dll.FindProc("uio_out")
	if err != nil {
		log.Fatal(err)
	}
	var (
		r1, r2 uintptr
	)
	if flag.NArg() > 0 {
		r1, r2, err = proc.Call(0, 255, 0)
	} else {
		r1, r2, err = proc.Call(0, 0, 0)
	}

	fmt.Printf("%v, %v, %v\n", r1, r2, err)
}
