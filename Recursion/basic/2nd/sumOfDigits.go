package main

import "fmt"

func main(){
	// sum of digits of a number
	fmt.Println(sumOfDigits(134))
}

func sumOfDigits(n int) int {
	// 2. lets add a base condition. i.e, something that we already know
	if n <= 1 {
		return 1
	}

	// lets see how to make it to recursive
	// total = n%10 + n/10
	// the above thing can be done recursively
	// since the function has a return type, we will also add a return type
	return n%10 + sumOfDigits(n/10)
}
