package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

	var count int
	var varName string = "data"
	if *includeflag {
		if filename != "" {
			varName = sanitize(filepath.Base(filename))
		}
		fmt.Printf("const unsigned char %s[] = {", varName)
	}

	for {
		buf := make([]byte, bytesPerLine)
		n, err := reader.Read(buf)
		if n > 0 {
			if *includeflag {
				for i := 0; i < n; i++ {
					if count%bytesPerLine == 0 {
						fmt.Print("\n  ")
					}
					fmt.Printf("0x%02x, ", buf[i])
					count++
				}
			}
		}
		if err != nil {
			break
		}
	}
	if *includeflag {
		fmt.Printf("\n};\nconst unsigned int %s_len = %d;\n", varName, count)
	}
}

func sanitize(name string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	sanitizedName := reg.ReplaceAllString(name, "_")
	return sanitizedName
}
