package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print("--- ")
		var cmd, val, base string
		n, _ := fmt.Scanln(&cmd, &val, &base)

		if n == 0 {
			continue
		}
		if cmd == "quit" {
			fmt.Println("Goodbye my dear!")
			return
		}
		if n < 3 || cmd != "convert" {
			continue
		}

		switch base {
		case "hex":
			if d, err := strconv.ParseInt(val, 16, 64); err == nil {
				fmt.Printf(" Decimal: %d\n", d)
			} else {
				fmt.Printf("Error: \"%s\" is not valid hex\n", val)
			}
		case "bin":
			if d, err := strconv.ParseInt(val, 2, 64); err == nil {
				fmt.Printf(" Decimal: %d\n", d)
			} else {
				fmt.Printf("Error: \"%s\" is not valid binary\n", val)
			}
		case "dec":
			if d, err := strconv.ParseInt(val, 10, 64); err == nil {
				fmt.Printf(" Binary:  %s\n", strconv.FormatInt(d, 2))
				fmt.Printf(" Hex:     %s\n", strings.ToUpper(strconv.FormatInt(d, 16)))
			} else {
				fmt.Printf("Error: \"%s\" is not a valid decimal\n", val)
			}
		default:
			fmt.Println("Error: base must be hex, bin, or dec")
		}
	}
}
