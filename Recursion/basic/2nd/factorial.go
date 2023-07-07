package main

import "fmt"

func main(){
	// use recursion to print factorial of n
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	// we also need a base condition, and it is generally a thing which we already know
	if (n <= 1){
		return 1
	}
	// this is the formula so this will be called recursively --- > fact n = n * fact(n-1)
	return n * factorial(n-1)
}
