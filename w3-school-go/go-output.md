## Go Output Functions

This has to do with printing of an output.

GO output has three functions to output text:

1. Print(): prints its arguments with their default format. Example: fmt.Print(t)

2. Println(): it is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end. Example: 
  fmt.Println(i,j)


3. Printf():it first formats its argument based on the given formatting verb and then prints them.

There are two formatting verbs:

1. %v is used to print the value of the arguments.

2. %T is used to print the type of the arguments.

Example:  fmt.Printf("x has value: %v and type: %T\n", x, x)