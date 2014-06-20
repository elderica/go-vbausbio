package main

import (
	"github.com/sbwhitecap/go-vbausbio/usbio"
	"flag"
	"strconv"
	"regexp"
	"log"
)

var (
	port = flag.Uint("port", 0, "USB-IO Port [0, 1]")
	data = flag.String("data", "0x00", "Bit pattern(10bit) (ex: 0b10101010, 0xaa, 170)")
)

func main() {
	flag.Parse()

	if *data == "" {
		flag.PrintDefaults()
		return
	}

	if *port > 1 {
		log.Fatalln("invalid port number")
	}

	// Parse data flag
	h := regexp.MustCompile("^(0x|0b)?").FindString(*data)
	var b int
	if h == "0x" {
		b = 16
	} else if h == "0b" {
		b = 2
	} else {
		b = 10
	}
	var bitpat uint64
	var err error
	if b == 10 {
		bitpat, err = strconv.ParseUint(*data, b, 10)
	} else {
		bitpat, err = strconv.ParseUint((*data)[2:], b, 10)
	}
	if err != nil {
		log.Fatalln(err)
	}

	if bitpat > 0x1ff {
		log.Println("warning: data is over 0x1ff")
	}

	if err := usbio.Start(); err != nil {
		log.Fatalf("Error in init: %s\n", err)
	}
	if err := usbio.Find(); err != nil {
		log.Fatalf("Error in search: %s\n", err)
	}
	if err := usbio.Set(uint8(*port), uint16(bitpat), 0); err != nil {
		log.Fatalf("Error in writing: %s\n", err)
	}
	log.Printf("Successfully wrote in USB-IO: port is %d, data is 0x%010b.\n", *port, bitpat)
}
