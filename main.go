package main

import (
	"fmt"
	"reflect"
)

func main() {
	// a := nil this will be a compile-time error
	var a *int = nil
	var b int  // if we not initialize this variable the value were be ZERO 0 !!1!
	var c *int // if wer not initialize this pointer the value were be <nil> !!!!!

	fmt.Printf("Value of a => %v\n", a)
	// The folowing code not working because theres no memmory address to point to
	// fmt.Printf("Value of a pointer to => %v\n", *a)
	fmt.Printf("Type of a => %v\n", reflect.TypeOf(a))
	fmt.Println("----------------")

	fmt.Printf("Value of b => %v\n", b)
	fmt.Printf("Type of b => %v\n", reflect.TypeOf(b))
	fmt.Println("----------------")

	fmt.Printf("Value of c => %v\n", c)
	// The folowing code not working because theres no memmory address to point to
	// fmt.Printf("Value of c pointer to => %v\n", *c)
	fmt.Printf("Type of c => %v\n", reflect.TypeOf(c))
	fmt.Println("----------------")

	// if b == nil => we can't do that due to mismatch typing
	if a == nil {
		fmt.Println("The a is a nil pointer so 'a == nil' works")
	}

	if c == nil {
		fmt.Println("The c is a nil pointer so 'c == nil' works")
	}
	fmt.Println("----------------")

	value := 10
	c = &value

	// if b == nil => we can't do that due to mismatch typing
	if c == nil {
		fmt.Printf("After assigned a value to the pointer The 'c == nil' returns false")
	}

	fmt.Printf("Value of c => %v\n", c)
	fmt.Printf("Address memory of 'value' => %v\n", &value)
	fmt.Printf("Value of c pointer to => %v\n", *c)
	fmt.Printf("Type of c => %v\n", reflect.TypeOf(c))
	fmt.Println("----------------")

	value = 15
	fmt.Printf("Value of c pointer to => %v\n", *c)
}
