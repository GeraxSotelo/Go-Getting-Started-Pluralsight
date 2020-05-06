package main

import (
	"fmt"
)

func main() {
	//Initialize in separate lines. var keyword, variable name, type
	var a int;
	a = 42
	fmt.Println(a)

	//Initialize on same line. With floating-point variables, size must be specified. (float32, float64)
	var f float32 = 3.14
	fmt.Println(f);

	//Implicit initialization syntax. Using :=  Go will imply the data type
	firstName := "Gerax"
	fmt.Println(firstName)

	b := true;
	fmt.Println(b);

	c := complex(3, 4);
	fmt.Println(c);
	
	//real & img functions
	r, i := real(c), imag(c);
	fmt.Println(r, i);

	//Pointer data type. A variable will point to another location in memory that's going to hold the relevant information
	//lastName is going to be a pointer to a string value. It must be initialized before assigning anything to it
	var lastName *string = new(string)
	//Dereference operation (*). Dereferencing is reaching through the pointer, grabing the data, and bringing it back
	*lastName = "Sotelo";
	fmt.Println(*lastName);

	//Address of operator (&). Go allows to create & declare a variable, then create a pointer that points to the variable.
	middleInit := "Nef";
	ptr := &middleInit;
	fmt.Println(ptr, *ptr);

	middleInit = "N";
	//The value of the pointer 'ptr' changes, but the memory address doesn't. The 'middleInit' variable is still in the same location
	fmt.Println(ptr, *ptr);
}