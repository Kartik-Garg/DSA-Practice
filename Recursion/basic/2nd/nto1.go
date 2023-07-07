package main

import "fmt"

func main(){
	// use recursion to print from n to 1
	recursive(5)
}

func recursive(n int){
	// base condition will be what we already know
	if (n<1){
		return
	}

	fmt.Println(n)
	// also since the function is being called at the end, so it is tail recursion.
	recursive(n-1)
}
