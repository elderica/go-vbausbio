package main

import (
	"github.com/sbwhitecap/go-vbausbio/usbio"
	"flag"
	"log"
)

var (
	port = flag.Uint("port", 0, "USB-IO Port [0, 1]")
	data = flag.Uint("data", 0, "Data [0x00, 0xff]")
)

func main() {
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.PrintDefaults()
		return
	}
	if err := usbio.Start(); err != nil {
		log.Fatalf("Error in init: %s\n", err)
	}
	if err := usbio.Find(); err != nil {
		log.Fatalf("Error in search: %s\n", err)
	}
	if err := usbio.Set(uint8(*port), uint8(*data), 0); err != nil {
		log.Fatalf("Error in writing: %s\n", err)
	}
	log.Println("Successfully wrote out for USB-IO")
}
