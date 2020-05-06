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
}