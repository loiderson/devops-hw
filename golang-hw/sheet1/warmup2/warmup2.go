// package main

// (IF PACKAGE MAIN DELETED)

// The program/compiler is looking for the main package, but instead finds the import delcaration instead. All programs must have the package main declared.

import "fmt"

func start() {

	fmt.Println("Hello, this is warmup 2.")

}

// (IF PACKAGE MAIN IS INCLUDED BUT THE FUNC IS NOT CALLED MAIN)

// "function main is undeclared in the main package", because we specified that we are using the main package, it is expecting us to have a main function in our program. Otherwise the compiler will throw an error.
