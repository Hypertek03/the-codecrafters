package main

import "fmt"

func main() {
	var cmd string
	fmt.Println(" Manuel's Calculator ")
	fmt.Println("Type 'help' for list of commands")

	for {
		fmt.Print(" -- ")
		fmt.Scan(&cmd)

		if cmd == "quit" {
			fmt.Println("Goodbye my dear!")
			return
		}

		if cmd == "help" {
			fmt.Println("Commands:")
			fmt.Println(" add a b = Addition ")
			fmt.Println(" sub a b = Subtraction ")
			fmt.Println(" mul a b = Multiplication ")
			fmt.Println(" div a b = Division ")
			fmt.Println(" quit")
			continue
		}

		if cmd != "add" && cmd != "sub" && cmd != "mul" && cmd != "div" {
			fmt.Println(" Error: unknown command (type 'help' for list of command)")
			continue
		}

		var a, b float64
		var s1, s2 string

		n, _ := fmt.Scan(&s1, &s2)

		if n != 2 {
			fmt.Println(" Error: wrong number of arguments (need 2)")
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		_, err1 := fmt.Sscan(s1, &a)
		_, err2 := fmt.Sscan(s2, &b)

		if err1 != nil || err2 != nil {
			fmt.Println(" Error: arguments must be numbers")
			continue
		}

		if cmd == "div" && b == 0 {
			fmt.Println(" Error: cannot divide by zero")
			continue
		}

		switch cmd {
		case "add":
			fmt.Println(" Result:", a+b)
		case "sub":
			fmt.Println(" Result:", a-b)
		case "mul":
			fmt.Println(" Result:", a*b)
		case "div":
			fmt.Println(" Result:", a/b)
		}
	}
}
