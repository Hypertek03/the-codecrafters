package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print(" -- ")
		var cmd, num, base string
		// Read three inputs
		if _, err := fmt.Scan(&cmd, &num, &base); err != nil {
			fmt.Println("✦ Error: invalid input")
			continue
		}

		if cmd == "quit" {
			return
		}
		if cmd != "convert" {
			fmt.Println("✦ Error: invalid command")
			continue
		}

		base = strings.ToLower(base)
		var n int64
		var err error

		if base == "bin" {
			n, err = strconv.ParseInt(num, 2, 64)
		} else if base == "hex" {
			n, err = strconv.ParseInt(num, 16, 64)
		} else if base == "dec" {
			n, err = strconv.ParseInt(num, 10, 64)
		} else {
			fmt.Println("✦ Error: unknown base")
			continue
		}

		if err != nil {
			fmt.Println("✦ Error: invalid number for base", base)
			continue
		}

		if base == "dec" {
			fmt.Println("✦ Binary:", strconv.FormatInt(n, 2))
			fmt.Println("✦ Hex:", strings.ToUpper(strconv.FormatInt(n, 16)))
		} else {
			fmt.Println("✦ Decimal:", n)
		}
	}
}
