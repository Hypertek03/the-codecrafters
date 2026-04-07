## Functions
A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.

# Create a Function
To create (often referred to as declare) a function, do the following:

* Use the func keyword.

* Specify a name for the function, followed by parentheses ().

* Finally, add code that defines what the function should do, inside curly braces {}.

# Call a Function
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

* In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():

# Naming Rules for Go Functions

* A function name must start with a letter.

* A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ ).

* Function names are case-sensitive.

* A function name cannot contain spaces.

* If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used.

## Function Parameters and Arguments

Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma:

# Function With Parameter Example
The following example has a function with one parameter (fname) of type string. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name.

## Function Returns
If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function.

## Recursion Functions
Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use the x variable as the data, which increments with 1 (x + 1) every time we recurse. The recursion ends when the x variable equals to 11 (x == 11). 



