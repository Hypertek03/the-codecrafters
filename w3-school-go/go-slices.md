## Slices
Are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

There are several ways to create a slic:

1. Using the []datatype{values} format: E.g. Slice_name := []datatype{values} 

2. Create a slice from an array: E.g. var myarray = [length]datatype{values}

3. Using the make() function: E.g. slice_name := make([]type, length, capacity) 


