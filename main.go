package main

import (
	"fmt"
)

//constant block
const (
	first = 1;
	second = "second";
	//Every time iota is reused it increments its value by one. Long chains of constants can be built and the values of the constants will change.
	third = iota;
	fourth = iota;
	fifth = iota + 2;
	sixth = 2 << iota; //bit shift
	//If a value is not specified for a constant below another constant, it'll reuse the constant expression above it
	seventh
)

const (
	//iota will reset in a new constant block
	eigth = iota;
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

	//Constants have to be initialized when they are declared
	//Value of a constant has to be able to be determined at 'compile time'. Can't assign it to a return value of a function that gets evaluated at 'runtime'
	const pi = 3.1415
	fmt.Println(pi)

	//Go will dynamically interpret the type of a constant, unless the constant's type is specified
	const d = 3;
	fmt.Println(d + 3);
	fmt.Println(d + 1.2);

	const e int = 3;
	fmt.Println(float32(e) + 1.5);

	//from constant block
	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eigth);

	//all elements in an array must be the same type
	var arr [3]int;
	arr[0] = 1;
	arr[1] = 2;
	arr[2] = 3;
	fmt.Println(arr);

	arr1 := [3]int{4, 5, 6}
	fmt.Println(arr1);
}