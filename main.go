//package main specifies an entry point for the application
package main

import (
	"net/http"
	"github.com/pluralsight/go-gettingstarted/webservice/controllers"
	"fmt"
	"errors"
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
	//register routing
	controllers.RegisterControllers();
	//http server
	//use nil to use the dfault ServeMux
	http.ListenAndServe(":3000", nil);


	fmt.Println(testFunction(1, 2, 3));
	fmt.Println(testErrorFunc());
	num, err := testMultipleReturnFunc(5);
	fmt.Println(num, err);
	//can use a write-only variable, which is an underscore.
	//Allows you to dump data into it without having to use it, so you don't get an error for an unused variable
	_, err1 :=  testMultipleReturnFunc(10);
	fmt.Println(err1);
}

//if all params are the same type, can pass in data type once
//return type comes after parenthesis 
func testFunction(num1, num2, num3 int) bool {
	fmt.Println(num1);
	return true;
}

func testErrorFunc() error {
	return errors.New("Something is wrong");
}

//multiple returns
func testMultipleReturnFunc(num1 int) (int, error) {
	return num1, nil;
}

func notes() {
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

	//A colon in the square brackets of an array creates a slice of that array from the beginning of the array to the end of it.
	sliceArr := arr1[:];
	fmt.Println(arr1, sliceArr);

	//The slice is pointing to the data that the array is keeping
	//Any changes made to the array are going to be reflected in the slice, and any changes made in the slice are going to be reflected back to the underlying array.
	//still bound to a fixed sized array
	arr1[1] = 10;
	sliceArr[2] = 15;
	fmt.Println(arr1, sliceArr);

	//A slice is not a fixed size entity
	//Go will automatically resize the underlying array by creating a new array & having the slice point to it
	 sliceArr2 := []int{1, 2, 3}
	 fmt.Println(sliceArr2)
	 sliceArr2 = append(sliceArr2, 4, 5);
	 fmt.Println(sliceArr2);

	 //Creates a slice of SliceArr2 starting from index 1 to the end
	 sliceArr3 := sliceArr2[1:];
	 //Creates a slice of SliceArr2 starting from index 0 up to & not including index 2
	 sliceArr4 := sliceArr2[:2];
	 //Creates a slice of SliceArr2 starting from index 1 up to & not including index 2
	 sliceArr5 := sliceArr2[1:2];
	 fmt.Println(sliceArr3, sliceArr4, sliceArr5);

	 //keys are of type string, values are of type int
	 myMap := map[string]int{"foo":42}
	 fmt.Println(myMap);
	 fmt.Println(myMap["foo"]);
	 
	 delete(myMap, "foo");
	 fmt.Println(myMap);

	 //Define a type called 'user'. That user type is going to be a struct
	 type user struct {
		ID int
		FirstName string
		LastName string
	 }

	 var u user;
	 u.ID = 1;
	 u.FirstName = "Gerax"
	 u.LastName = "Sotelo"
	 fmt.Println(u);

	 //Must end final line with a comma or a closing curly brace, or else Go will add a semicolon
	 u2 := user { ID: 1, 
		FirstName: "Gerax", 
		LastName: "Sotelo",
	 }
	 fmt.Println(u2);

	 //Loop till condition construct
	 var num int
	 for num < 5 {
		 println(num);
		 num++;
	 }

	 //Loop till condition with a post clause
	 //has to have 3 terms in the for loop to indicate to the compiler that this is a loop till condition with a post clause
	 for num := 0; num < 5; num++ {
		 println(num);
	 }

	 //Infinite loop
	 var num2 int
	 for { // instead of for ; ; {
		 if i == 5 {
			 break;
		 }
		 println(num2);
		 num2++;
	 }

	 //Looping over collections
	 //loop over a slice
	 loopingSlice := []int{1, 2, 3};
	 for i := 0; i < len(loopingSlice); i++ {
		 println(loopingSlice[i]);
	 }
	 //SAME THING
	 //use (i) indexer & (v) value. Set them equal to the 'range' keyword followed by the collection being iterated over.
	 //The 'range' keyword tells the compiler that we're passing in a collection type & then it returns out the index in the 1st variable that gets returned & the value at that index in the 2nd variable that's returned
	 for i, v := range loopingSlice {
		println(i, v);
	 }
	 //loop over a map
	 wellKnownPorts := map[string]int {"http": 80, "https": 443}
	 for k, v := range wellKnownPorts {
		println(k, v);
	 }
	 //ignore the second value
	 wellKnownPorts2 := map[string]int {"http": 80, "https": 443}
	 for k := range wellKnownPorts2 {
		println(k);
	 }
	 //ignore the first value
	 wellKnownPorts3 := map[string]int {"http": 80, "https": 443}
	 for _, v := range wellKnownPorts3 {
		println(v);
	 }

	 //a panic is similar to an exception in other languages
	 // panic("Uh oh. Something went wrong.");

	 //if statement
	 number1 := 5;
	 number2 := 5;
	 if number1 == number2 {
		 println("numbers are equal");
	 } else if number1 > number2 {
		println("first number is greater");
	 }
}