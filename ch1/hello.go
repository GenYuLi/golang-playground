package main

import "fmt"

func try_defer() {
	defer fmt.Println("i got defered haha")
	fmt.Println("im here")
}

func main() {
	try_defer()
	fmt.Printf("hello, %s!\n", "jenny")
	
}
