package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const bytesPerLine = 12 // For -i option, 12 bytes per line is common

var includeflag = flag.Bool("i", false, "Output in C include file style")

func main() {
	flag.Parse()

	var filename string
	if len(flag.Args()) > 0 {
		filename = flag.Args()[0]
	}

	var reader *bufio.Reader
	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	if *includeflag {
		fmt.Print("unsigned char data[] = {")
	}
	var offset int64 = 0
	for {
		buf := make([]byte, bytesPerLine)
		n, err := reader.Read(buf)
		if n > 0 {
			if *includeflag {
				for i := 0; i < n; i++ {
					if offset%bytesPerLine == 0 {
						fmt.Print("\n  ")
					}
					fmt.Printf("0x%02x, ", buf[i])
					offset++
				}
			} else {
				fmt.Printf("%08x: ", offset)
				for i := 0; i < n; i++ {
					fmt.Printf("%02x ", buf[i])
				}
				for i := n; i < bytesPerLine; i++ {
					fmt.Print("   ")
				}
				fmt.Print(" ")
				for i := 0; i < n; i++ {
					if buf[i] >= 32 && buf[i] < 127 {
						fmt.Printf("%c", buf[i])
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print("\n")
				offset += int64(n)
			}
		}
		if err != nil {
			break
		}
	}
	if *includeflag {
		fmt.Print("\n};")
	}
}
