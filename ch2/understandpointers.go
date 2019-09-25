package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println("&x = %g\n",&x) 
	fmt.Println("x = %g\n", x)
	fmt.Println("p = %g\n", *p)
}
